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
	"github.com/mrsimonemms/gobblr/pkg/drivers/mongodb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dbMongodbOpts struct {
	Database      string
	ConnectionURI string
}

// dbMongodbCmd represents the mongodb command
var dbMongodbCmd = &cobra.Command{
	Use:   "mongodb",
	Short: "MongoDB ingestion commands",
	Run: func(cmd *cobra.Command, args []string) {
		dbOpts.Driver = mongodb.New(
			dbMongodbOpts.ConnectionURI,
			dbMongodbOpts.Database,
		)
	},
}

func init() {
	dbCmd.AddCommand(dbMongodbCmd)

	bindEnv("connection-uri")
	bindEnv("database")

	viper.SetDefault("connection-uri", "mongodb://localhost:27017")
	dbMongodbCmd.Flags().
		StringVarP(&dbMongodbOpts.ConnectionURI, "connection-uri", "u", viper.GetString("connection-uri"), "database connection uri")
	dbMongodbCmd.Flags().StringVarP(&dbMongodbOpts.Database, "database", "d", viper.GetString("database"), "name of the database to use")
}
