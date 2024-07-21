package env

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DBName string
	JWTSecret string
}

var AppEnv *Env

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	AppEnv = &Env{
		DBName: getEnv("DB_NAME", false),
		JWTSecret: getEnv("JWT_SECRET", false),
	}
}	

func getEnv(key string, nullable bool) string {
	envVar := os.Getenv(key)
	if envVar == "" && !nullable {
		panic("Environment variable " + key + " is required")
	}
	return envVar
}