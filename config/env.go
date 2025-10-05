package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost           string
	Port                 string
	DBUser               string
	DBPassword           string
	DBAddress            string
	DBName               string
	DBNet                string
	JWTExpirationSeconds int
	JWTSecret            string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load("/home/dmin/go/ecom/.env") //自动加载根目录下的 .env 文件，把里面的环境变量读到程序的环境变量中
	if err != nil {
		log.Printf("未能加载 .env 文件: %v", err)
	}
	return Config{
		PublicHost:           getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                 getEnv("PORT", "8080"),
		DBUser:               getEnv("DB_USER", "root"),
		DBPassword:           getEnv("DB_PASSWORD", "password"),
		DBAddress:            fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
		DBName:               getEnv("DB_NAME", "ecom"),
		DBNet:                getEnv("DB_NET", "tcp"),
		JWTExpirationSeconds: getEnvInt("JWT_EXP", 3600*24*7), // 24 hours
		JWTSecret:            getEnv("JWT_SECRET", "your_jwt_secret_key"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		var intValue int
		_, err := fmt.Sscanf(value, "%d", &intValue)
		if err == nil {
			return intValue
		}
	}
	return fallback
}
