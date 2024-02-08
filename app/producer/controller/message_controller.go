package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"gitlab.com/rezaif79-ri/go-rabbitmq-101/internal/rabbitdev"
)

type MessageController interface {
	SendMessageToServer(c *fiber.Ctx) error
	SendMessageToServerV2(c *fiber.Ctx) error
}

type MessageControllerImpl struct {
	rbmdConn *rabbitdev.RabbitMqDev
}

// SendMessageToServerV2 implements MessageController.
func (*MessageControllerImpl) SendMessageToServerV2(c *fiber.Ctx) error {
	panic("unimplemented")
}

func InitMessageController(conn *rabbitdev.RabbitMqDev) MessageController {
	return &MessageControllerImpl{
		rbmdConn: conn,
	}
}

// SendMessageToServer implements MessageController.
func (mc *MessageControllerImpl) SendMessageToServer(c *fiber.Ctx) error {
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
	mc.rbmdConn.Connect()
	defer mc.rbmdConn.Connection.Close()

	// opening a channel over the connection established to interact with RabbitMQ
	channel, err := mc.rbmdConn.Connection.Channel()
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

	unique := uuid.New()
	// publishing a message
	err = channel.Publish(
		"",            // exchange
		"sendMessage", // key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bodyBytes,
			MessageId:   unique.String(),
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
}
