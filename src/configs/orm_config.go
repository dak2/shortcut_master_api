package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigList struct {
	DbHost     string
	DbName     string
	DbUser     string
	DbPassword string
}

func Init() ConfigList {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config := ConfigList{
		DbHost:     os.Getenv("DBHOST"),
		DbName:     os.Getenv("DBNAME"),
		DbUser:     os.Getenv("DBUSER"),
		DbPassword: os.Getenv("ROOTPASS"),
	}
	return Config
}
