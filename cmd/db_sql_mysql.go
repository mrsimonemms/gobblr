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

var dbSqlMysqlOpts struct {
	Database string
	Host     string
	Password string
	Port     int
	User     string
}

// dbSqlMysqlCmd represents the mysql command
var dbSqlMysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "MySQL ingestion commands",
	Run: func(cmd *cobra.Command, args []string) {
		dbOpts.Driver = sql.MySQL(
			dbSqlMysqlOpts.Database,
			dbSqlMysqlOpts.Host,
			dbSqlMysqlOpts.Password,
			dbSqlMysqlOpts.Port,
			dbSqlMysqlOpts.User,
		)
	},
}

func init() {
	dbSqlCmd.AddCommand(dbSqlMysqlCmd)

	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 3306)
	viper.SetDefault("username", "root")
	dbSqlMysqlCmd.Flags().StringVarP(&dbSqlMysqlOpts.Database, "database", "d", viper.GetString("database"), "name of the database to use")
	dbSqlMysqlCmd.Flags().StringVarP(&dbSqlMysqlOpts.Host, "host", "H", viper.GetString("host"), "database host name")
	dbSqlMysqlCmd.Flags().StringVarP(&dbSqlMysqlOpts.Password, "password", "p", viper.GetString("password"), "database password")
	dbSqlMysqlCmd.Flags().IntVarP(&dbSqlMysqlOpts.Port, "port", "P", viper.GetInt("port"), "database port")
	dbSqlMysqlCmd.Flags().StringVarP(&dbSqlMysqlOpts.User, "username", "u", viper.GetString("username"), "database username")
}
