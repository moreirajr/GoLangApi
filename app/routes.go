package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"dependencies/docs"
)

func (app *App) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("healthy")
}

func ConfigureSwagger(fApp *fiber.App) {
	docs.SwaggerInfo.Title = "GoLang Example API"
	docs.SwaggerInfo.Description = "A simple GoLang Api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	fApp.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/swagger/doc.json",
	}))
}

func (app *App) Routes(fApp *fiber.App) {

	ConfigureSwagger(fApp)
	fApp.Get("/health", app.HealthCheck)

	apiV1 := fApp.Group("/api/v1")

	apiV1.Get("/employee/:id", app.GetEmployeeById)
	apiV1.Get("/employee", app.GetEmployee)
	apiV1.Post("/employee", app.AddEmployee)
	apiV1.Put("/employee", app.UpdateEmployee)
	apiV1.Delete("/employee", app.DeleteEmployee)
}
