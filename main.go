package main

import (
	"fmt"
	"log"
	"ptv-tui/api"
	"ptv-tui/client/departures"
	"ptv-tui/ui"
	"time"
)

func main() {

	ptv := api.NewPTVClient()
	var routeList []ui.RouteRow
	//params := routes.NewRoutesOneOrMoreRoutesParams()
	//params.SetRouteTypes([]int32{1})
	//routeRespones, err := ptv.Routes.RoutesOneOrMoreRoutes(params)
	//if err != nil {
	//	log.Fatal(err)
	//}
	maxResults := int32(5)
	lookBack := false

	dparams := departures.NewDeparturesGetForStopParams()
	dparams.SetRouteType(int32(1))
	dparams.SetStopID(int32(2266))
	dparams.SetMaxResults(&maxResults)
	dparams.SetLookBackwards(&lookBack)
	depaturesResponses, err := ptv.Departures.DeparturesGetForStop(dparams)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(depaturesResponses.Code())

	for _, d := range depaturesResponses.Payload.Departures {
		localTime := time.Time(d.EstimatedDepartureUtc).Local().Format("Mon 2 3:04 PM")
		routeList = append(routeList, ui.RouteRow{
			Name:    d.DirectionID,
			Number:  localTime,
			RouteID: d.RouteID})
		log.Printf("RouteID: %d", d.RouteID)
	}

	for _, dep := range depaturesResponses.Payload.Departures {
		fmt.Printf("Run: %s, Route: %d, Departure Time: %s\n", dep.RunRef, dep.DirectionID, dep.ScheduledDepartureUtc)
	}

	//
	//for _, r := range routeRespones.Payload.Routes {
	//	routeList = append(routeList, ui.RouteRow{
	//		Name:    r.RouteName,
	//		Number:  r.RouteNumber,
	//		RouteID: r.RouteID})
	//}
	//
	if err := ui.StartNewTea(routeList); err != nil {
		log.Fatal(err)
	}

	//	params := routes.NewRoutesOneOrMoreRoutesParams()
	//	params.SetRouteTypes(nil)

}
