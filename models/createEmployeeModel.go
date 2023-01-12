package models

type CreateEmployeeModel struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
}
