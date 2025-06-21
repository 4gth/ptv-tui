// Package api provides functions to interact with the PTV API.
// It includes client initialization, query signing, and data retrieval functions.
// It uses the go-openapi runtime for HTTP client operations and handles authentication via HMAC-SHA1.
// It also includes a structure for responses and parameters for API queries.

package api

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"net/url"
	"os"
	"ptv-tui/client"
	"sort"
	"strings"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/joho/godotenv"
)

// Responses represents the structure of the responses from the PTV API.
type Responses struct {
	Name          string  `col:"Name"`
	Number        string  `col:"Number"`
	RouteID       int32   `col:"Route ID"`
	SchDepTime    string  `col:"Scheduled Depature Time"`
	EstDepTime    string  `col:"Estimated Depature Time"`
	DirectionID   int32   `col:"Direction"`
	DirectionName string  `col:"Direction"`
	StopID        []int32 `col:"Stop"`
	StopName      string  `col:"Stop Name"`
	RouteType     int32   `col:"Route Type"`
	Routes        []any   `col:"Routes"`
	Distance      float32 `col:"Distance"` // Distance in meters
}

// PTVAuth holds the developer ID and API key for authenticating with the PTV API.
type PTVAuth struct {
	DevID  string
	APIKey string
}

// Parameters defines the parameters for querying the PTV API.
type Parameters struct {
	RouteTypes            []int32 `json:"route_types"`
	RouteID               *int32  `json:"route_id"`
	RouteName             *string `json:"route_name"`
	RouteNumber           *string `json:"route_number"`
	RouteType             int32   `json:"route_type"`
	DirectionID           *int32  `json:"direction_id"`
	StopID                int32   `json:"stop_id"`
	StopIDs               []int32 `json:"stop_ids"`
	LookBackwards         bool    `json:"look_backwards"`
	MaxResults            int32   `json:"max_results"`
	SearchTerm            string  `json:"search_term"`
	Latitude              float32 `json:"latitude"`
	Longitude             float32 `json:"longitude"`
	MaxDistance           float64 `json:"max_distance"`
	MatchStopByGtfsStopID bool    `json:"match_stop_by_gtfs_stop_id"`
}

// NewPTVClient initializes a new PTV client with authentication.
func NewPTVClient() client.Ptvclient {
	// Load environment variables from .env file
	// This is useful for local development to avoid hardcoding sensitive information.
	// In production, you might want to set these variables in your environment directly.
	_ = godotenv.Load()
	auth := &PTVAuth{

		DevID:  os.Getenv("PTV_DEV_ID"),
		APIKey: os.Getenv("PTV_API_KEY"),
	}
	// Create a new HTTP transport with the default host, base path, and schemes
	// The transport will use the custom authentication writer to sign requests.
	transport := httptransport.New(client.DefaultHost, client.DefaultBasePath, client.DefaultSchemes)
	transport.DefaultAuthentication = auth.NewAuthInfoWriter()

	// Set the transport to use the custom authentication writer
	ptv := client.NewHTTPClient(nil)
	ptv.SetTransport(transport)

	return *ptv

}

// SignQuery signs the query parameters for the PTV API using HMAC-SHA1 with the provided API key and developer ID.
// It returns the signed query parameters as url.Values.
// The function sorts the query parameters, constructs a string to sign, and appends the signature
// to the query parameters before returning them.
// It is used to authenticate requests to the PTV API.
func SignQuery(path string, query url.Values) url.Values {
	apiKey := os.Getenv("PTV_API_KEY")
	devID := os.Getenv("PTV_DEV_ID")

	//clone url to make safe (dont change original)
	safeQuery := url.Values{}
	for k, vs := range query {
		for _, v := range vs {
			safeQuery.Add(k, v)
		}
	}
	//add devid
	safeQuery.Set("devid", devID)

	//for each k make string
	keys := make([]string, 0, len(safeQuery))
	for k := range safeQuery {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	//build string
	var queryParts []string
	for _, k := range keys {
		for _, v := range safeQuery[k] {
			queryParts = append(queryParts, url.QueryEscape(k)+"="+url.QueryEscape(v))
		}
	}
	sortedQuery := strings.Join(queryParts, "&")
	stringToSign := path + "?" + sortedQuery

	mac := hmac.New(sha1.New, []byte(apiKey))
	mac.Write([]byte(stringToSign))
	signature := hex.EncodeToString(mac.Sum(nil))

	safeQuery.Set("signature", signature)

	return safeQuery

}

func (a *PTVAuth) NewAuthInfoWriter() runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(req runtime.ClientRequest, reg strfmt.Registry) error {

		origURL := req.GetPath()
		query := req.GetQueryParams()

		signedParams := SignQuery(origURL, query)

		for k, vs := range signedParams {
			for _, v := range vs {
				req.SetQueryParam(k, v)
			}
		}

		return nil
	})
}
