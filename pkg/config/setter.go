package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Set(configPath, envFilePath string) error {
	// load .env before viper
	if err := godotenv.Load(envFilePath); err != nil {
		log.Printf("Warning: unable to load .env file: %v", err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	// ENV override support
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config: %w", err)
	}

	if err := viper.Unmarshal(&configurations); err != nil {
		return fmt.Errorf("unable to decode into struct: %w", err)
	}

	return nil
}
