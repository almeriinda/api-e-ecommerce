package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"runtime"
)

func LoadEnv() {
	var err error
	if runtime.GOOS == "linux" {
		err = godotenv.Load(".env")
	} else {
		err = godotenv.Load(".env")
	}
	if err != nil {
		log.Println(err.Error())
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
