package app

import (
	"context"
	"dependencies/entities"
	"dependencies/models"
	"dependencies/repositories"
	"dependencies/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// AddEmployee godoc
// @Summary Adds a new employee
// @Description Adds a new employee
// @Router /api/v1/employee
func (app *App) AddEmployee(c *fiber.Ctx) error {
	ctx := context.Background()

	createEmployee := &models.CreateEmployeeModel{}

	if err := c.BodyParser(&createEmployee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Invalid request"))
	}

	repo := repositories.Create[entities.Employee](app.db)
	employee := utils.CreateEmployee(createEmployee)
	err := repo.Add(&employee, ctx)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Error occurred while adding employee"))
	}
	return c.Status(fiber.StatusOK).JSON(utils.StatusOK("Employee added successfully"))
}

func (app *App) GetEmployee(c *fiber.Ctx) error {
	ctx := context.Background()

	repo := repositories.Create[entities.Employee](app.db)
	employees, err := repo.GetAll(ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Error occurred while getting employees"))
	}
	return c.Status(fiber.StatusOK).JSON(utils.StatusOK(employees))
}

// @title Employee
// @version 1.0
// @description Get Employee by Id
// @BasePath /
func (app *App) GetEmployeeById(c *fiber.Ctx) error {
	ctx := context.Background()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Invalid request"))
	}

	repo := repositories.Create[entities.Employee](app.db)
	employee, err := repo.GetById(id, ctx)
	if employee.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(utils.StatusNotFound("Employee not found"))
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Error occurred while getting employee"))
	}

	return c.Status(fiber.StatusOK).JSON(utils.StatusOK(employee))
}

func (app *App) UpdateEmployee(c *fiber.Ctx) error {
	ctx := context.Background()
	updateEmployee := &models.UpdateEmployeeModel{}
	if err := c.BodyParser(&updateEmployee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Invalid request"))
	}

	repo := repositories.Create[entities.Employee](app.db)
	employee := utils.UpdateEmployee(updateEmployee)
	err := repo.Update(&employee, ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Error occurred while updating employee"))
	}
	return c.Status(fiber.StatusOK).JSON(utils.StatusOK("Employee updated successfully"))
}

func (app *App) DeleteEmployee(c *fiber.Ctx) error {
	ctx := context.Background()
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Invalid request"))
	}
	repo := repositories.Create[entities.Employee](app.db)
	err = repo.Delete(id, ctx)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Error occurred while deleting employee"))
	}
	return c.Status(fiber.StatusOK).JSON(utils.StatusOK("Employee deleted successfully"))
}
