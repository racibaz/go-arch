package config

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Config struct {
	App      App      `mapstructure:"app"`
	Server   Server   `mapstructure:"server"`
	DB       DB       `mapstructure:"db"`
	Grpc     Grpc     `mapstructure:"grpc"`
	Swagger  Swagger  `mapstructure:"swagger"`
	RabbitMQ RabbitMQ `mapstructure:"rabbitmq"`
	Jaeger   Jaeger   `mapstructure:"jaeger"`
	TestDB   TestDB   `mapstructure:"test_db"`
}

type App struct {
	Name             string `mapstructure:"name"`
	Debug            string `mapstructure:"debug"`
	Env              string `mapstructure:"env"`
	LogLevel         string `mapstructure:"log_level"`
	JWTKey           string `mapstructure:"jwt_key"`
	JWTWebTimeout    int    `mapstructure:"jwt_web_timeout"`
	JWTMobileTimeout int    `mapstructure:"jwt_mobile_timeout"`
	Version          string `mapstructure:"version"`
	Local            string `mapstructure:"local"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type DB struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
}

type TestDB struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
}

type Swagger struct {
	Path string `mapstructure:"path"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Grpc struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type RabbitMQ struct {
	Host            string `mapstructure:"host"`
	DefaultExchange string `mapstructure:"default_exchange"`
	DefaultQueue    string `mapstructure:"default_queue"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Port            string `mapstructure:"port"`
	TestPort        string `mapstructure:"test_port"`
}

type Jaeger struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

// DatabaseConnectionString constructs the database connection URL based on the environment.
func (config *Config) DatabaseConnectionString() string {
	port := config.DB.Port
	host := config.DB.Host

	switch config.App.Env {
	case "test":
		port = config.TestDB.Port
		host = config.App.Local
	case "local":
		port = config.DB.Port
		host = config.App.Local
	default:
		port = config.DB.Port
	}

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		port,
	)
}

// RabbitMQConnectionString constructs the RabbitMQ connection URL based on the environment.
func (config *Config) RabbitMQConnectionString() string {
	host := config.RabbitMQ.Host

	switch config.App.Env {
	case "test":
		host = config.App.Local
	case "local":
		host = config.App.Local
	}

	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		config.RabbitMQ.Username,
		config.RabbitMQ.Password,
		host,
		config.RabbitMQ.Port,
	)

	return dsn
}

// JaegerConnectionString constructs the Jaeger connection URL based on the environment.
func (config *Config) JaegerConnectionString() string {
	host := config.Jaeger.Host

	switch config.App.Env {
	case "test":
		host = config.App.Local
	case "local":
		host = config.App.Local
	}

	dsn := fmt.Sprintf("%s:%s",
		host,
		config.Jaeger.Port,
	)

	return dsn
}

// GinMode returns the appropriate Gin mode based on the application environment.
func (config *Config) GinMode() string {
	mode := gin.DebugMode

	switch config.App.Env {
	case "test":
		mode = gin.TestMode
	case "local":
		mode = gin.DebugMode
	case "prod":
		mode = gin.ReleaseMode
	case "debug":
		mode = gin.DebugMode
	default:
		mode = gin.DebugMode
	}

	return mode
}

func (config *Config) IsProduction() bool {
	if "prod" == config.App.Env {
		return true
	}
	return false
}
