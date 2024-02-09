package main

import (
	"fmt"

	"gitlab.com/rezaif79-ri/go-rabbitmq-101/app/config"
	queuehandler "gitlab.com/rezaif79-ri/go-rabbitmq-101/app/consumer/queue_handler"
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

	// print consumed messages from queue
	forever := make(chan bool)

	queuehandler.SendMessageHandler(channel)
	queuehandler.SendMessageV2Handler(channel)

	fmt.Println("Waiting for messages...")
	<-forever
}
