package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tommyp1ckles/places/places"
)

var (
	GotoCmds = &cobra.Command{
		Use:   "goto",
		Short: "<image_path>",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				cmd.Help()
			}
			path := args[0]
			err := places.Goto(path)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
		},
	}
)
