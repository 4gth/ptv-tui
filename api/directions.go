package api

import (
	"log"
	"ptv-tui/client"

	"ptv-tui/client/directions"
)

func GetDirections(client client.Ptvclient, parameters Parameters) []Responses {
	var params directions.DirectionsForRouteParams

	if parameters.RouteID != nil {
		params.SetRouteID(*parameters.RouteID)
	}

	GetResponse, err := client.Directions.DirectionsForRoute(&params)
	if err != nil {
		log.Fatal("Unable to read directions", GetResponse.Code())
	}

	var results []Responses
	for _, r := range GetResponse.Payload.Directions {
		resp := Responses{
			DirectionName: r.DirectionName,
			DirectionID:   r.DirectionID,
		}
		results = append(results, resp)

	}

	return results
}
