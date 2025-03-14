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

package gobblr

import (
	"fmt"

	"github.com/cenkalti/backoff/v4"
	"github.com/mrsimonemms/gobblr/pkg/drivers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger *zerolog.Logger = &log.Logger

type Inserted struct {
	Table string `json:"table"`
	Count int    `json:"count"`
}

func Execute(dataPath string, db drivers.Driver, retries uint64) ([]Inserted, error) {
	attempt := 0
	return backoff.RetryWithData(
		func() ([]Inserted, error) {
			logger.Debug().Int("attempt", attempt).Msg("Attempting to gobble data")

			return retryExecute(dataPath, db)
		},
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), retries),
	)
}

func retryExecute(dataPath string, db drivers.Driver) ([]Inserted, error) {
	inserted := make([]Inserted, 0)
	var err error

	logCtx := (&log.Logger).With().Str("dbType", db.DriverName()).Logger()

	// Connect to database
	logCtx.Debug().Msg("Authenticating database")
	if err := db.Auth(); err != nil {
		logCtx.Error().Err(err).Msg("Failed to connect to database")
		return nil, err
	}

	// We've finished - let's clear up after ourselves
	defer func() {
		logCtx.Debug().Msg("Closing database connection")
		err = db.Close()
	}()

	// Find the files
	files, err := FindFiles(dataPath)
	if err != nil {
		logCtx.Error().Err(err).Msg("Failed to connect to database")
		return nil, err
	}
	logCtx.Debug().Str("path", dataPath).Int("count", len(files)).Msg("Found files to ingest")

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
