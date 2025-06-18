package api

import (
	"log"
	"ptv-tui/client"
	"ptv-tui/client/stops"
	"ptv-tui/models"
)

func GetStopDetails(client client.Ptvclient, parameters []Parameters) []Responses {
	var params stops.StopsStopDetailsParams

	for _, p := range parameters {
		if p.RouteType != 0 {
			params.SetRouteType(p.RouteType)
		}
		if p.StopID != 0 {
			params.SetStopID(p.StopID)
		}
	}
	stopResponse, err := client.Stops.StopsStopDetails(&params)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(stopResponse.Code())
	stopList := BuildStopDetailsResponses(stopResponse.Payload)
	log.Println(stopResponse.Payload)
	return stopList

}

func BuildStopDetailsResponses(payload *models.V3StopResponse) []Responses {
	if payload == nil {
		return nil
	}

	resp := Responses{
		StopID:   payload.Stop.StopID,
		StopName: payload.Stop.StopName,
		Routes:   payload.Stop.Routes,
	}
	return []Responses{resp}
}
