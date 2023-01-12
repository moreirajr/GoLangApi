package entities

type Employee struct {
	Id        uint     `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	AddressId uint     `json:"-"`
	Address   *Address `json:"address,omitempty" gorm:"foreignKey:AddressId;references:Id"`
}

type Address struct {
	Id    uint   `json:"id,omitempty"`
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
