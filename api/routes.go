package api

import (
	"ptv-tui/client"
	"ptv-tui/client/routes"
)

// GetRoutes retrieves route information from the PTV API based on the provided parameters.
// It accepts a client and a slice of Parameters, which can include route types, route names,
// and other options. The function constructs a request to the PTV API and returns a slice of
// Responses containing the route details.
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

	}
	routeList := BuildRoutesResponses(routeRespones)
	return routeList

}

// BuildRoutesResponses converts the RoutesOneOrMoreRoutesOK payload into a slice of Responses.
// It extracts relevant route information such as RouteID, RouteName, and RouteNumber.
// If the payload is nil, it returns an empty slice.
// Each route is represented by a Responses struct, which includes the route ID, name, and number.
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
