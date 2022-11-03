package db

import (
	"fmt"
	"log"

	"github.com/agustiawanilham/go-api-karyawan/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("karyawan.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s", err))
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Karyawan{})

	log.Println("Database Migration Completed!")
	DB = db
}
