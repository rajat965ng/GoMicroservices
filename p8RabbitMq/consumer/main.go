package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("Error connecting rabbitMq producer: %s\n", err.Error())
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		fmt.Printf("Error creating rabbitMq channel: %s\n", err.Error())
	}

	delivery, err := channel.Consume("office", "consumerA", true, false, false, false, nil)

	if err != nil {
		fmt.Printf("Error consulming rabbitMq queue: %s\n", err.Error())
	}

	for msg := range delivery {
		fmt.Printf("Message is: %s\n", msg.Body)
	}

}
