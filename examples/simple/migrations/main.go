// The SQL databases need to have the tables created before
// we can do anything. Normally, this would be your own migration
// package's responsibility, but this simulates it for our
// purposes

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
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

type Item struct {
	gorm.Model
	Item     int
	SomeDate time.Time
}

func execute(dialector gorm.Dialector) error {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(&User{}, &Item{}); err != nil {
		return err
	}

	return nil
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
	case "sqlite":
		dialector = sqlite.Open(connection)
	case "sqlserver":
		dialector = sqlserver.Open(connection)
	default:
		panic(fmt.Errorf("unknown db type: %s", dbType))
	}

	tries := 0

	err := backoff.Retry(
		func() error {
			if tries > 0 {
				fmt.Println("Retrying database connection...")
			}
			err := execute(dialector)

			tries += 1

			return err
		},
		backoff.NewExponentialBackOff(),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Migration completed")
}
