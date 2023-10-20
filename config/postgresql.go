package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var db *gorm.DB

func OpenDB() *gorm.DB {
	databaseURL := os.Getenv("database")
	database, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
	}
	return database
}

func CloseDB(con *gorm.DB) {
	db, err := con.DB()
	if err != nil {
		log.Fatalln(err)
	}
	db.Close()
}
