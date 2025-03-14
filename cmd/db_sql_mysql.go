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

var dbSQLMysqlOpts struct {
	Database string
	Host     string
	Password string
	Port     int
	User     string
}

// dbSQLMysqlCmd represents the mysql command
var dbSQLMysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "MySQL ingestion commands",
	Run: func(cmd *cobra.Command, args []string) {
		dbOpts.Driver = sql.MySQL(
			dbSQLMysqlOpts.Database,
			dbSQLMysqlOpts.Host,
			dbSQLMysqlOpts.Password,
			dbSQLMysqlOpts.Port,
			dbSQLMysqlOpts.User,
		)
	},
}

func init() {
	dbSQLCmd.AddCommand(dbSQLMysqlCmd)

	bindEnv("database")
	bindEnv("host")
	bindEnv("password")
	bindEnv("port")
	bindEnv("username")

	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 3306)
	viper.SetDefault("username", "root")
	dbSQLMysqlCmd.Flags().StringVarP(&dbSQLMysqlOpts.Database, "database", "d", viper.GetString("database"), "name of the database to use")
	dbSQLMysqlCmd.Flags().StringVarP(&dbSQLMysqlOpts.Host, "host", "H", viper.GetString("host"), "database host name")
	dbSQLMysqlCmd.Flags().StringVarP(&dbSQLMysqlOpts.Password, "password", "p", viper.GetString("password"), "database password")
	dbSQLMysqlCmd.Flags().IntVarP(&dbSQLMysqlOpts.Port, "port", "P", viper.GetInt("port"), "database port")
	dbSQLMysqlCmd.Flags().StringVarP(&dbSQLMysqlOpts.User, "username", "u", viper.GetString("username"), "database username")
}
