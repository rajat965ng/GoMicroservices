package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"p10RedisPubSub/domain"
	"time"
)

func Subscriber(client *redis.Client) {
	fmt.Println("::::::Init Subscriber:::::")
	ctx := context.Background()

	time.Sleep(3 * time.Second)
	topic := client.Subscribe(ctx, "user_channel")
	defer topic.Close()

	if _, err := topic.Receive(ctx); err != nil {
		fmt.Printf("failed to receive from control PubSub", err.Error())
		return
	}

	channel := topic.Channel()
	for msg := range channel {
		usr := &domain.User{}
		usr.UnMarshal([]byte(msg.Payload))
		fmt.Printf("Subscribe: %s\n", usr.Marshal())
	}

	fmt.Println("::::::Stop Subscriber:::::")
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	defer client.Close()

	ch := make(chan bool, 1)
	go Subscriber(client)

	<-ch
}
