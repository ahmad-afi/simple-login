package http

import (
	"fmt"
	"net/http"
	"simple-login/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func SetupServerHTTP(containerConf *container.Container) {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New())

	HTTPRouteInit(app, containerConf)

	port := fmt.Sprintf("%s:%d", containerConf.Apps.Host, containerConf.Apps.HttpPort)

	if err := app.Listen(port); err != nil && err != http.ErrServerClosed {
		app.Server().Logger.Printf("shutting down the server : ", err)
	}
}
