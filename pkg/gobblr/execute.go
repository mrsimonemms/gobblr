package gobblr

import (
	"fmt"

	"github.com/cenkalti/backoff/v4"
	"github.com/mrsimonemms/gobblr/pkg/drivers"
)

func Execute(dataPath string, db drivers.Driver, retries uint64) (map[string]int, error) {
	return backoff.RetryWithData(
		func() (map[string]int, error) {
			return retryExecute(dataPath, db)
		},
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), retries),
	)
}

func retryExecute(dataPath string, db drivers.Driver) (map[string]int, error) {
	inserted := make(map[string]int, 0)
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
		inserted[file.TableName] = tableInserted
	}

	return inserted, err
}
