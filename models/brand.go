package models

import (
	"api-e-commerce/config"
	"fmt"
)

type Brand struct {
	ID   int    `gorm:"primary_key" rql:"filter,sort" json:"id"`
	Name string `json:"name"`
}

func (e Brand) BrandsList() ([]Brand, error) {
	connection := config.OpenDB()
	defer config.CloseDB(connection)
	var brands []Brand
	err := connection.Table("brands").Find(&brands).Error
	if err != nil {
		fmt.Println("Erro")
	}
	return brands, err
}
