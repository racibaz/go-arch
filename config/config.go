package config

import "fmt"

type Config struct {
	App      App
	Server   Server
	DB       DB
	Grpc     Grpc
	Swagger  Swagger
	RabbitMQ RabbitMQ
}

type App struct {
	Name     string
	Debug    string
	Env      string
	LogLevel string
	Version  string
}

type Server struct {
	Host string
	Port string
}

type DB struct {
	Username string
	Password string
	Host     string
	Port     string
	TestPort string
	Name     string
}

type Swagger struct {
	Path string
	Host string
	Port string
}

type Grpc struct {
	Host string
	Port string
}

type RabbitMQ struct {
	Url string
}

func (config *Config) DatabaseUrl() string {

	var dbPort = config.DB.Port

	switch config.App.Env {
	case "test":
		dbPort = config.DB.TestPort
	case "dev", "local":
		dbPort = config.DB.Port
	case "prod":
		dbPort = config.DB.Port
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DB.Host,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		dbPort,
	)
}
