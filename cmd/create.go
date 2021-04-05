package cmd

import (
	"encoding/json"
	"fmt"
	"yokanban-cli/internal/api"
	"yokanban-cli/internal/elements"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var createName string
var createOnBoardID string

// createCmd represents the root create command
var createCmd = &cobra.Command{
	Use:       "create [element]",
	Short:     "Create yokanban resources like boards, cards, etc.",
	Example:   "yokanban create board --name test-board",
	ValidArgs: []string{string(elements.Board), string(elements.Column)},
	Args:      cobra.ExactValidArgs(1),
	Run:       func(cmd *cobra.Command, args []string) {},
}

// createBoardSubCmd has the responsibility to create a yokanban board
var createBoardSubCmd = &cobra.Command{
	Use:     "board",
	Short:   "Create a yokanban board",
	Example: "yokanban create board --name test-board",
	Run: func(cmd *cobra.Command, args []string) {
		body := api.CreateBoard(api.CreateBoardDTO{Name: createName})
		fmt.Println(body)
	},
}

// createColumnSubCmd has the responsibility to create a yokanban column on a board
var createColumnSubCmd = &cobra.Command{
	Use:     "column",
	Short:   "Create a yokanban column",
	Example: "yokanban create column --name test-column --board-id 605f574e26f0535cfd7fd6cd",
	Run: func(cmd *cobra.Command, args []string) {
		details := api.CreateColumn(createOnBoardID, createName)

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

	createCmd.PersistentFlags().StringVarP(&createName, "name", "n", "", "The name of the resource")

	// create board
	createCmd.AddCommand(createBoardSubCmd)

	// create column
	createCmd.AddCommand(createColumnSubCmd)
	createColumnSubCmd.PersistentFlags().StringVarP(&createOnBoardID, "board-id", "", "", "The id of the board")
	if err := createColumnSubCmd.MarkPersistentFlagRequired("board-id"); err != nil {
		log.Error(err)
	}
}
