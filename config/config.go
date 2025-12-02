package config

import "fmt"

type Config struct {
	App      App      `mapstructure:"app"`
	Server   Server   `mapstructure:"server"`
	DB       DB       `mapstructure:"db"`
	Grpc     Grpc     `mapstructure:"grpc"`
	Swagger  Swagger  `mapstructure:"swagger"`
	RabbitMQ RabbitMQ `mapstructure:"rabbitmq"`
}

type App struct {
	Name     string `mapstructure:"name"`
	Debug    string `mapstructure:"debug"`
	Env      string `mapstructure:"env"`
	LogLevel string `mapstructure:"log_level"`
	Version  string `mapstructure:"version"`
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
	TestPort string `mapstructure:"test_port"`
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
	DefaultExchange string `mapstructure:"default_exchange"`
	DefaultQueue    string `mapstructure:"default_queue"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Port            string `mapstructure:"port"`
	TestPort        string `mapstructure:"test_port"`
	DockerHost      string `mapstructure:"docker_host"`
	LocalHost       string `mapstructure:"local_host"`
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

func (config *Config) RabbitMQUrl() string {

	var host, port string = "", ""

	switch config.App.Env {
	case "test":
		port = config.RabbitMQ.TestPort
		host = config.RabbitMQ.DockerHost
	case "local":
		port = config.RabbitMQ.TestPort
		host = config.RabbitMQ.LocalHost
	case "dev":
		port = config.RabbitMQ.TestPort
		host = config.RabbitMQ.DockerHost
	case "prod":
		port = config.RabbitMQ.Port
		host = config.RabbitMQ.DockerHost
	}

	return fmt.Sprintf("amqp://%s:%s@%s:%s/",
		config.RabbitMQ.Username,
		config.RabbitMQ.Password,
		host,
		port,
	)
}
