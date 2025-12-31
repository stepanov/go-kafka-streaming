package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	var (
		broker  = flag.String("broker", getEnv("KAFKA_BROKER", "localhost:9092"), "Kafka broker address")
		topic   = flag.String("topic", getEnv("KAFKA_TOPIC", "test-topic"), "Kafka topic")
		groupID = flag.String("group", getEnv("KAFKA_GROUP", "group-1"), "Consumer group ID")
	)
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{*broker},
		GroupID:  *groupID,
		Topic:    *topic,
		MinBytes: 1,    // 1B
		MaxBytes: 10e6, // 10MB
	})
	defer r.Close()

	log.Printf("subscribing to broker=%s topic=%s group=%s", *broker, *topic, *groupID)

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				log.Println("shutting down subscriber")
				return
			}
			log.Printf("read error: %v", err)
			time.Sleep(time.Second)
			continue
		}
		log.Printf("received: partition=%d offset=%d key=%s value=%s", m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}

func getEnv(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}
