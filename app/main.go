package main

import (
	"simple-login/internal/infrastructure/container"
	rest "simple-login/internal/server/http"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	containerConf := container.InitContainer()
	rest.SetupServerHTTP(containerConf)
}
