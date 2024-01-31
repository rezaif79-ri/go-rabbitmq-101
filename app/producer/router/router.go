package router

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/rezaif79-ri/go-rabbitmq-101/internal/rabbitdev"
)

func SetupMainRouter(router fiber.Router, rbmdConn *rabbitdev.RabbitMqDev) {
	messageRoute := router.Group("messages")
	messageRoute.Post("messages", func(c *fiber.Ctx) error {
		return nil
	})
}
