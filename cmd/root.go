package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func CreatePlacesRootCommand() *cobra.Command {
	PlacesRootCommand := cobra.Command{
		Use: "places",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				os.Exit(-1)
			}
		},
	}
	PlacesRootCommand.AddCommand(RecursivePlacesCmds)
	PlacesRootCommand.AddCommand(SingleFileCmds)
	PlacesRootCommand.AddCommand(GotoCmds)
	return &PlacesRootCommand
}
