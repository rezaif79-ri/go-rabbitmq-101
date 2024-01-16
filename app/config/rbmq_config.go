package config

import (
	envutil "gitlab.com/rezaif79-ri/go-rabbitmq-101/internal/env_util"
	"gitlab.com/rezaif79-ri/go-rabbitmq-101/internal/rabbitdev"
)

func InitRabbitMqDevConn() *rabbitdev.RabbitMqDev {
	var rbmqHost = envutil.GetEenv("RBMQ_HOST", "localhost")
	var rbmqPort = envutil.GetEenv("RBMQ_PORT", "5672")
	var rbmqUser = envutil.GetEenv("RBMQ_USER", "guest")
	var rbmqPassword = envutil.GetEenv("RBMQ_PASSWORD", "guest")

	return rabbitdev.NewRabbitMqDevConn(
		rabbitdev.WithUser(rbmqUser),
		rabbitdev.WithPassword(rbmqPassword),
		rabbitdev.WithHost(rbmqHost),
		rabbitdev.WithPort(rbmqPort),
	)
}
