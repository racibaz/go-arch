package config

import (
	"strings"
	"testing"
)

func getTestConfig() Config {
	return Config{
		App: App{
			Name:     "testapp",
			Debug:    "true",
			Env:      "test",
			LogLevel: "debug",
			Version:  "1.0.0",
			Local:    "127.0.0.1",
		},
		Server: Server{
			Host: "localhost",
			Port: "8080",
		},
		DB: DB{
			Username: "user",
			Password: "pass",
			Host:     "localhost",
			Port:     "9876",
			Name:     "testdb",
		},
		TestDB: TestDB{
			Username: "user",
			Password: "pass",
			Host:     "localhost",
			Port:     "9877",
			Name:     "testdb",
		},
		Grpc: Grpc{
			Host: "localhost",
			Port: "9090",
		},
		Swagger: Swagger{
			Path: "/swagger",
			Host: "localhost",
			Port: "9000",
		},
		RabbitMQ: RabbitMQ{
			Host:            "rabbit",
			DefaultExchange: "default",
			DefaultQueue:    "queue",
			Username:        "rmquser",
			Password:        "rmqpass",
			Port:            "5151",
			TestPort:        "6262",
		},
		Jaeger: Jaeger{
			Host: "jaegerhost",
			Port: "3132",
		},
	}
}

func TestConfig_DatabaseUrl(t *testing.T) {
	cfg := getTestConfig()
	dbUrl := cfg.DatabaseConnectionString()
	if dbUrl == "" {
		t.Fatalf("DatabaseConnectionString should not be empty")
	}
	if !strings.Contains(dbUrl, cfg.TestDB.Port) {
		t.Errorf("Expected port (%s) in DatabaseConnectionString, got: %s", cfg.TestDB.Port, dbUrl)
	}
	if !strings.Contains(dbUrl, cfg.DB.Username) {
		t.Errorf("Expected username in DatabaseConnectionString, got: %s", dbUrl)
	}
}

func TestConfig_DatabaseUrl_LocalEnv(t *testing.T) {
	cfg := getTestConfig()
	cfg.App.Env = "local"
	dbUrl := cfg.DatabaseConnectionString()
	if !strings.Contains(dbUrl, cfg.App.Local) {
		t.Errorf("Expected local host (%s) in DatabaseConnectionString, got: %s", cfg.App.Local, dbUrl)
	}
}

func TestConfig_RabbitMQUrl(t *testing.T) {
	cfg := getTestConfig()
	rmqUrl := cfg.RabbitMQConnectionString()
	if rmqUrl == "" {
		t.Fatalf("RabbitMQConnectionString should not be empty")
	}
	if !strings.Contains(rmqUrl, cfg.RabbitMQ.Username) || !strings.Contains(rmqUrl, cfg.RabbitMQ.Password) {
		t.Errorf("User/pass not in RabbitMQConnectionString: %s", rmqUrl)
	}
	if !strings.Contains(rmqUrl, cfg.RabbitMQ.Port) {
		t.Errorf("Expected local host in test env, got: %s", rmqUrl)
	}
}

func TestConfig_JaegerUrl(t *testing.T) {
	cfg := getTestConfig()
	jurl := cfg.JaegerConnectionString()
	if jurl == "" {
		t.Fatalf("JaegerConnectionString should not be empty")
	}
	if !strings.Contains(jurl, cfg.Jaeger.Port) {
		t.Errorf("Port not in JaegerConnectionString: %s", jurl)
	}
	if !strings.Contains(jurl, cfg.App.Local) {
		t.Errorf("Expected local host in test env, got: %s", jurl)
	}
}

func TestConfig_GinMode(t *testing.T) {
	cfg := getTestConfig()
	mode := cfg.GinMode()
	if mode == "" {
		t.Fatal("GinMode should not be empty")
	}
	if mode != "test" {
		t.Errorf("Expected test mode, got: %s", mode)
	}
}

func TestConfig_GinMode_ProdEnv(t *testing.T) {
	cfg := getTestConfig()
	cfg.App.Env = "prod"
	mode := cfg.GinMode()
	if mode != "release" {
		t.Errorf("Expected release mode for prod env, got: %s", mode)
	}
}
