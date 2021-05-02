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
	"yokanban-cli/internal/api"
	"yokanban-cli/internal/elements"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var createTitle string
var createOnBoardID string

// createCmd represents the root create command
var createCmd = &cobra.Command{
	Use:       "create [element]",
	Short:     "Create yokanban resources like boards, cards, etc.",
	Example:   "yokanban create board --title test-board",
	ValidArgs: []string{string(elements.Board), string(elements.Column), string(elements.Card)},
	Args:      cobra.ExactValidArgs(1),
	Run:       func(cmd *cobra.Command, args []string) {},
}

// createBoardSubCmd has the responsibility to create a yokanban board
var createBoardSubCmd = &cobra.Command{
	Use:     "board",
	Short:   "Create a yokanban board",
	Example: "yokanban create board --title test-board",
	Run: func(cmd *cobra.Command, args []string) {
		a := getAPI()
		body := a.CreateBoard(api.CreateBoardDTO{Name: createTitle})
		fmt.Println(body)
	},
}

// createColumnSubCmd has the responsibility to create a yokanban column on a board
var createColumnSubCmd = &cobra.Command{
	Use:     "column",
	Short:   "Create a yokanban column",
	Example: "yokanban create column --title test-column --board-id 605f574e26f0535cfd7fd6cd",
	Run: func(cmd *cobra.Command, args []string) {
		a := getAPI()
		details := a.CreateColumn(createOnBoardID, createTitle, uuid.New())

		// generate the pretty printed output
		pretty, err := json.MarshalIndent(details, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(pretty))
	},
}

// createCardSubCmd has the responsibility to create a yokanban card on a board
var createCardSubCmd = &cobra.Command{
	Use:     "card",
	Short:   "Create a yokanban card",
	Example: "yokanban create card --title test-card --board-id 605f574e26f0535cfd7fd6cd",
	Run: func(cmd *cobra.Command, args []string) {
		a := getAPI()
		details := a.CreateCard(createOnBoardID, createTitle, uuid.New())

		// generate the pretty printed output
		pretty, err := json.MarshalIndent(details, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(pretty))
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	createCmd.PersistentFlags().StringVarP(&createTitle, "title", "n", "", "The title of the resource")

	// create board
	createCmd.AddCommand(createBoardSubCmd)

	// create column
	createCmd.AddCommand(createColumnSubCmd)
	createColumnSubCmd.PersistentFlags().StringVarP(&createOnBoardID, "board-id", "", "", "The id of the board")
	if err := createColumnSubCmd.MarkPersistentFlagRequired("board-id"); err != nil {
		log.Error(err)
	}

	// create card
	createCmd.AddCommand(createCardSubCmd)
	createCardSubCmd.PersistentFlags().StringVarP(&createOnBoardID, "board-id", "", "", "The id of the board")
	if err := createCardSubCmd.MarkPersistentFlagRequired("board-id"); err != nil {
		log.Error(err)
	}
}
