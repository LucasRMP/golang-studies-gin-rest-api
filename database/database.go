package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connectionString := "host=host.docker.internal port=5432 user=postgres password=root dbname=students sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	log.Println("Database connected")
}
