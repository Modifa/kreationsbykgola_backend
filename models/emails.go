package models

type Order struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	Quantity    string `json:"quantity"`
	Flavour     string `json:"flavor"`
	Product     string `json:"product"`
	Address     string `json:"address"`
}
