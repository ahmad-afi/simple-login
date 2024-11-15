package handler

import (
	"net/http"
	"simple-login/internal/helper"
	"simple-login/internal/usecase/user"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userUsc user.UserUsc
}

func NewUserHandler(userUsc user.UserUsc) userHandler {
	return userHandler{userUsc: userUsc}
}

func (h *userHandler) CreateUser(ctx *fiber.Ctx) error {
	data := new(user.CreateUserReq)
	if err := ctx.BodyParser(data); err != nil {
		return helper.BuildResponse(ctx, false, err.Error(), nil, http.StatusBadRequest)
	}

	err := h.userUsc.CreateUser(ctx.Context(), *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEED, helper.SUCCEED, http.StatusOK)
}
