package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tommyp1ckles/places/places"
)

var (
	config = places.Config{}
)

const (
	StatsMessage = "Show stats about your photo places"
)

var (
	RecursivePlacesCmds = &cobra.Command{
		Use:   "list",
		Short: "<dir> ...",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cmd.Help()
			}
			for _, path := range args {
				err := places.ListPlacesRecursively(path, &config)
				if err != nil && err == places.ErrNoToken {
					fmt.Println("Could not connect to Google Maps API, likely an invalid API Token.")
					return
				}
			}
		},
	}
)
