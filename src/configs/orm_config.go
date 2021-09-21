package config

import (
	"os"
)

type ConfigList struct {
	DbHost     string
	DbName     string
	DbUser     string
	DbPassword string
}

var Config ConfigList

func Init() {

	Config = ConfigList{
		DbHost:     os.Getenv("DBHOST"),
		DbName:     os.Getenv("DBNAME"),
		DbUser:     os.Getenv("DBUSER"),
		DbPassword: os.Getenv("ROOTPASS"),
	}
}
