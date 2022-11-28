package sql

import (
	"gorm.io/gorm"
)

type SQL struct {
	driver           gorm.Dialector
	activeConnection *gorm.DB
}

func (db *SQL) Auth() (err error) {
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
	defer func() {
		err = rows.Close()
	}()

	return err
}

func (db *SQL) Close() error {
	return nil
}

func (db *SQL) InsertBulk(table string, data []map[string]interface{}) (int, error) {
	return len(data), nil
}

func (db *SQL) Truncate(table string) error {
	return nil
}
