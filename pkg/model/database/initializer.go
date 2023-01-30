package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := os.Getenv("DATABASE_DSN")
	fmt.Println(dsn)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	db = database

	autoMigrate()
}

func GetDatabase() *gorm.DB {
	return db
}
