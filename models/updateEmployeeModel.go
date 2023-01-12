package models

type UpdateEmployeeModel struct {
	Id        uint               `json:"id,omitempty"`
	FirstName string             `json:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty"`
	Address   UpdateAddressModel `json:"adress,omitempty"`
}

type UpdateAddressModel struct {
	Id    uint   `json:"id,omitempty"`
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
