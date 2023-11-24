package main

import (
	"fmt"

	"gitlab.com/rezaif79-ri/go-rabbitmq-101/internal/rabbitdev"
)

func main() {
	fmt.Println("RabbitMQ in Golang: Getting started tutorial")

	var rbmd = rabbitdev.NewRabbitMqDevConn(
		rabbitdev.WithUser("guest"),
		rabbitdev.WithPassword("guest"),
		rabbitdev.WithHost("localhost"),
		rabbitdev.WithPort("5672"),
	)

	rbmd.Connect()
	defer rbmd.Connection.Close()

	fmt.Println("Successfully connected to RabbitMQ instance")

}
