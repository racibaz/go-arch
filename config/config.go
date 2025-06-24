package config

type Config struct {
	App     App
	Server  Server
	DB      DB
	Grpc    Grpc
	Swagger Swagger
}

type App struct {
	Name string
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
