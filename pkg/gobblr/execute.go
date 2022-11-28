package gobblr

import (
	"fmt"

	"github.com/mrsimonemms/gobblr/pkg/drivers"
)

func Execute(dataPath string, db drivers.Driver) (map[string]int, error) {
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
	}

	return inserted, err
}
