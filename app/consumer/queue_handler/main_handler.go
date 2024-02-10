package queuehandler

import "github.com/streadway/amqp"

func SetupRbmqHandler(channel *amqp.Channel) error {
	if err := SendMessageHandler(channel); err != nil {
		return err
	}

	if err := SendMessageV2Handler(channel); err != nil {
		return err
	}

	return nil
}
