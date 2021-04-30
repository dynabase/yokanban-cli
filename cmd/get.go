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
	"encoding/json"
	"fmt"
	"yokanban-cli/internal/elements"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var getID string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:       "get",
	Short:     "Get yokanban resources like boards, cards, etc.",
	Example:   "yokanban get board --id 605f526126f0535cfd7fd6c7",
	ValidArgs: []string{string(elements.Board)},
	Args:      cobra.ExactValidArgs(1),
	Run:       func(cmd *cobra.Command, args []string) {},
}

// getBoardSubCmd has the responsibility to get a yokanban board
var getBoardSubCmd = &cobra.Command{
	Use:     "board",
	Short:   "Get a single yokanban board",
	Example: "yokanban get board --id 605f526126f0535cfd7fd6c7",
	Run: func(cmd *cobra.Command, args []string) {
		a := getAPI()
		details := a.GetBoard(getID)

		// generate the pretty printed output
		pretty, err := json.MarshalIndent(details, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(pretty))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	getCmd.PersistentFlags().StringVarP(&getID, "id", "i", "", "The id of the resource")
	if err := getCmd.MarkPersistentFlagRequired("id"); err != nil {
		log.Error(err)
	}

	// subCommands
	getCmd.AddCommand(getBoardSubCmd)
}
