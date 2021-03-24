package cmd

import (
	"github.com/spf13/cobra"
	"yokanban-cli/internal/api"
	"yokanban-cli/internal/elements"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:       "create",
	Short:     "Create yokanban resources like boards, cards, etc.",
	Long:      `Create yokanban resources like boards, cards, etc.`,
	Example:   "yokanban create board",
	ValidArgs: []string{"board"}, // TODO add more commands later on
	Args:      cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "board" {
			api.Create(elements.Board)
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
}
