package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/streadway/amqp"
	"gitlab.com/rezaif79-ri/go-rabbitmq-101/internal/rabbitdev"
)

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func main() {
	var rbmqHost = getEnv("RBMQ_HOST", "localhost")
	var rbmqPort = getEnv("RBMQ_PORT", "5672")
	var rbmqUser = getEnv("RBMQ_USER", "guest")
	var rbmqPassword = getEnv("RBMQ_PASSWORD", "guest")

	var rbmd = rabbitdev.NewRabbitMqDevConn(
		rabbitdev.WithUser(rbmqUser),
		rabbitdev.WithPassword(rbmqPassword),
		rabbitdev.WithHost(rbmqHost),
		rabbitdev.WithPort(rbmqPort),
	)

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Use(logger.New())

	apiGroup := app.Group("api")
	apiGroup.Post("messages", func(c *fiber.Ctx) error {
		type message struct {
			Message   string     `json:"message"`
			CreatedAt *time.Time `json:"created_at"`
		}
		var bodyIn message
		if err := c.BodyParser(&bodyIn); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
				"data":    nil,
			})
		}
		dtNow := time.Now()
		bodyIn.CreatedAt = &dtNow

		// Logic to handle body and publish message to rabbitmq
		rbmd.Connect()
		defer rbmd.Connection.Close()

		// opening a channel over the connection established to interact with RabbitMQ
		channel, err := rbmd.Connection.Channel()
		if err != nil {
			panic(err)
		}
		defer channel.Close()

		// declaring queue with its properties over the the channel opened
		queue, err := channel.QueueDeclare(
			"sendMessage", // name
			false,         // durable
			false,         // auto delete
			false,         // exclusive
			false,         // no wait
			nil,           // args
		)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
				"data":    nil,
			})
		}

		// Serialize bodyIn
		bodyBytes, err := json.Marshal(bodyIn)
		if err != nil {
			return c.Status(http.StatusConflict).JSON(fiber.Map{
				"status":  http.StatusConflict,
				"message": err.Error(),
				"data":    nil,
			})
		}

		// publishing a message
		err = channel.Publish(
			"",            // exchange
			"sendMessage", // key
			false,         // mandatory
			false,         // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        bodyBytes,
			},
		)
		if err != nil {
			panic(err)
		}

		fmt.Println("Queue status:", queue)
		fmt.Println("Successfully published message")

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": "OK",
			"data":    bodyIn,
		})
	})

	app.Listen(":8001")
}
