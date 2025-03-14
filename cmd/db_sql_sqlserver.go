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

var dbSQLSqlserverOpts struct {
	Database string
	Host     string
	Password string
	Port     int
	User     string
}

// dbSQLSqlserverCmd represents the sqlserver command
var dbSQLSqlserverCmd = &cobra.Command{
	Use:   "sqlserver",
	Short: "SQL server ingestion commands",
	Run: func(cmd *cobra.Command, args []string) {
		dbOpts.Driver = sql.SQLServer(
			dbSQLSqlserverOpts.Database,
			dbSQLSqlserverOpts.Host,
			dbSQLSqlserverOpts.Password,
			dbSQLSqlserverOpts.Port,
			dbSQLSqlserverOpts.User,
		)
	},
}

func init() {
	dbSQLCmd.AddCommand(dbSQLSqlserverCmd)

	bindEnv("database")
	bindEnv("host")
	bindEnv("password")
	bindEnv("port")
	bindEnv("username")

	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 1433)
	viper.SetDefault("username", "sa")
	dbSQLSqlserverCmd.Flags().
		StringVarP(&dbSQLSqlserverOpts.Database, "database", "d", viper.GetString("database"), "name of the database to use")
	dbSQLSqlserverCmd.Flags().StringVarP(&dbSQLSqlserverOpts.Host, "host", "H", viper.GetString("host"), "database host name")
	dbSQLSqlserverCmd.Flags().StringVarP(&dbSQLSqlserverOpts.Password, "password", "p", viper.GetString("password"), "database password")
	dbSQLSqlserverCmd.Flags().IntVarP(&dbSQLSqlserverOpts.Port, "port", "P", viper.GetInt("port"), "database port")
	dbSQLSqlserverCmd.Flags().StringVarP(&dbSQLSqlserverOpts.User, "username", "u", viper.GetString("username"), "database username")
}
