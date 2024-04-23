package main

import (
	"log"
	"shops/models"
	"shops/routers"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	router := routers.Setup()
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&[]models.Item{})
	if err != nil {
		log.Fatal(err)
	}
	router.Run() // instead of router.Run()
}
