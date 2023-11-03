package models

type Customer struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Address string `json:"address"`
	City	string `json:"city"`
}