package app

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func CreateLogger() logger.Config {
	file, err := os.OpenFile("./logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return logger.Config{
		Output:     file,
		TimeFormat: "01:01:01 01-01-2000",
	}
}
