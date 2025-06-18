package main

import (
	"log"
	"ptv-tui/api"
	joinutil "ptv-tui/helper"
	"ptv-tui/ui"
)

func main() {
	//var Routes []api.Responses
	var Depatures []api.Responses
	var StopDetails []api.Responses
	var Enriched []api.Responses
	ptv := api.NewPTVClient()

	params := []api.Parameters{
		{
			RouteTypes:    []int32{1},
			RouteType:     1,
			RouteID:       func(i int32) *int32 { return &i }(721),
			StopID:        2266,
			MaxResults:    int32(5),
			LookBackwards: false,
		},
	}

	Depatures = api.GetDepatures(ptv, params)
	//Routes = api.GetRoutes(ptv, params)
	StopDetails = api.GetStopDetails(ptv, params)
	enrichedMap := make(map[int32]*api.Responses)

	joinutil.MergeToMap(enrichedMap, Depatures, func(d api.Responses) int32 {
		return d.RouteID
	}, func(dst *api.Responses, src api.Responses) {
		dst.SchDepTime = src.SchDepTime
		dst.EstDepTime = src.EstDepTime
		dst.DirectionID = src.DirectionID
		dst.StopID = src.StopID
		dst.StopName = src.StopName
	})

	//joinutil.MergeToMap(enrichedMap, Routes, func(r api.Responses) int32 {
	//	return r.RouteID
	//}, func(dst *api.Responses, src api.Responses) {
	//	dst.Name = src.Name
	//	dst.Number = src.Number
	//	dst.RouteType = src.RouteType
	//})

	joinutil.MergeToMap(enrichedMap, StopDetails, func(s api.Responses) int32 {
		return s.StopID
	}, func(dst *api.Responses, src api.Responses) {
		dst.StopName = src.StopName
		dst.Routes = src.Routes
		dst.StopID = src.StopID
	})

	for _, enriched := range enrichedMap {
		Enriched = append(Enriched, *enriched)
	}

	//log.Println("Routes:", Routes)
	log.Println("Depatures:", Depatures)
	log.Println("StopDetails:", StopDetails)
	log.Println("Enriched:", Enriched)
	if err := ui.StartNewTea(&StopDetails); err != nil {
		log.Fatal(err)
	}
	if err := ui.StartNewTea(&Enriched); err != nil {
		log.Fatal(err)
	}
}
