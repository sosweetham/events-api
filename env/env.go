package env

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}	

func GetEnv(key string, nullable bool) string {
	envVar := os.Getenv(key)
	if envVar == "" && !nullable {
		panic("Environment variable " + key + " is required")
	}
	return envVar
}