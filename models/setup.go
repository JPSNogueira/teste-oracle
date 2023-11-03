// models/setup.go

package models

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "sqlserver://sa:sql123@localhost:1433?database=racoesoltp"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&Customer{})
	if err != nil {
		return
	}

	DB = db
}
