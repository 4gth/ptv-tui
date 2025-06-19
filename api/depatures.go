package api

import (
	"log"
	"ptv-tui/client"
	"ptv-tui/client/departures"
	"ptv-tui/models"
	"time"
)

// GetDepatures retrieves departure information for a given stop from the PTV API.
// It accepts a client and a slice of Parameters, which can include route IDs, stop IDs, and other options.
func GetDepatures(client client.Ptvclient, parameters []Parameters) []Responses {

	var params departures.DeparturesGetForStopParams

	for _, p := range parameters {
		if *p.RouteID != 0 {
			params.SetRouteType(*p.RouteID)
		}
		if p.StopID != 0 {
			params.SetStopID(p.StopID)
		}
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
			params.SetRouteType(p.RouteType)
		}
	}

	depaturesResponses, err := client.Departures.DeparturesGetForStop(&params)
	if err != nil {
		log.Printf("[%d]Failed to get Depatures: %s", depaturesResponses.Code(), err)
	}

	payload := BuildDepaturesResponses(depaturesResponses.Payload)

	return payload
}

// BuildDepaturesResponses converts the V3DeparturesResponse payload into a slice of Responses.
// It formats the scheduled and estimated departure times into a human-readable format.
// If the payload is nil, it returns an empty slice.
// Each departure is represented by a Responses struct, which includes route ID, direction ID, stop ID, and formatted departure times.
func BuildDepaturesResponses(payload *models.V3DeparturesResponse) []Responses {
	if payload == nil {
		return nil
	}

	var results []Responses

	for _, dep := range payload.Departures {
		resp := Responses{
			RouteID:     dep.RouteID,
			DirectionID: dep.DirectionID,
			StopID:      dep.StopID,
		}

		if !time.Time(dep.ScheduledDepartureUtc).IsZero() {
			resp.SchDepTime = time.Time(dep.ScheduledDepartureUtc).Local().Format("Mon 2 3:04 PM")
		}
		if !time.Time(dep.EstimatedDepartureUtc).IsZero() {
			resp.EstDepTime = time.Time(dep.EstimatedDepartureUtc).Local().Format("Mon 2 3:04 PM")
		}

		results = append(results, resp)
	}
	return results
}
