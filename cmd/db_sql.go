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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type dbSQLOpts struct {
	Database string
	Host     string
	Password string
	Port     int
	User     string
}

func dbSQLFlags(cmd *cobra.Command, opts *dbSQLOpts) {
	bindEnv("database")
	bindEnv("host")
	bindEnv("password")
	bindEnv("port")
	bindEnv("username")

	viper.SetDefault("host", opts.Host)
	viper.SetDefault("port", opts.Port)
	viper.SetDefault("username", opts.User)
	cmd.Flags().StringVarP(&opts.Database, "database", "d", viper.GetString("database"), "name of the database to use")
	cmd.Flags().StringVarP(&opts.Host, "host", "H", viper.GetString("host"), "database host name")
	cmd.Flags().StringVarP(&opts.Password, "password", "p", viper.GetString("password"), "database password")
	cmd.Flags().IntVarP(&opts.Port, "port", "P", viper.GetInt("port"), "database port")
	cmd.Flags().StringVarP(&opts.User, "username", "u", viper.GetString("username"), "database username")
}

// dbSQLCmd represents the sql command
var dbSQLCmd = &cobra.Command{
	Use:   "sql",
	Short: "SQL ingestion commands",
}

func init() {
	dbCmd.AddCommand(dbSQLCmd)
	dbSQLCmd.AddCommand(MakeSQLMysql())
	dbSQLCmd.AddCommand(MakeSQLPostgres())
	dbSQLCmd.AddCommand(MakeSQLSqlserver())
}
