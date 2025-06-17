package main

import (
	"log"
	"ptv-tui/api"
	"ptv-tui/ui"
)

func main() {

	ptv := api.NewPTVClient()

	//Routes := api.GetRoutes(ptv)
	Depatures := api.GetDepatures(ptv)

	log.Println(Depatures)

	//if err := ui.StartNewTea(Routes); err != nil {
	//	log.Fatal(err)
	//}
	if err := ui.StartNewTea(Depatures); err != nil {
		log.Fatal(err)
	}
}
