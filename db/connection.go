// Package db Handling connection to database
package db

import (
	"fmt"
	"log"

	"github.com/agustiawanilham/go-api-karyawan/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB ...
var DB *gorm.DB

// ConnectDatabase Handling Connecting to database and Migration
func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("karyawan.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s", err))
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Karyawan{})
	db.AutoMigrate(&models.Departement{})

	log.Println("Database Migration Completed!")
	DB = db
}
