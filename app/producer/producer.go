package main

import (
	"fmt"

	"github.com/streadway/amqp"
	"gitlab.com/rezaif79-ri/go-rabbitmq-101/internal/rabbitdev"
)

func main() {
	var rbmd = rabbitdev.NewRabbitMqDevConn(
		rabbitdev.WithUser("guest"),
		rabbitdev.WithPassword("guest"),
		rabbitdev.WithHost("localhost"),
		rabbitdev.WithPort("5672"),
	)

	rbmd.Connect()
	defer rbmd.Connection.Close()

	// opening a channel over the connection established to interact with RabbitMQ
	channel, err := rbmd.Connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	// declaring queue with its properties over the the channel opened
	queue, err := channel.QueueDeclare(
		"testing", // name
		false,     // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // args
	)
	if err != nil {
		panic(err)
	}

	// publishing a message
	err = channel.Publish(
		"",        // exchange
		"testing", // key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Test Message"),
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Queue status:", queue)
	fmt.Println("Successfully published message")
}
