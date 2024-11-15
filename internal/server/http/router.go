package http

import (
	"simple-login/internal/infrastructure/container"
	"simple-login/internal/server/http/handler"

	"github.com/gofiber/fiber/v2"
)

func HTTPRouteInit(r *fiber.App, containerConf *container.Container) {
	api := r.Group("/api/v1")
	h := handler.SetupHandler(*containerConf)

	userGr := api.Group("/user")
	{
		userGr.Post("", middlewareGetHeader, h.UserHandler.CreateUser)
	}
}
