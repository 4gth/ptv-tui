package api

import (
	"log"
	"ptv-tui/client"
	"ptv-tui/client/routes"
)

func GetRoutes(client client.Ptvclient, parameters []Parameters) []Responses {
	var params routes.RoutesOneOrMoreRoutesParams

	for _, p := range parameters {
		if p.RouteTypes != nil {
			params.SetRouteTypes(p.RouteTypes)
		}
		if p.RouteName != nil {
			params.SetRouteName(p.RouteName)
		}
	}

	routeRespones, err := client.Routes.RoutesOneOrMoreRoutes(&params)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(routeRespones.Code())
	routeList := BuildRoutesResponses(routeRespones)

	log.Println(routeRespones.Payload)
	return routeList

}

func BuildRoutesResponses(payload *routes.RoutesOneOrMoreRoutesOK) []Responses {
	if payload == nil {
		return nil
	}

	var results []Responses

	for _, route := range payload.Payload.Routes {
		resp := Responses{
			RouteID:   route.RouteID,
			Name:      route.RouteName,
			Number:    route.RouteNumber,
			RouteType: route.RouteType,
		}

		results = append(results, resp)
	}
	return results
}
