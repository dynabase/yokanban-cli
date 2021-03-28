package cmd

import (
	"yokanban-cli/internal/api"
	"yokanban-cli/internal/elements"

	logger "github.com/sirupsen/logrus"

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
		api.DeleteBoard(deleteID)
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
		logger.Error(err)
	}

	// subCommands
	deleteCmd.AddCommand(deleteBoardSubCmd)
}
