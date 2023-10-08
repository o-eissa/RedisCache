package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVars() (string, string, string) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	address := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")
	password := os.Getenv("PASSWORD")

	return address, port, password
}
