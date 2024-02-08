package router

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/rezaif79-ri/go-rabbitmq-101/app/producer/controller"
	"gitlab.com/rezaif79-ri/go-rabbitmq-101/internal/rabbitdev"
)

func SetupMainRouter(router fiber.Router, rbmdConn *rabbitdev.RabbitMqDev) {
	messageRoute := router.Group("messages")

	messageController := controller.InitMessageController(rbmdConn)

	messageRoute.Post("", messageController.SendMessageToServer)
	messageRoute.Post("v2", messageController.SendMessageToServerV2)
}
