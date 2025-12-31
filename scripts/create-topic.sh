#!/usr/bin/env bash
set -euo pipefail

TOPIC=${1:-}
PARTITIONS=${2:-1}
RF=${3:-1}

if [ -z "$TOPIC" ]; then
  echo "Usage: $0 <topic> [partitions] [replication-factor]"
  exit 2
fi

# Run kafka-topics.sh inside the kafka container
# Note: replication-factor >1 requires multiple brokers

docker compose exec kafka /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 \
  --create --topic "$TOPIC" --partitions "$PARTITIONS" --replication-factor "$RF"
