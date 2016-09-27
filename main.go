package main

import (
	"log"

	"github.com/tommyp1ckles/places/cmd"
)

func main() {
	placesRootCommand := cmd.CreatePlacesCommand()
	if err := placesRootCommand.Execute(); err != nil {
		log.Fatalln(err)
	}
}
