package cmd

import (
	"yokanban-cli/internal/api"
	"yokanban-cli/internal/elements"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var getID string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:       "get",
	Short:     "Get single yokanban resources like boards, cards, etc.",
	Example:   "yokanban get board --id 605f526126f0535cfd7fd6c7",
	ValidArgs: []string{string(elements.Board)},
	Args:      cobra.ExactValidArgs(1),
	Run:       func(cmd *cobra.Command, args []string) {},
}

// getBoardSubCmd has the responsibility to get a yokanban board
var getBoardSubCmd = &cobra.Command{
	Use:     "board",
	Short:   "Get a single yokanban board",
	Example: "yokanban get board --id 605f526126f0535cfd7fd6c7",
	Run: func(cmd *cobra.Command, args []string) {
		api.GetBoard(getID)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	getCmd.PersistentFlags().StringVarP(&getID, "id", "i", "", "The id of the resource")
	if err := getCmd.MarkPersistentFlagRequired("id"); err != nil {
		log.Error(err)
	}

	// subCommands
	getCmd.AddCommand(getBoardSubCmd)
}
