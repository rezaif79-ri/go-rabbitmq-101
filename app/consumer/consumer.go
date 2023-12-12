package main

import (
	"fmt"
	"os"

	"gitlab.com/rezaif79-ri/go-rabbitmq-101/internal/rabbitdev"
)

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func main() {
	var rbmqHost = getEnv("RBMQ_HOST", "localhost")
	var rbmqPort = getEnv("RBMQ_PORT", "5672")
	var rbmqUser = getEnv("RBMQ_USER", "guest")
	var rbmqPassword = getEnv("RBMQ_PASSWORD", "guest")

	var rbmd = rabbitdev.NewRabbitMqDevConn(
		rabbitdev.WithUser(rbmqUser),
		rabbitdev.WithPassword(rbmqPassword),
		rabbitdev.WithHost(rbmqHost),
		rabbitdev.WithPort(rbmqPort),
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
