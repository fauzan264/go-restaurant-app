package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB(dbAddress string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	SeedDB(dbAddress)
	return db
}