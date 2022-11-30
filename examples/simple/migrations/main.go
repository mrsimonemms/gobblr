package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Registered   bool
	Confirmed    bool
	Username     string
	EmailAddress string
	Name         string
	Password     string
}

func main() {
	connection := os.Getenv("DB_CONNECTION")
	dbType := os.Getenv("DB_TYPE")

	var dialector gorm.Dialector
	switch dbType {
	case "mysql":
		dialector = mysql.Open(connection)
	case "postgres":
		dialector = postgres.Open(connection)
	default:
		panic(fmt.Errorf("unknown db type: %s", dbType))
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	fmt.Println("Migration completed")
}
