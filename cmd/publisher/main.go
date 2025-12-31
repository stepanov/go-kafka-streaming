package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	var (
		broker = flag.String("broker", getEnv("KAFKA_BROKER", "localhost:9092"), "Kafka broker address")
		topic  = flag.String("topic", getEnv("KAFKA_TOPIC", "test-topic"), "Kafka topic")
		count  = flag.Int("count", 0, "Number of messages to send (0 = unlimited)")
		intv   = flag.Duration("interval", time.Second, "Interval between messages")
	)
	flag.Parse()

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{*broker},
		Topic:    *topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	log.Printf("publishing to broker=%s topic=%s", *broker, *topic)
	msgIdx := 1
	for {
		select {
		case <-ctx.Done():
			log.Println("shutting down publisher")
			return
		default:
			if *count > 0 && msgIdx > *count {
				log.Println("done sending messages")
				return
			}
			value := fmt.Sprintf("message-%d", msgIdx)
			msg := kafka.Message{
				Key:   []byte(strconv.Itoa(msgIdx)),
				Value: []byte(value),
			}
			if err := w.WriteMessages(ctx, msg); err != nil {
				log.Printf("write error: %v", err)
				time.Sleep(time.Second)
				continue
			}
			log.Printf("sent: %s", value)
			msgIdx++
			if *intv > 0 {
				time.Sleep(*intv)
			}
		}
	}
}

func getEnv(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}
