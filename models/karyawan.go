// Package models ...
package models

import (
	"time"

	"gorm.io/gorm"
)

// Karyawan Modyel
type Karyawan struct {
	gorm.Model
	ID   uint       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Name string     `gorm:"varchar(300)" json:"name"`
	DOB  *time.Time `gorm:"date" json:"dob"`
}
