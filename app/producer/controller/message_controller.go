package controller

import "github.com/gofiber/fiber/v2"

type MessageController interface {
	SendMessageToServer(ctx *fiber.Ctx)
}

type MessageControllerImpl struct {
}

func InitMessageController() MessageController {
	return &MessageControllerImpl{}
}

// SendMessageToServer implements MessageController.
func (*MessageControllerImpl) SendMessageToServer(ctx *fiber.Ctx) {
	panic("unimplemented")
}
