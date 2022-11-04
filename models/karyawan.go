// Package models ...
package models

import (
	"time"

	"gorm.io/gorm"
)

// Karyawan Model
type Karyawan struct {
	gorm.Model
	ID              uint        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Name            string      `gorm:"varchar(300)" json:"name"`
	DOB             *time.Time  `gorm:"date" json:"dob"`
	NoHandphone     string      `gorm:"varchar(300)" json:"no_handphone"`
	Address         string      `gorm:"varchar(max)" json:"address"`
	JoinDate        *time.Time  `gorm:"date" json:"joindate"`
	Role            string      `json:"role"`
	DepartementCode string      `json:"departementcode"`
	Departement     Departement `gorm:"foreignKey:DepartementCode;references:DepartementCode"`
}

// Departement model
type Departement struct {
	gorm.Model
	ID              uint   `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	DepartementCode string `gorm:"unique; not null"`
	NameDepartement string `json:"namedepartment"`
	Description     string `json:"description"`
}
