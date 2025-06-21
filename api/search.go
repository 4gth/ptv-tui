package api

import (
	"log"
	"ptv-tui/client"
	"ptv-tui/client/stops"
)

func GetStopsFromLoc(client client.Ptvclient, parameters []Parameters) []Responses {
	var params stops.StopsStopsByGeolocationParams
	for _, p := range parameters {
		if p.Latitude != 0 {
			params.SetLatitude(p.Latitude)
		}
		if p.Longitude != 0 {
			params.SetLongitude(p.Longitude)
		}
		if p.MaxDistance != 0 {
			params.SetMaxDistance(&p.MaxDistance)
		}
		if p.RouteTypes != nil {
			params.SetRouteTypes(p.RouteTypes)
		}
	}
	payload, err := client.Stops.StopsStopsByGeolocation(&params)
	if err != nil {
		log.Fatal(err)
	}
	return BuildSearchRoutesResponses(payload)
}

//func GetStopsFromSearch(client client.Ptvclient, parameters []Parameters) []Responses {
//	var params search.SearchSearchParams
//
//	for _, p := range parameters {
//		if p.RouteType != 0 {
//			params.SetRouteTypes(p.RouteTypes)
//		}
//		if p.SearchTerm != "" {
//			params.SetSearchTerm(p.SearchTerm)
//		}
//		if p.Latitude != 0 {
//			params.SetLatitude(&p.Latitude)
//		}
//		if p.Longitude != 0 {
//			params.SetLongitude(&p.Longitude)
//		}
//		if p.MaxDistance != 0 {
//			params.SetMaxDistance(&p.MaxDistance)
//		}
//
//	}
//	searchResponse, err := client.Search.SearchSearch(&params)
//	if err != nil {
//
//	}
//
//	return BuildSearchRoutesResponses(searchResponse)
//}

func BuildSearchRoutesResponses(payload *stops.StopsStopsByGeolocationOK) []Responses {
	if payload == nil {

		return nil
	}

	var results []Responses

	for _, route := range payload.Payload.Stops {
		if route.RouteType == 1 {
			resp := Responses{
				StopID:   []int32{route.StopID},
				StopName: route.StopName,
			}

			results = append(results, resp)
		}
	}
	return results
}
