package api

import (
	"log"
	"ptv-tui/client"
	"ptv-tui/client/stops"
	"ptv-tui/models"
)

// GetStopDetails retrieves detailed information about a specific stop from the PTV API.
// It accepts a client and a slice of Parameters, which can include route types and stop IDs.
// The function constructs a request to the PTV API and returns a slice of Responses containing the stop details.
func GetStopDetails(client client.Ptvclient, parameters []Parameters) []Responses {
	var params stops.StopsStopDetailsParams
	var StopList []Responses
	for _, p := range parameters {
		if p.RouteType != 0 {
			params.SetRouteType(p.RouteType)
		}
		//if p.StopID != 0 {
		//	params.SetStopID(p.StopID)
		//}
		if p.StopIDs != nil {
			for _, stopID := range p.StopIDs {
				params.SetStopID(stopID)
			}

		}
		log.Printf("Fetching stop details for StopIDs: %+v using params: %+v", p.StopIDs, params.StopID)
		stopResponse, err := client.Stops.StopsStopDetails(&params)
		if err != nil {
			log.Printf("[%d]Failed to get Stop Details: %s", stopResponse.Code(), err)
			return nil // Return nil if an error occurs
		}
		StopList = append(StopList, BuildStopDetailsResponses(stopResponse.Payload)...)
		log.Printf("Found %d stops", len(StopList))

	}
	return StopList
}

// BuildStopDetailsResponses converts the StopsStopDetailsOK payload into a slice of Responses.
// It extracts relevant stop information such as StopID, StopName, and Routes.
// If the payload is nil, it returns an empty slice.
// Each stop is represented by a Responses struct, which includes the stop ID, name, and routes.
func BuildStopDetailsResponses(payload *models.V3StopResponse) []Responses {
	if payload == nil {
		return nil
	}

	resp := Responses{
		StopID:   []int32{payload.Stop.StopID},
		StopName: payload.Stop.StopName,
		Routes:   payload.Stop.Routes,
	}
	return []Responses{resp}
}
