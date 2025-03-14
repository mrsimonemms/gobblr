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

package drivers

// NB. In this file, the interfaces refer to "table", as per a SQL
// type database. This naming convention may differ, such as "collection"
// in a MongoDB instance. However, the functionality will remain the
// same from an interface point of view.

type Driver interface {
	// Authorize the connection to the database
	Auth() error

	// Close the database connection and free up resources
	Close() error

	// Name of the driver
	DriverName() string

	// Insert data in bulk
	InsertBulk(table string, data []map[string]interface{}) (inserted int, err error)

	// Remove all the data from the table
	Truncate(table string) error
}
