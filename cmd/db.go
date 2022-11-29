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

	"github.com/mrsimonemms/gobblr/pkg/drivers"
	"github.com/mrsimonemms/gobblr/pkg/gobblr"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dbOpts struct {
	DataPath string
	Driver   drivers.Driver
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

		inserted, err := gobblr.Execute(dbOpts.DataPath, dbOpts.Driver)
		if err != nil {
			return err
		}

		jsonData, err := json.MarshalIndent(inserted, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(jsonData))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)

	currentPath, err := os.Getwd()
	cobra.CheckErr(err)

	viper.SetDefault("path", path.Join(currentPath, "data"))
	dbCmd.PersistentFlags().StringVar(&dbOpts.DataPath, "path", viper.GetString("path"), "location of the data files")
}
