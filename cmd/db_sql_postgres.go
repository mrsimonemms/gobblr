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

var dbSQLPostgresOpts struct {
	Database string
	Host     string
	Password string
	Port     int
	User     string
}

// dbSQLPostgresCmd represents the postgres command
var dbSQLPostgresCmd = &cobra.Command{
	Use:   "postgres",
	Short: "PostgreSQL ingestion commands",
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

func init() {
	dbSQLCmd.AddCommand(dbSQLPostgresCmd)

	bindEnv("database")
	bindEnv("host")
	bindEnv("password")
	bindEnv("port")
	bindEnv("username")

	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 5432)
	viper.SetDefault("username", "postgres")
	dbSQLPostgresCmd.Flags().
		StringVarP(&dbSQLPostgresOpts.Database, "database", "d", viper.GetString("database"), "name of the database to use")
	dbSQLPostgresCmd.Flags().StringVarP(&dbSQLPostgresOpts.Host, "host", "H", viper.GetString("host"), "database host name")
	dbSQLPostgresCmd.Flags().StringVarP(&dbSQLPostgresOpts.Password, "password", "p", viper.GetString("password"), "database password")
	dbSQLPostgresCmd.Flags().IntVarP(&dbSQLPostgresOpts.Port, "port", "P", viper.GetInt("port"), "database port")
	dbSQLPostgresCmd.Flags().StringVarP(&dbSQLPostgresOpts.User, "username", "u", viper.GetString("username"), "database username")
}
