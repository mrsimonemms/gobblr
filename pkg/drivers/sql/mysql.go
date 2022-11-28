package sql

import (
	"fmt"

	"gorm.io/driver/mysql"
)

func MySQL(database string, host string, password string, port int, user string) *SQL {
	var dsn string

	if user != "" && password != "" {
		dsn += fmt.Sprintf("%s:%s@", user, password)
	}
	if host != "" {
		dsn += fmt.Sprintf("tcp(%s:%d)", host, port)
	}
	dsn += fmt.Sprintf("/%s", database)

	return &SQL{
		driver: mysql.Open(dsn),
	}
}
