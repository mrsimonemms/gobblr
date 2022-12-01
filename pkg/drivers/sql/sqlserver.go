package sql

import (
	"fmt"

	"gorm.io/driver/sqlserver"
)

func SQLServer(database string, host string, password string, port int, user string) *SQL {
	dsn := "sqlserver://"

	if user != "" && password != "" {
		dsn += fmt.Sprintf("%s:%s@", user, password)
	}
	if host != "" {
		dsn += fmt.Sprintf("%s:%d", host, port)
	}
	if database != "" {
		dsn += fmt.Sprintf("?database=%s", database)
	}

	return &SQL{
		driver: sqlserver.Open(dsn),
	}
}
