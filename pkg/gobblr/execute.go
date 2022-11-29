package gobblr

import (
	"fmt"

	"github.com/mrsimonemms/gobblr/pkg/drivers"
)

type Inserted struct {
	Table string `json:"table"`
	Count int    `json:"count"`
}

func Execute(dataPath string, db drivers.Driver) ([]Inserted, error) {
	inserted := make([]Inserted, 0)
	var err error

	// Connect to database
	if err := db.Auth(); err != nil {
		return nil, err
	}

	// We've finished - let's clear up after ourselves
	defer func() {
		err = db.Close()
	}()

	// Find the files
	files, err := FindFiles(dataPath)
	if err != nil {
		return nil, err
	}

	// Iterate over each file, delete and then ingest data
	for _, file := range files {
		fileData, err := file.LoadFile()
		if err != nil {
			return nil, err
		}

		if len(fileData.Data) == 0 {
			return inserted, fmt.Errorf("no data in file: %s", file.Path)
		}

		// Clear out any existing data
		if err := db.Truncate(file.TableName); err != nil {
			return nil, err
		}

		// Insert the data to the table
		tableInserted, err := db.InsertBulk(file.TableName, fileData.Data)
		if err != nil {
			// Insertion failed - truncate
			if err := db.Truncate(file.TableName); err != nil {
				// Failed to truncate
				return nil, err
			}
			return nil, err
		}

		// Store the result for output
		inserted = append(inserted, Inserted{
			Table: file.TableName,
			Count: tableInserted,
		})
	}

	return inserted, err
}
