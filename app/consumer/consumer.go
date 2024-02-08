package main

import (
	"fmt"

	"gitlab.com/rezaif79-ri/go-rabbitmq-101/app/config"
)

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

	// declaring queue with its properties over the the channel opened
	queue, err = channel.QueueDeclare(
		"sendMessageV2", // name
		false,           // durable
		false,           // auto delete
		false,           // exclusive
		false,           // no wait
		nil,             // args
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("check queue v2: ", queue)

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

	// declaring consumer with its properties over channel opened
	msgsV2, err := channel.Consume(
		"sendMessageV2", // queue
		"",              // consumer
		true,            // auto ack
		false,           // exclusive
		false,           // no local
		false,           // no wait
		nil,             //args
	)
	if err != nil {
		panic(err)
	}

	// print consumed messages from queue
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Println("cons tag:", msg.ConsumerTag)
			fmt.Println("deliv tag:", msg.DeliveryTag)
			fmt.Println("message id:", msg.MessageId)
			fmt.Printf("Received Message: %s\n", msg.Body)
		}

		for msg := range msgsV2 {
			fmt.Println("V2 - cons tag:", msg.ConsumerTag)
			fmt.Println("V2 - deliv tag:", msg.DeliveryTag)
			fmt.Println("V2 - message id:", msg.MessageId)
			fmt.Printf("V2 - Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}
