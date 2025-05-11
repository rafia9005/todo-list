package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dns := os.Getenv("DATABASE")
	fmt.Println(dns)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed connect to DB")
		return err
	}

	DB = db

	fmt.Println("Success connect to DB")
	return nil
}
