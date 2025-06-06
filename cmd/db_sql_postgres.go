/*
Copyright © 2022 Simon Emms <simon@simonemms.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/mrsimonemms/gobblr/pkg/drivers/sql"
	"github.com/spf13/cobra"
)

func MakeSQLPostgres() *cobra.Command {
	dbSQLPostgresOpts := &dbSQLOpts{
		Host: "localhost",
		Port: 5432,
		User: "postgres",
	}

	// dbSQLPostgresCmd represents the postgres command
	dbSQLPostgresCmd := &cobra.Command{
		Use:     "postgres",
		Aliases: []string{"pgsql"},
		Short:   "PostgreSQL ingestion commands",
		Run: func(cmd *cobra.Command, args []string) {
			dbOpts.Driver = sql.PostgreSQL(
				dbSQLPostgresOpts.Database,
				dbSQLPostgresOpts.Host,
				dbSQLPostgresOpts.Password,
				dbSQLPostgresOpts.Port,
				dbSQLPostgresOpts.User,
			)
		},
	}

	dbSQLFlags(dbSQLPostgresCmd, dbSQLPostgresOpts)

	return dbSQLPostgresCmd
}
