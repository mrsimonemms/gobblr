/*
 * Copyright 2022 Simon Emms <simon@simonemms.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
	if gormDB, err := gorm.Open(db.driver, &gorm.Config{}); err != nil {
		return err
	} else {
		// Success - store the connection
		db.activeConnection = gormDB
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

func (db *SQL) DriverName() string {
	return db.driver.Name()
}

func (db *SQL) InsertBulk(table string, data []map[string]any) (int, error) {
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
