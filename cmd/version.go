/*
Copyright © 2021 The Yokanban CLI Authors.

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
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var version string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of yokanban-cli",
	Long:  "Print the version number of yokanban-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yokanban-cli version", determineVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func determineVersion() string {
	defaultVersion := "0.0.1-SNAPSHOT"

	absPath, err := filepath.Abs("./VERSION")
	if err != nil {
		return defaultVersion
	}

	// I: Return version provided by build flag
	if version != "" {
		return version
	}

	// II: Try to determine version

	// Return defaultVersion, if no VERSION file can be found
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return defaultVersion
	}

	// default to SNAPSHOT version, if no build flag is provided though a VERSION file exists
	if version == "" {
		data, err := os.ReadFile(absPath)
		if err != nil {
			return defaultVersion
		}

		// cleanup read file contents
		detectedVersion := strings.TrimSpace(strings.TrimRight(string(data), "\r\n"))

		return fmt.Sprintf("%s-SNAPSHOT", detectedVersion)
	}

	return defaultVersion
}
