package cmd

import (
	"encoding/json"
	"fmt"
	"yokanban-cli/internal/api"
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
		boardList := api.ListBoards()

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
