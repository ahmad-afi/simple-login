package container

import (
	"os"
	"path/filepath"
	"regexp"
	userd "simple-login/internal/domain/user"
	"simple-login/internal/helper"
	"simple-login/internal/infrastructure/mysql"
	"simple-login/internal/usecase/user"

	"github.com/spf13/viper"
)

var v *viper.Viper

type (
	Container struct {
		Apps    *Apps
		UserUsc user.UserUsc
	}

	Apps struct {
		Name      string `mapstructure:"name"`
		Host      string `mapstructure:"host"`
		Version   string `mapstructure:"version"`
		Address   string `mapstructure:"address"`
		HttpPort  int    `mapstructure:"httpport"`
		SecretJwt string `mapstructure:"secretJwt"`
	}
)

func loadEnv() {
	projectDirName := "simple-login"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	v.SetConfigFile(string(rootPath) + `/.env`)
}

func init() {
	v = viper.New()

	v.AutomaticEnv()
	loadEnv()

	path, err := os.Executable()
	if err != nil {
		helper.Logger(helper.LoggerLevelPanic, "os.Executable panic", err)
	}

	dir := filepath.Dir(path)
	v.AddConfigPath(dir)

	if err := v.ReadInConfig(); err != nil {
		helper.Logger(helper.LoggerLevelPanic, "failed read config", err)
	}

	helper.Logger(helper.LoggerLevelInfo, "Succeed read configuration file", nil)
}

func AppsInit(v *viper.Viper) (apps Apps) {
	err := v.Unmarshal(&apps)
	if err != nil {
		helper.Logger(helper.LoggerLevelPanic, "Error when unmarshal configuration file", err)
	}
	helper.Logger(helper.LoggerLevelInfo, "Succeed when unmarshal configuration file", nil)
	return
}

func InitContainer() (cont *Container) {
	apps := AppsInit(v)
	mysqldb := mysql.DatabaseInit(v)

	userRepo := userd.NewUserDomain(mysqldb)
	userUsc := user.NewUserUsecase(userRepo)

	return &Container{
		Apps:    &apps,
		UserUsc: userUsc,
	}
}
