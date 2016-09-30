package main

import (
	"log"

	"github.com/tommyp1ckles/places/cmd"
)

func main() {
	root := cmd.CreatePlacesRootCommand()
	if err := root.Execute(); err != nil {
		log.Fatalln(err)
	}
}
