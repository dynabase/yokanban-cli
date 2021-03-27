package cmd

import (
	"github.com/spf13/cobra"
	"yokanban-cli/internal/api"
	"yokanban-cli/internal/elements"
)

var createName string

// createCmd represents the root create command
var createCmd = &cobra.Command{
	Use:       "create [element]",
	Short:     "Create yokanban resources like boards, cards, etc.",
	Example:   "yokanban create board --name test-board",
	ValidArgs: []string{string(elements.Board)},
	Args:      cobra.ExactValidArgs(1),
	Run:       func(cmd *cobra.Command, args []string) {},
}

// createBoardSubCmd has the responsibility to create a yokanban board
var createBoardSubCmd = &cobra.Command{
	Use:     "board",
	Short:   "Create a yokanban board",
	Example: "yokanban create board --name test-board",
	Run: func(cmd *cobra.Command, args []string) {
		api.CreateBoard(api.CreateBoardModel{Name: createName})
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

	// subCommands
	createCmd.AddCommand(createBoardSubCmd)
}
