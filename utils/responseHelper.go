package utils

import (
	"dependencies/models"

	"github.com/gofiber/fiber/v2"
)

func StatusOK(data interface{}) models.ResponseModel {
	return models.ResponseModel{
		Status:  fiber.StatusOK,
		Data:    data,
		Message: "OK",
		Success: true,
	}
}

func StatusFail(message string) models.ResponseModel {
	return models.ResponseModel{
		Status:  fiber.StatusBadRequest,
		Message: message,
		Success: false,
	}
}

func StatusUnauthorized(message string) models.ResponseModel {
	return models.ResponseModel{
		Status:  fiber.StatusUnauthorized,
		Message: message,
		Success: false,
	}
}

func StatusNotFound(message string) models.ResponseModel {
	return models.ResponseModel{
		Status:  fiber.StatusNotFound,
		Message: message,
		Success: true,
	}
}
