package sql

import (
	"fmt"
	"strings"

	"gorm.io/driver/postgres"
)

func PostgreSQL(database string, host string, password string, port int, user string) *SQL {
	var dsn []string

	if user != "" {
		dsn = append(dsn, fmt.Sprintf("user=%s", user))
	}
	if password != "" {
		dsn = append(dsn, fmt.Sprintf("password=%s", password))
	}
	if host != "" {
		dsn = append(dsn, fmt.Sprintf("host=%s", host))
	}
	dsn = append(dsn, fmt.Sprintf("port=%d", port))
	dsn = append(dsn, fmt.Sprintf("dbname=%s", database))

	return &SQL{
		driver: postgres.Open(strings.Join(dsn, " ")),
	}
}
