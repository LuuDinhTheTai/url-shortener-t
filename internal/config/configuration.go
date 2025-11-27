package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Server struct {
		Port string
		Host string
	}
	Database struct {
		URI      string
		Name     string
		Username string
		Password string
	}
}

type Configuration struct {
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}

func (c *Configuration) LoadEnv() (EnvConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return EnvConfig{}, err
	}

	cfg := EnvConfig{}
	cfg.Server.Port = os.Getenv("PORT")
	cfg.Server.Host = os.Getenv("HOST")
	cfg.Database.URI = os.Getenv("MONGO_URI")
	cfg.Database.Name = os.Getenv("DB_NAME")
	cfg.Database.Username = os.Getenv("DB_USERNAME")
	cfg.Database.Password = os.Getenv("DB_PASSWORD")

	return cfg, nil
}
