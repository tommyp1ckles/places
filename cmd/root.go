package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func CreatePlacesCommand() *cobra.Command {
	PlacesRootCommand := cobra.Command{
		Use: "places",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				os.Exit(-1)
			}
		},
	}
	PlacesRootCommand.Flags().BoolVarP(
		&config.Stats,
		"stats",
		"b",
		false,
		"Show places stats table",
	)
	PlacesRootCommand.AddCommand(RecursivePlacesCmds)
	return &PlacesRootCommand
}
