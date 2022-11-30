package sql

import (
	"fmt"

	"gorm.io/gorm"
)

type SQL struct {
	driver           gorm.Dialector
	activeConnection *gorm.DB
}

func (db *SQL) Auth() error {
	// Connect to the database
	if gormDb, err := gorm.Open(db.driver, &gorm.Config{}); err != nil {
		return err
	} else {
		// Success - store the connection
		db.activeConnection = gormDb
	}

	// Perform a simple query to ensure connectivity
	rows, err := db.activeConnection.Raw("SELECT 1 + 1 AS result").Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (db *SQL) Close() error {
	return nil
}

func (db *SQL) InsertBulk(table string, data []map[string]interface{}) (int, error) {
	result := db.activeConnection.Table(table).CreateInBatches(data, 100)
	if err := result.Error; err != nil {
		return 0, err
	}
	return len(data), nil
}

func (db *SQL) Truncate(table string) error {
	var sql string
	switch db.activeConnection.Name() {
	case "sqlite":
		sql = fmt.Sprintf("DELETE FROM %s", table)
	default:
		sql = fmt.Sprintf("TRUNCATE TABLE %s", table)
	}

	rows, err := db.activeConnection.Raw(sql).Rows()
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}
