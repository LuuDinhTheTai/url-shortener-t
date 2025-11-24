package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
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

func LoadConfig() *Configuration {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	cfg := &Configuration{}
	cfg.Server.Port = os.Getenv("PORT")
	cfg.Server.Host = os.Getenv("HOST")
	cfg.Database.URI = os.Getenv("MONGO_URI")
	cfg.Database.Name = os.Getenv("DB_NAME")
	cfg.Database.Username = os.Getenv("DB_USERNAME")
	cfg.Database.Password = os.Getenv("DB_PASSWORD")

	return cfg
}
