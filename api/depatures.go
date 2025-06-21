package api

import (
	"ptv-tui/client"
	"ptv-tui/client/departures"
	"ptv-tui/models"
	"time"
)

// GetDepatures retrieves departure information for a given stop from the PTV API.
// It accepts a client and a slice of Parameters, which can include route IDs, stop IDs, and other options.
func GetDepatures(client client.Ptvclient, parameters []Parameters) []Responses {

	var params departures.DeparturesGetForStopParams
	var DepList []Responses

	for _, p := range parameters {

		if p.LookBackwards {
			params.SetLookBackwards(&p.LookBackwards)
		}
		if p.MaxResults != 0 {
			params.SetMaxResults(&p.MaxResults)
		}
		if p.DirectionID != nil {
			params.SetDirectionID(p.DirectionID)
		}
		if p.RouteTypes != nil {
			params.SetRouteType(p.RouteTypes[0])
		}
		if p.StopID != 0 {
			params.SetStopID(p.StopID)
		}

		depaturesResponses, err := client.Departures.DeparturesGetForStop(&params)
		if err != nil {

			return nil
		}
		//
		DepList = append(DepList, BuildDepaturesResponses(depaturesResponses.Payload)...)
	}
	return DepList
}

// BuildDepaturesResponses converts the V3DeparturesResponse payload into a slice of Responses.
// It formats the scheduled and estimated departure times into a human-readable format.
// If the payload is nil, it returns an empty slice.
// Each departure is represented by a Responses struct, which includes route ID, direction ID, stop ID, and formatted departure times.
func BuildDepaturesResponses(payload *models.V3DeparturesResponse) []Responses {
	if payload == nil {
		return nil
	}

	var Results []Responses
	for _, dep := range payload.Departures {
		resp := Responses{
			RouteID:     dep.RouteID,
			DirectionID: dep.DirectionID,
			StopID:      []int32{dep.StopID},
		}

		//

		if !time.Time(dep.ScheduledDepartureUtc).IsZero() {
			resp.SchDepTime = time.Time(dep.ScheduledDepartureUtc).Local().Format("Today 2 3:04 PM")
		}
		if !time.Time(dep.EstimatedDepartureUtc).IsZero() {
			resp.EstDepTime = time.Time(dep.EstimatedDepartureUtc).Local().Format("Today 2 3:04 PM")
		}

		Results = append(Results, resp)
	}
	return Results
}
