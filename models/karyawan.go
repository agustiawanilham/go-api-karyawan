package models

import "time"


type Karyawan struct {
	Id   int64  `gorm:"primaryKey" json:"id"`
	Name string `gorm:"varchar(300)" json:"name"`
    DOB time.Time `gorm:"datetime" json:"dob"`
}
