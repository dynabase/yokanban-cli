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
	"yokanban-cli/internal/elements"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var deleteID string

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:       "delete [element]",
	Short:     "Delete yokanban resources like boards, cards, etc.",
	Example:   "yokanban delete board --id 605f526126f0535cfd7fd6c7",
	ValidArgs: []string{string(elements.Board)},
	Args:      cobra.ExactValidArgs(1),
	Run:       func(cmd *cobra.Command, args []string) {},
}

// deleteBoardSubCmd has the responsibility to delete a yokanban board
var deleteBoardSubCmd = &cobra.Command{
	Use:     "board",
	Short:   "Delete a yokanban board",
	Example: "yokanban delete board --id 605f526126f0535cfd7fd6c7",
	Run: func(cmd *cobra.Command, args []string) {
		a := getAPI()
		body := a.DeleteBoard(deleteID)
		fmt.Println(body)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	deleteCmd.PersistentFlags().StringVarP(&deleteID, "id", "i", "", "The id of the resource")
	if err := deleteCmd.MarkPersistentFlagRequired("id"); err != nil {
		log.Error(err)
	}

	// subCommands
	deleteCmd.AddCommand(deleteBoardSubCmd)
}
