package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr                   string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = envHandler()

func envHandler() *Config {
	godotenv.Load()

	return &Config{
		Addr:                   getEnv("PORT", "8080"),
		Port:                   getEnv("PORT", "8080"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", "mypassword"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                 getEnv("DB_NAME", "ecom"),
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
	}
}

func getEnv(key, fallback string) string {
	if value, err := os.LookupEnv(key); err {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, err := os.LookupEnv(key); err {
		i, err := strconv.ParseInt(value, 10, 64)

		if err != nil {
			return fallback
		}
		return i
	}

	return fallback
}
