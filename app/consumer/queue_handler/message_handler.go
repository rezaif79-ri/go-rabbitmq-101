package queuehandler

import (
	"fmt"

	"github.com/streadway/amqp"
)

func SendMessageHandler(rbmqChan *amqp.Channel) error {
	// declaring queue with its properties over the the channel opened
	_, err := rbmqChan.QueueDeclare(
		"sendMessage", // name
		false,         // durable
		false,         // auto delete
		false,         // exclusive
		false,         // no wait
		nil,           // args
	)
	if err != nil {
		return err
	}

	// declaring consumer with its properties over channel opened
	msgs, err := rbmqChan.Consume(
		"sendMessage", // queue
		"",            // consumer
		true,          // auto ack
		false,         // exclusive
		false,         // no local
		false,         // no wait
		nil,           //args
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			fmt.Println("cons tag:", msg.ConsumerTag)
			fmt.Println("deliv tag:", msg.DeliveryTag)
			fmt.Println("message id:", msg.MessageId)
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	return nil
}

func SendMessageV2Handler(rbmqChan *amqp.Channel) error {
	// declaring queue with its properties over the the channel opened
	_, err := rbmqChan.QueueDeclare(
		"sendMessageV2", // name
		false,           // durable
		false,           // auto delete
		false,           // exclusive
		false,           // no wait
		nil,             // args
	)
	if err != nil {
		return err
	}

	// declaring consumer with its properties over channel opened
	msgs, err := rbmqChan.Consume(
		"sendMessageV2", // queue
		"",              // consumer
		true,            // auto ack
		false,           // exclusive
		false,           // no local
		false,           // no wait
		nil,             //args
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			fmt.Println("V2 - cons tag:", msg.ConsumerTag)
			fmt.Println("V2 - deliv tag:", msg.DeliveryTag)
			fmt.Println("V2 - message id:", msg.MessageId)
			fmt.Printf("V2 - Received Message: %s\n", msg.Body)
		}
	}()

	return nil
}
