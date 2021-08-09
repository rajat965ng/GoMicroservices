package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"p10RedisPubSub/domain"
	"time"
)

func Publisher(client *redis.Client) {
	fmt.Println("::::::Init Publisher:::::")
	for {
		time.Sleep(2 * time.Second)
		i := rand.Intn(100)
		user := &domain.User{
			Name: fmt.Sprintf("TestUser%q", i),
			Age:  i,
		}

		fmt.Printf("Publish: %s\n", user)
		client.Publish(context.Background(), "user_channel", user.Marshal())
	}
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	defer client.Close()

	ch := make(chan bool, 1)
	go Publisher(client)

	<-ch
}
