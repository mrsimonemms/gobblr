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

var dbSqlSqliteOpts struct {
	File string
}

// dbSqlSqliteCmd represents the sqlite command
var dbSqlSqliteCmd = &cobra.Command{
	Use:   "sqlite",
	Short: "SQLite ingestion commands",
	Run: func(cmd *cobra.Command, args []string) {
		dbOpts.Driver = sql.SQLite(dbSqlSqliteOpts.File)
	},
}

func init() {
	dbSqlCmd.AddCommand(dbSqlSqliteCmd)

	bindEnv("file")

	viper.SetDefault("file", "sqlite.db")
	dbSqlSqliteCmd.Flags().StringVarP(&dbSqlSqliteOpts.File, "file", "f", viper.GetString("file"), "SQLite database file")
}
