package main

import (
	"fmt"

	altqueuehandler "gitlab.com/rezaif79-ri/go-rabbitmq-101/app/alt_consumer/alt_queue_handler"
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

	// print consumed messages from queue
	forever := make(chan bool)

	if err := altqueuehandler.SetupRbmqHandler(channel); err != nil {
		panic(err)
	}

	fmt.Println("Waiting for messages...")
	<-forever
}
