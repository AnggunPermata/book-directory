package config

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBInstance *sql.DB
}

func LoadConfig() Config {
	godotenv.Load(".env")
	databaseURL := os.Getenv("DATABASE_URL")
	dbInstance := InitDB(databaseURL)
	port := os.Getenv("PORT")

	return Config{
		DBInstance: dbInstance,
		Port:       port,
	}
}

func LoadEnv(key string) (value string) {
	godotenv.Load(".env")
	value = os.Getenv(key)
	return
}
