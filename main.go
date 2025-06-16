package main

import (
	"log"
	"ptv-tui/api"
	"ptv-tui/ui"
)

func main() {

	ptv := api.NewPTVClient()

	Routes := api.GetRoutes(ptv)

	if err := ui.StartNewTea(Routes); err != nil {
		log.Fatal(err)
	}

}
