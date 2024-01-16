package main

import (
	"fmt"
	"os"

	"gitlab.com/rezaif79-ri/go-rabbitmq-101/app/consumer/config"
)

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func main() {
	rbmdConn := config.InitRabbitMqDevConn()
	rbmdConn.Connect()
	defer rbmdConn.Connection.Close()

	// opening a channel over the connection established to interact with RabbitMQ
	channel, err := rbmdConn.Connection.Channel()
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
