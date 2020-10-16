package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Token string
)

func ReadConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Print(err)
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN_DISCORD")

	if len(token) > 0 {
		Token = token
	}

	return nil
}
