package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tommyp1ckles/places/places"
)

var (
	SingleFileCmds = &cobra.Command{
		Use:   "image",
		Short: "<dir> ...",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cmd.Help()
			}
			for _, path := range args {
				err := places.Image(path)
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}
			}
		},
	}
)
