package main

import (
	"fmt"

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
		"sendMessage", // name
		false,         // durable
		false,         // auto delete
		false,         // exclusive
		false,         // no wait
		nil,           // args
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("check queue: ", queue)

	// declaring consumer with its properties over channel opened
	msgs, err := channel.Consume(
		"sendMessage", // queue
		"",            // consumer
		true,          // auto ack
		false,         // exclusive
		false,         // no local
		false,         // no wait
		nil,           //args
	)
	if err != nil {
		panic(err)
	}

	// print consumed messages from queue
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}

// func sendMessageConsumer(*rabbitdev.RabbitMqDev)
