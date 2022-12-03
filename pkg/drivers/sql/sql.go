package sql

import (
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

func (db *SQL) DriverName() string {
	return db.driver.Name()
}

func (db *SQL) InsertBulk(table string, data []map[string]interface{}) (int, error) {
	result := db.activeConnection.Table(table).CreateInBatches(data, 100)
	if err := result.Error; err != nil {
		return 0, err
	}
	return len(data), nil
}

func (db *SQL) Truncate(table string) error {
	// Gorm doesn't have a truncate method - this will not reset the "id" field
	result := db.activeConnection.
		Session(&gorm.Session{AllowGlobalUpdate: true}).
		Table(table).
		Delete("")

	return result.Error
}
