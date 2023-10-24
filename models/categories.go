package models

import (
	"api-e-commerce/config"
	"fmt"
)

type Categories struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (e Categories) CategoriesList() ([]Categories, error) {
	connection := config.OpenDB()
	defer config.CloseDB(connection)
	var categories []Categories
	err := connection.Table("categories").Find(&categories).Error
	if err != nil {
		fmt.Println("Erro")
	}
	return categories, err
}
