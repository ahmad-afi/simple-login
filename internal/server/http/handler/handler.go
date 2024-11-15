package handler

import "simple-login/internal/infrastructure/container"

type handler struct {
	UserHandler userHandler
}

func SetupHandler(cont container.Container) handler {
	return handler{
		UserHandler: NewUserHandler(cont.UserUsc),
	}
}
