#!/usr/bin/env bash
set -euo pipefail

TOPIC=${1:-}

if [ -z "$TOPIC" ]; then
  echo "Usage: $0 <topic>"
  exit 2
fi

docker compose exec kafka /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --describe --topic "$TOPIC"
