package helper

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	SUCCEED = "Success"
)

func BuildResponse(ctx *fiber.Ctx, status bool, message string, data interface{}, code int) error {

	return ctx.Status(code).JSON(&Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
