package models

import (
	"api-e-commerce/config"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Products struct {
	ID              int            `gorm:"primary_key" rql:"filter,sort" json:"id"`
	Name            string         `gorm:"type:varchar" json:"name"`
	Description     string         `gorm:"type:varchar" json:"description"`
	Price           int            `json:"price"`
	CategoryID      int            `json:"category_id"`
	Category        Categories     `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	BrandID         int            `json:"brand_id"`
	Brand           Brand          `gorm:"foreignKey:BrandID" json:"brand,omitempty"`
	StockQuantity   int            `json:"stock_quantity"`
	ProductImageUrl string         `json:"product_image_url"`
	UpdatedAt       time.Time      `json:"updated_at,omitempty"`
	CreatedAt       time.Time      `rql:"filter" json:"created_at,omitempty"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (e *Products) CreateProducts() error {
	connection := config.OpenDB()
	defer config.CloseDB(connection)
	err := connection.Table("products").Create(&e).Error
	if err != nil {
		fmt.Println("Erro")
	}
	return err
}

func (e Products) ProductsList() ([]Products, error) {
	connection := config.OpenDB()
	defer config.CloseDB(connection)
	var products []Products
	err := connection.Table("products").Find(&products).Error
	if err != nil {
		fmt.Println("Erro")
	}
	return products, err
}
