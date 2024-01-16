package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gitlab.com/rezaif79-ri/go-rabbitmq-101/app/config"
	"gitlab.com/rezaif79-ri/go-rabbitmq-101/app/producer/router"
)

func main() {
	rbmdConn := config.InitRabbitMqDevConn()
	rbmdConn.Connect()
	defer rbmdConn.Connection.Close()

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Use(logger.New())

	apiGroup := app.Group("api")
	router.SetupMainRouter(apiGroup, rbmdConn)

	app.Listen(":8001")
}
