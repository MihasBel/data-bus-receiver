FROM golang:1.20.4-alpine as builder

ENV CGO_ENABLED=1

RUN apk update && apk add --no-cache make git build-base musl-dev librdkafka librdkafka-dev
WORKDIR /go/src/github.com/MihasBel/data-bus-receiver
COPY . ./

RUN echo "build binary" && \
    export PATH=$PATH:/usr/local/go/bin && \
    go mod download && \
    go build -tags musl /go/src/github.com/MihasBel/data-bus-bus/cmd/main.go && \
    mkdir -p /data-bus-bus && \
    mv main /data-bus-bus/main && \
    rm -Rf /usr/local/go/src

FROM alpine:latest as app
WORKDIR /data-bus-receiver
COPY --from=builder /data-bus-receiver/. /data-bus-receiver/
CMD ./main
