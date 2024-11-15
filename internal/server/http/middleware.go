package http

import (
	"encoding/json"
	"net/http"
	"simple-login/internal/helper"

	"github.com/gofiber/fiber/v2"
)

type simpleHeader struct {
	Name      string `reqHeader:"name"`
	Pass      string `reqHeader:"pass"`
	ID        string `reqHeader:"id"`
	Timestamp string `reqHeader:"timestamp"`
}

func middlewareGetHeader(ctx *fiber.Ctx) error {
	header := new(simpleHeader)
	if err := ctx.ReqHeaderParser(header); err != nil {
		return helper.BuildResponse(ctx, false, err.Error(), nil, http.StatusUnauthorized)
	}

	ress, err := json.MarshalIndent(header, "", " ")
	if err != nil {
		helper.Logger(helper.LoggerLevelError, "error", err)
	}
	helper.Logger(helper.LoggerLevelInfo, string(ress), nil)

	return ctx.Next()
}
