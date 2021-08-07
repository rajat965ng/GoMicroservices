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

	err = channel.Publish("email", "test", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello World !!"),
	})

	if err != nil {
		fmt.Printf("Error Publishing payload in rabbitMq channel: %s\n", err.Error())
	}
}
