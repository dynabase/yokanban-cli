package cmd

import (
	"github.com/spf13/cobra"
	"yokanban-cli/internal/api"
	"yokanban-cli/internal/elements"
)

var name string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:       "create",
	Short:     "Create yokanban resources like boards, cards, etc.",
	Long:      `Create yokanban resources like boards, cards, etc.`,
	Example:   "yokanban create board --name test-board",
	ValidArgs: []string{string(elements.Board)},
	Args:      cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == string(elements.Board) {
			api.CreateBoard(api.CreateBoardModel{Name: name})
		}
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

	createCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the resource")
}
