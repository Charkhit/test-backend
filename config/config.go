package config

import (
	"fmt"
	"log"
	"os"

	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DBPassword string
	DBUrl      string
	ServerPort string
}

func LoadConfig() *Config {
	envFile := fmt.Sprintf(".env")
	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("ไม่พบไฟล์ %s จะใช้ค่า environment variable ของระบบแทน", envFile)
	}

	return &Config{
		DBPassword: getEnv("DB_PASSWORD", "xxx"),
		DBUrl:      strings.ReplaceAll(getEnv("DB_URL", "xxx"), "{DB_PASSWORD}", getEnv("DB_PASSWORD", "xxx")),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
