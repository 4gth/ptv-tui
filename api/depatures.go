package api

import (
	"log"
	"ptv-tui/client"
	"ptv-tui/client/departures"
	"ptv-tui/models"
	"time"
)

// Takes a client and constructs a API call with params returns a struct
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
		log.Fatal(err)
	}

	log.Print(depaturesResponses.Code())
	payload := BuildDepaturesResponses(depaturesResponses.Payload)

	log.Println(depaturesResponses.Payload)
	return payload
}

// Takes the payload from the API call that returns a map[] and constructs a response struct
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
