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
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/mrsimonemms/gobblr/pkg/drivers"
	"github.com/mrsimonemms/gobblr/pkg/gobblr"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dbOpts struct {
	DataPath     string
	Driver       drivers.Driver
	Retries      uint64
	RunWebServer bool
	WebPort      int
}

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Control the dataset in your database",
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		// Use a persistent post run command as each of the subcommands is there
		// to create the configure only. The execution happens here.
		//
		// There can be only one PersistentPostRun command.

		inserted, err := gobblr.Execute(dbOpts.DataPath, dbOpts.Driver, dbOpts.Retries)
		if err != nil {
			return err
		}

		if dbOpts.RunWebServer {
			// Runs the execution as a server to make it easy to trigger it from integration tests etc
			return gobblr.Serve(dbOpts.DataPath, dbOpts.Driver, dbOpts.Retries, dbOpts.WebPort)
		}

		jsonData, err := json.MarshalIndent(inserted, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(jsonData))

		return nil
	},
}

const (
	envvarPrefix = "GOBBLR"
)

func bindEnv(key string) {
	envvarName := fmt.Sprintf("%s_%s", envvarPrefix, key)
	envvarName = strings.Replace(envvarName, "-", "_", -1)
	envvarName = strings.ToUpper(envvarName)

	err := viper.BindEnv(key, envvarName)
	cobra.CheckErr(err)
}

func init() {
	rootCmd.AddCommand(dbCmd)

	currentPath, err := os.Getwd()
	cobra.CheckErr(err)

	bindEnv("path")
	bindEnv("retries")
	bindEnv("run")
	bindEnv("run-port")

	viper.SetDefault("path", path.Join(currentPath, "data"))
	viper.SetDefault("retries", 0)
	viper.SetDefault("run", false)
	viper.SetDefault("run-port", 5670) // Default to a random, normally-unused port
	dbCmd.PersistentFlags().StringVar(&dbOpts.DataPath, "path", viper.GetString("path"), "location of the data files")
	dbCmd.PersistentFlags().Uint64Var(&dbOpts.Retries, "retries", viper.GetUint64("retries"), "number of retries before declaring a failure")
	dbCmd.PersistentFlags().BoolVar(&dbOpts.RunWebServer, "run", viper.GetBool("run"), "run as a web server")
	dbCmd.PersistentFlags().IntVar(&dbOpts.WebPort, "run-port", viper.GetInt("run-port"), "port for web server to listen on")

}
