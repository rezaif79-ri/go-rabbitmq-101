package config

import (
	envutil "gitlab.com/rezaif79-ri/go-rabbitmq-101/internal/env_util"
	"gitlab.com/rezaif79-ri/go-rabbitmq-101/internal/rabbitdev"
)

func InitRabbitMqDevConn() *rabbitdev.RabbitMqDev {
	var rbmqHost = envutil.GetEnv("RBMQ_HOST", "localhost")
	var rbmqPort = envutil.GetEnv("RBMQ_PORT", "5672")
	var rbmqUser = envutil.GetEnv("RBMQ_USER", "guest")
	var rbmqPassword = envutil.GetEnv("RBMQ_PASSWORD", "guest")

	return rabbitdev.NewRabbitMqDevConn(
		rabbitdev.WithUser(rbmqUser),
		rabbitdev.WithPassword(rbmqPassword),
		rabbitdev.WithHost(rbmqHost),
		rabbitdev.WithPort(rbmqPort),
	)
}
