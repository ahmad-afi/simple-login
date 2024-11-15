package mysql

import (
	"fmt"
	"simple-login/internal/helper"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConf struct {
	Username           string `mapstructure:"mysql_username"`
	Password           string `mapstructure:"mysql_password"`
	DbName             string `mapstructure:"mysql_Dbname"`
	Host               string `mapstructure:"mysql_host"`
	Port               int    `mapstructure:"mysql_port"`
	Schema             string `mapstructure:"mysql_schema"`
	LogMode            bool   `mapstructure:"mysql_logMode"`
	MaxLifetime        int    `mapstructure:"mysql_maxLifetime"`
	MinIdleConnections int    `mapstructure:"mysql_minIdleConnections"`
	MaxOpenConnections int    `mapstructure:"mysql_maxOpenConnections"`
}

func DatabaseInit(v *viper.Viper) *sqlx.DB {
	var mysqlConfig MysqlConf
	err := v.Unmarshal(&mysqlConfig)
	if err != nil {
		helper.Logger(helper.LoggerLevelPanic, "failed init database mysql", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DbName)

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		helper.Logger(helper.LoggerLevelPanic, "Cannot conenct to database", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(mysqlConfig.MinIdleConnections)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(mysqlConfig.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	maxLifeTime := time.Duration(mysqlConfig.MaxLifetime) * time.Second
	db.SetConnMaxLifetime(maxLifeTime)

	if err := db.Ping(); err != nil {
		helper.Logger(helper.LoggerLevelError, "⇨ MySQL status is disconnected", err)
	}

	return db
}
