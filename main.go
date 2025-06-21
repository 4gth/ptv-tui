package main

import (
	"log"
	"ptv-tui/api"
	"ptv-tui/ui"
)

func main() {
	// Initialize the PTV client and parameters for fetching departure information.
	// This example fetches departures for a specific stop (ID: 2206) and
	// enriches the data with route and stop details.
	// TODO: Make the parameters configurable via UI.
	// TODO: Add error handling for API calls and UI rendering.

	//var Routes []api.Responses
	//var Depatures []api.Responses
	//var StopDetails []api.Responses
	var SearchResults []api.Responses
	var Depatures []api.Responses
	var Routes []api.Responses
	//var Direction []api.Responses
	var Enriched []api.Responses

	ptv := api.NewPTVClient()

	params := []api.Parameters{
		{
			RouteTypes:            []int32{1},
			RouteType:             0,
			RouteID:               func(i int32) *int32 { return &i }(721),
			StopID:                2206,
			MaxResults:            int32(5),
			LookBackwards:         false,
			SearchTerm:            "Southbank",
			Latitude:              -37.820648,
			Longitude:             144.965286,
			MaxDistance:           600,
			MatchStopByGtfsStopID: true,
			StopIDs:               []int32{},
		},
	}

	SearchResults = api.GetStopsFromLoc(ptv, params)
	Routes = api.GetRoutes(ptv, params)

	for _, s := range SearchResults {
		params[0].StopID = s.StopID[0]
		params[0].RouteName = &s.Name
		Depatures = api.GetDepatures(ptv, params)
		for i := range Depatures {
			dirparams := api.Parameters{RouteID: &Depatures[i].RouteID}
			directions := api.GetDirections(ptv, dirparams)
			for _, dir := range directions {
				if Depatures[i].DirectionID == dir.DirectionID {
					Depatures[i].DirectionName = dir.DirectionName
				}
			}
		}
		for _, d := range Depatures {
			var matchedRouteName string
			var matchedRouteNumber string
			//var matchedDireciton string
			for _, r := range Routes {

				if r.RouteID == d.RouteID {
					matchedRouteName = r.Name
					matchedRouteNumber = r.Number
					//matchedDireciton = r.DirectionName

					break
				}
			}

			Enriched = append(Enriched, api.Responses{
				StopID:        s.StopID,
				StopName:      s.StopName,
				DirectionName: d.DirectionName,
				EstDepTime:    d.EstDepTime,
				SchDepTime:    d.SchDepTime,
				RouteID:       d.RouteID,
				Name:          matchedRouteName,
				Number:        matchedRouteNumber,
			})
		}
	}

	//for _, p := range Enriched {
	//
	//}
	if err := ui.BrewDepatureTea(Enriched); err != nil {
		log.Fatal(err)
	}

}
