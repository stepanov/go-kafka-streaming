# go-kafka-streaming

Simple Go examples for publishing and subscribing to Kafka (using segmentio/kafka-go).

## Requirements

- Go 1.20+
- Docker & Docker Compose (to run Apache Kafka in KRaft mode, no Zookeeper)

## Quick start

1. Start Apache Kafka (KRaft, no Zookeeper):

   docker compose up -d

2. Build binaries:

   make build

3. Run subscriber (in one terminal):

   ./bin/subscriber -topic test-topic

4. Run publisher (in another terminal):

   ./bin/publisher -topic test-topic -count 10 -interval 500ms

Messages will be printed by the subscriber.

## Environment variables / flags

- `KAFKA_BROKER` or `-broker`: broker address (default localhost:9092)
- `KAFKA_TOPIC` or `-topic`: topic name (default `test-topic`)
- `KAFKA_GROUP` or `-group`: consumer group id for the subscriber

## Notes

- This example uses Apache Kafka in KRaft mode (no Zookeeper) in `docker-compose.yml` for easy local testing.

### Topic management helpers

This repo includes helper scripts to manage topics via the Kafka CLI inside the container:

- Create: `./scripts/create-topic.sh <topic> [partitions] [replication-factor]`
- List: `./scripts/list-topics.sh`
- Describe: `./scripts/describe-topic.sh <topic>`
- Increase partitions: `./scripts/increase-partitions.sh <topic> <new-partition-count>`

You can also use the `Makefile` targets:

- `make topic-create TOPIC=my-topic PARTITIONS=3 RF=1`
- `make topic-list`
- `make topic-describe TOPIC=my-topic`
- `make topic-alter TOPIC=my-topic PARTITIONS=6`
- Topics are auto-created by the broker if allowed; you can also create topics manually if desired.
