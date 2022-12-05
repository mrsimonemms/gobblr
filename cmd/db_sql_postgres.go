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

var dbSqlPostgresOpts struct {
	Database string
	Host     string
	Password string
	Port     int
	User     string
}

// dbSqlPostgresCmd represents the postgres command
var dbSqlPostgresCmd = &cobra.Command{
	Use:   "postgres",
	Short: "PostgreSQL ingestion commands",
	Run: func(cmd *cobra.Command, args []string) {
		dbOpts.Driver = sql.PostgreSQL(
			dbSqlPostgresOpts.Database,
			dbSqlPostgresOpts.Host,
			dbSqlPostgresOpts.Password,
			dbSqlPostgresOpts.Port,
			dbSqlPostgresOpts.User,
		)
	},
}

func init() {
	dbSqlCmd.AddCommand(dbSqlPostgresCmd)

	bindEnv("database")
	bindEnv("host")
	bindEnv("password")
	bindEnv("port")
	bindEnv("username")

	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 5432)
	viper.SetDefault("username", "postgres")
	dbSqlPostgresCmd.Flags().StringVarP(&dbSqlPostgresOpts.Database, "database", "d", viper.GetString("database"), "name of the database to use")
	dbSqlPostgresCmd.Flags().StringVarP(&dbSqlPostgresOpts.Host, "host", "H", viper.GetString("host"), "database host name")
	dbSqlPostgresCmd.Flags().StringVarP(&dbSqlPostgresOpts.Password, "password", "p", viper.GetString("password"), "database password")
	dbSqlPostgresCmd.Flags().IntVarP(&dbSqlPostgresOpts.Port, "port", "P", viper.GetInt("port"), "database port")
	dbSqlPostgresCmd.Flags().StringVarP(&dbSqlPostgresOpts.User, "username", "u", viper.GetString("username"), "database username")
}
