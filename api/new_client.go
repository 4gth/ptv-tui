package api

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"log"
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

type Responses struct {
	Name        string `col:"Name"`
	Number      string `col:"Number"`
	RouteID     int32  `col:"Route ID"`
	SchDepTime  string `col:"Scheduled Depature Time"`
	EstDepTime  string `col:"Estimated Depature Time"`
	DirectionID int32  `col:"Direction"`
	StopID      int32  `col:"Stop"`
	StopName    string `col:"Stop Name"`
	RouteType   int32  `col:"Route Type"`
	Routes      []any  `col:"Routes"`
}
type PTVAuth struct {
	DevID  string
	APIKey string
}

type Parameters struct {
	RouteTypes    []int32 `json:"route_types"`
	RouteID       *int32  `json:"route_id"`
	RouteName     *string `json:"route_name"`
	RouteNumber   *string `json:"route_number"`
	RouteType     int32   `json:"route_type"`
	DirectionID   *int32  `json:"direction_id"`
	StopID        int32   `json:"stop_id"`
	LookBackwards bool    `json:"look_backwards"`
	MaxResults    int32   `json:"max_results"`
}

func NewPTVClient() client.Ptvclient {

	// Load keys into env
	_ = godotenv.Load()
	auth := &PTVAuth{

		DevID:  os.Getenv("PTV_DEV_ID"),
		APIKey: os.Getenv("PTV_API_KEY"),
	}

	//Creat client transport with customer authentication writer.
	transport := httptransport.New(client.DefaultHost, client.DefaultBasePath, client.DefaultSchemes)
	transport.DefaultAuthentication = auth.NewAuthInfoWriter()

	//Create ptv client with default host and scheme with authwriter
	ptv := client.NewHTTPClient(nil)
	ptv.SetTransport(transport)

	return *ptv

}

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
	log.Printf("StringToSign: %s", stringToSign)
	mac := hmac.New(sha1.New, []byte(apiKey))
	mac.Write([]byte(stringToSign))
	signature := hex.EncodeToString(mac.Sum(nil))

	safeQuery.Set("signature", signature)
	log.Printf("safeQuery: %s", safeQuery)
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
		log.Printf("Calling endpoint: %s", req.GetQueryParams().Encode())

		return nil
	})
}
