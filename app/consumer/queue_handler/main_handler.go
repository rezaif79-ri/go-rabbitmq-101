package queuehandler

import "github.com/streadway/amqp"

func SetupRbmqHandler(channel *amqp.Channel) error {
	if err := sendMessageHandler(channel); err != nil {
		return err
	}

	if err := sendMessageV2Handler(channel); err != nil {
		return err
	}

	return nil
}
