package main

import (
	"log"
	"ptv-tui/api"
	joinutil "ptv-tui/helper"
	"ptv-tui/ui"
)

func main() {
	// Initialize the PTV client and parameters for fetching departure information.
	// This example fetches departures for a specific stop (ID: 2206) and
	// enriches the data with route and stop details.
	// TODO: Make the parameters configurable via UI.
	// TODO: Add error handling for API calls and UI rendering.

	log.Println("Starting PTV TUI...")

	var Routes []api.Responses
	var Depatures []api.Responses
	var StopDetails []api.Responses
	var Enriched []api.Responses

	ptv := api.NewPTVClient()

	params := []api.Parameters{
		{
			RouteTypes:    []int32{1},
			RouteType:     1,
			RouteID:       func(i int32) *int32 { return &i }(721),
			StopID:        2206,
			MaxResults:    int32(5),
			LookBackwards: false,
		},
	}

	Depatures = api.GetDepatures(ptv, params)
	Routes = api.GetRoutes(ptv, params)
	StopDetails = api.GetStopDetails(ptv, params)

	// Enrich the depatures with stop details and routes
	Enriched = joinutil.EnrichDepatures(Depatures, StopDetails, Routes)

	for _, l := range Enriched {
		log.Printf("Enriched: %+v", l)
	}

	if err := ui.BrewDepatureTea(Enriched); err != nil {
		log.Fatal(err)
	}

}
