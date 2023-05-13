package kafka

import (
	"context"
	"fmt"

	"github.com/MihasBel/data-bus/broker/bustopic"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

const (
	MsgErrCreateProducer    = "cant create producer connection to broker: %w "
	MsgErrCreateAdminClient = "cant create admin client connection to broker: %w"
	MsgErrCreateTopics      = "cant create topics in broker: %w"
	MsgWarnUnknownType      = "unknown message type in config: %v"
	MsgErrUnknownType       = "unknown message type in received message: %v"
	MsgErrEmptyMsg          = "empty message data in received message type: %v"
)

type Broker struct {
	log      *zerolog.Logger
	p        *kafka.Producer
	ac       *kafka.AdminClient
	msgTypes []string
	cfg      Config
}

func New(cfg Config, msgTypes []string, l zerolog.Logger) *Broker {
	l = l.With().Str("cmp", "broker").Logger()

	return &Broker{
		log:      &l,
		cfg:      cfg,
		msgTypes: msgTypes,
	}
}

func (b *Broker) Start(ctx context.Context) error {
	if !b.cfg.Enabled {
		return nil
	}

	// create an admin client connection
	ac, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": b.cfg.ServerURL,
	})
	if err != nil {
		b.log.Error().Err(err).Msg(MsgErrCreateAdminClient)
		return errors.Wrap(err, MsgErrCreateAdminClient)
	}

	// get enabled topics based on message types
	topics := b.getCurrentTopics(b.msgTypes)
	kafkaTopics := make([]kafka.TopicSpecification, len(topics))
	for i, topic := range topics {
		kafkaTopics[i] = kafka.TopicSpecification{
			Topic:         topic,
			NumPartitions: b.cfg.PartitionsCount,
		}
	}
	if _, err = ac.CreateTopics(ctx, kafkaTopics); err != nil {
		b.log.Error().Err(err).Msg(MsgErrCreateTopics)
		return errors.Wrap(err, MsgErrCreateTopics)
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": b.cfg.ServerURL,
		"message.max.bytes": 5 << 20, // 5 MB
	})
	if err != nil {
		b.log.Error().Err(err).Msg(MsgErrCreateProducer)
		return errors.New(MsgErrCreateProducer)
	}

	go func(drs chan kafka.Event) {
		for ev := range drs {
			m, ok := ev.(*kafka.Message)
			if !ok {
				continue
			}

			if err := m.TopicPartition.Error; err != nil {
				b.log.Error().Err(err).Msgf("Delivery error: %v", m.TopicPartition)
			}
		}
	}(p.Events())

	b.p = p
	b.ac = ac

	return nil
}

func (b *Broker) Stop(_ context.Context) error {
	if !b.cfg.Enabled {
		return nil
	}

	b.p.Flush(30 * 1000)

	b.p.Close()
	b.ac.Close()

	return nil
}

func (b *Broker) produce(topic string, data []byte) error {
	if !b.cfg.Enabled {
		return nil
	}

	err := b.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)

	if kafkaError, ok := err.(kafka.Error); ok && kafkaError.Code() == kafka.ErrQueueFull {
		b.log.Info().Str("topic", topic).Msg("Kafka local queue full error - Going to Flush then retry...")
		flushedMessages := b.p.Flush(30 * 1000)
		b.log.Info().Str("topic", topic).
			Msgf("Flushed kafka messages. Outstanding events still un-flushed: %d", flushedMessages)

		return b.produce(topic, data)
	}

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("produce %s fail", topic))
	}

	return nil
}

func (b *Broker) getCurrentTopics(types []string) (topics []string) {
	for _, t := range types {
		switch t {
		case *bustopic.Info:
			topics = append(topics, *bustopic.Info)
		case *bustopic.Errors:
			topics = append(topics, *bustopic.Errors)
		default:
			b.log.Warn().Msgf(MsgWarnUnknownType, t)
			continue
		}
	}
	return topics
}
