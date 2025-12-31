#!/usr/bin/env bash
set -euo pipefail

TOPIC=${1:-}
PARTITIONS=${2:-}

if [ -z "$TOPIC" ] || [ -z "$PARTITIONS" ]; then
  echo "Usage: $0 <topic> <new-partition-count>"
  exit 2
fi

docker compose exec kafka /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --alter --topic "$TOPIC" --partitions "$PARTITIONS"
