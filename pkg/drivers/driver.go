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

	// Insert data in bulk
	InsertBulk(table string, data []map[string]interface{}) (inserted int, err error)

	// Remove all the data from the table
	Truncate(table string) error
}
