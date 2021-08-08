package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

const (
	brokers = "localhost:9092"
	topic   = "junk"
	group   = "junkies"
)

func lintErr(err error, reason string) {
	fmt.Printf("Error in %s: %s\n", reason, err.Error())
}

func main() {
	go producer()
	consumer()
}

func producer() {
	writer := &kafka.Writer{
		Topic:    topic,
		Addr:     kafka.TCP(brokers),
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()
	time.Sleep(2 * time.Second)
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)

	lintErr(err, "writing messages")
}

func consumer() {
	rc := &kafka.ReaderConfig{
		Brokers: []string{brokers},
		Topic:   topic,
		GroupID: group,
	}
	reader := kafka.NewReader(*rc)

	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		lintErr(err, "Reading Message")
		fmt.Printf("key: %s, val:%s\n", msg.Key, msg.Value)
	}
}
