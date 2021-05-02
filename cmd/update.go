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
	"yokanban-cli/internal/api"
	"yokanban-cli/internal/elements"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var updateID string
var updateTitle string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:       "update",
	Short:     "Update yokanban resources like boards, cards, etc.",
	Example:   "yokanban update board --id 605f526126f0535cfd7fd6c7 --title test-board-update",
	ValidArgs: []string{string(elements.Board)},
	Args:      cobra.ExactValidArgs(1),
	Run:       func(cmd *cobra.Command, args []string) {},
}

// updateBoardSubCmd has the responsibility to update a yokanban board
var updateBoardSubCmd = &cobra.Command{
	Use:     "board",
	Short:   "Update a yokanban board",
	Example: "yokanban update board --id 605f526126f0535cfd7fd6c7 --title test-board-update",
	Run: func(cmd *cobra.Command, args []string) {
		a := getAPI()
		body := a.UpdateBoard(updateID, api.UpdateBoardDTO{NewName: updateTitle})
		fmt.Println(body)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	updateCmd.PersistentFlags().StringVarP(&updateID, "id", "i", "", "The id of the resource")
	if err := updateCmd.MarkPersistentFlagRequired("id"); err != nil {
		log.Error(err)
	}

	updateCmd.PersistentFlags().StringVarP(&updateTitle, "title", "n", "", "The title of the resource")

	// subCommands
	updateCmd.AddCommand(updateBoardSubCmd)
}
