package cmd

import (
	"yokanban-cli/internal/api"
	"yokanban-cli/internal/elements"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var updateID string
var updateName string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:       "update",
	Short:     "Updates yokanban resources like boards, cards, etc.",
	Example:   "yokanban update board --id 605f526126f0535cfd7fd6c7 --name test-board-update",
	ValidArgs: []string{string(elements.Board)},
	Args:      cobra.ExactValidArgs(1),
	Run:       func(cmd *cobra.Command, args []string) {},
}

// updateBoardSubCmd has the responsibility to update a yokanban board
var updateBoardSubCmd = &cobra.Command{
	Use:     "board",
	Short:   "Update a yokanban board",
	Example: "yokanban update board --id 605f526126f0535cfd7fd6c7 --name test-board-update",
	Run: func(cmd *cobra.Command, args []string) {
		api.UpdateBoard(updateID, api.UpdateBoardModel{NewName: updateName})
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
		logger.Error(err)
	}

	updateCmd.PersistentFlags().StringVarP(&updateName, "name", "n", "", "The name of the resource")

	// subCommands
	updateCmd.AddCommand(updateBoardSubCmd)
}
