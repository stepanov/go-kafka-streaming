BINARY := bin

.PHONY: build up down publish consume topic-create topic-list topic-describe topic-alter

build:
	go build -o $(BINARY)/publisher ./cmd/publisher
	go build -o $(BINARY)/subscriber ./cmd/subscriber

up:
	docker compose up -d

down:
	docker compose down -v

publish:
	# publish 5 messages quickly
	$(BINARY)/publisher -count 5 -interval 100ms

consume:
	$(BINARY)/subscriber

# Topic helpers (use TOPIC=... PARTITIONS=... RF=... )
topic-create:
	@./scripts/create-topic.sh "${TOPIC}" "${PARTITIONS}" "${RF}"

topic-list:
	@./scripts/list-topics.sh

topic-describe:
	@./scripts/describe-topic.sh "${TOPIC}"

topic-alter:
	@./scripts/increase-partitions.sh "${TOPIC}" "${PARTITIONS}"
