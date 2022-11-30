package sql

import (
	"gorm.io/driver/sqlite"
)

func SQLite(file string) *SQL {
	return &SQL{
		driver: sqlite.Open(file),
	}
}
