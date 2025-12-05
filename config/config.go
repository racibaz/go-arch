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
	Local    string `mapstructure:"local"`
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
	Host            string `mapstructure:"host"`
	DefaultExchange string `mapstructure:"default_exchange"`
	DefaultQueue    string `mapstructure:"default_queue"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Port            string `mapstructure:"port"`
	TestPort        string `mapstructure:"test_port"`
}

func (config *Config) DatabaseUrl() string {

	port := config.DB.Port
	host := config.DB.Host

	switch config.App.Env {
	case "test":
		port = config.DB.TestPort
		host = config.DB.Host
	case "local":
		port = config.DB.Port
		host = config.App.Local
	default:
		port = config.DB.Port
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		port,
	)
}

func (config *Config) RabbitMQUrl() string {

	host := config.RabbitMQ.Host

	switch config.App.Env {
	case "test":
		host = config.App.Local
	case "local":
		host = config.App.Local
	}

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		config.RabbitMQ.Username,
		config.RabbitMQ.Password,
		host,
		config.RabbitMQ.Port,
	)

	return url
}
