package utils

import (
	"dependencies/entities"
	"dependencies/models"
)

func CreateEmployee(model *models.CreateEmployeeModel) entities.Employee {
	return entities.Employee{
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Address: &entities.Address{
			City:  model.City,
			State: model.State,
		},
	}
}

func UpdateEmployee(model *models.UpdateEmployeeModel) entities.Employee {
	return entities.Employee{
		Id:        model.Id,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Address: &entities.Address{
			Id:    model.Address.Id,
			City:  model.Address.City,
			State: model.Address.State,
		},
	}
}
