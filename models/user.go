package models

import (
	"api-e-commerce/config"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                int            `gorm:"primary_key" rql:"filter,sort" json:"id"`
	Email             string         `gorm:"type:varchar" json:"email"`
	PasswordHash      string         `gorm:"type:varchar" json:"password_hash"`
	Name              string         `gorm:"type:varchar" json:"name"`
	DateOfBirth       string         `json:"date_of_birth"`
	Phone             string         `gorm:"type:varchar" json:"phone"`
	Active            bool           `json:"active"`
	CreatedAt         time.Time      `rql:"filter" json:"created_at,omitempty"`
	UpdatedAt         time.Time      `json:"updated_at,omitempty"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	ProfilePictureUrl string         `json:"profile_picture_url"`
	Address           int            `json:"address"`
}

func (e *User) CreateUser() error {
	connection := config.OpenDB()
	defer config.CloseDB(connection)
	err := connection.Table("users").Create(&e).Error
	if err != nil {
		fmt.Println("Erro")
	}
	return err
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	connection := config.OpenDB()
	defer config.CloseDB(connection)

	err := connection.Table("users").Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
