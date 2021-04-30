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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:       "list",
	Short:     "List yokanban resources like boards, cards, etc.",
	Example:   "yokanban list boards",
	ValidArgs: []string{string(elements.Boards), string(elements.Board)},
	Args:      cobra.ExactValidArgs(1),
	Run:       func(cmd *cobra.Command, args []string) {},
}

// listBoardsSubCmd has the responsibility to list yokanban boards of a user
var listBoardsSubCmd = &cobra.Command{
	Use:     "boards",
	Aliases: []string{"board"},
	Short:   "Lists yokanban boards current user has access to",
	Example: "yokanban list boards",
	Run: func(cmd *cobra.Command, args []string) {
		a := getAPI()
		boardList := a.ListBoards()

		// generate the pretty printed output
		boardsPretty, err := json.MarshalIndent(boardList, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(boardsPretty))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// subCommands
	listCmd.AddCommand(listBoardsSubCmd)
}
