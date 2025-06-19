package joinutil

import "ptv-tui/api"

// MergeToMap merges elements from a slice into a map, using a key function to determine the key for each element.
// If an element with the same key already exists in the map, it merges the existing value with the new value using a merge function.
// If the key does not exist, it adds a new entry to the map with a copy of the value.
// TODO: Intergrate this with the PTV API responses to enrich data

func MergeToMap[K comparable, V any](
	target map[K]*V,
	source []V,
	keyFunc func(V) K,
	mergeFunc func(dst *V, src V),
) {
	for _, item := range source {
		key := keyFunc(item)
		if existing, found := target[key]; found {
			mergeFunc(existing, item)
		} else {
			copy := item
			target[key] = &copy
		}
	}
}

// EnrichDepatures enriches a slice of departures with stop details and route information.
// It takes three slices: departures, stopDetails, and routes.
// It returns a new slice of departures with enriched information.
// Each departure is enriched with the stop name and route details (name and number) based on the StopID and RouteID.
// The function uses maps for quick lookups of stop details and routes
func EnrichDepatures(
	departures []api.Responses,
	stopDetails []api.Responses,
	routes []api.Responses,
) []api.Responses {
	// Create a map for quick lookup of stop details by StopID
	stopMap := make(map[int32]*api.Responses)
	for _, stop := range stopDetails {
		stopMap[stop.StopID] = &stop
	}

	// Create a map for quick lookup of routes by RouteID
	routeMap := make(map[int32]*api.Responses)
	for _, route := range routes {
		routeMap[route.RouteID] = &route
	}

	// Enrich departures with stop details and route information
	for i, dep := range departures {
		if stop, found := stopMap[dep.StopID]; found {
			departures[i].StopName = stop.StopName
		}
		if route, found := routeMap[dep.RouteID]; found {
			departures[i].Name = route.Name
			departures[i].Number = route.Number
		}
	}

	return departures
}
