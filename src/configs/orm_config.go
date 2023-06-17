package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigList struct {
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
}

func Init() ConfigList {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config := ConfigList{
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_USERPASS"),
	}
	return Config
}
