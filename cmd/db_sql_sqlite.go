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
	"github.com/spf13/viper"
)

var dbSQLSqliteOpts struct {
	File string
}

// dbSQLSqliteCmd represents the sqlite command
var dbSQLSqliteCmd = &cobra.Command{
	Use:   "sqlite",
	Short: "SQLite ingestion commands",
	Run: func(cmd *cobra.Command, args []string) {
		dbOpts.Driver = sql.SQLite(dbSQLSqliteOpts.File)
	},
}

func init() {
	dbSQLCmd.AddCommand(dbSQLSqliteCmd)

	bindEnv("file")

	viper.SetDefault("file", "sqlite.db")
	dbSQLSqliteCmd.Flags().StringVarP(&dbSQLSqliteOpts.File, "file", "f", viper.GetString("file"), "SQLite database file")
}
