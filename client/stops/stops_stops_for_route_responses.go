// Code generated by go-swagger; DO NOT EDIT.

package stops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"ptv-tui/models"
)

// StopsStopsForRouteReader is a Reader for the StopsStopsForRoute structure.
type StopsStopsForRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopsStopsForRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopsStopsForRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStopsStopsForRouteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopsStopsForRouteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v3/stops/route/{route_id}/route_type/{route_type}] Stops_StopsForRoute", response, response.Code())
	}
}

// NewStopsStopsForRouteOK creates a StopsStopsForRouteOK with default headers values
func NewStopsStopsForRouteOK() *StopsStopsForRouteOK {
	return &StopsStopsForRouteOK{}
}

/*
StopsStopsForRouteOK describes a response with status code 200, with default header values.

All stops on the specified route.
*/
type StopsStopsForRouteOK struct {
	Payload *models.V3StopsOnRouteResponse
}

// IsSuccess returns true when this stops stops for route o k response has a 2xx status code
func (o *StopsStopsForRouteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stops stops for route o k response has a 3xx status code
func (o *StopsStopsForRouteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stops stops for route o k response has a 4xx status code
func (o *StopsStopsForRouteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stops stops for route o k response has a 5xx status code
func (o *StopsStopsForRouteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stops stops for route o k response a status code equal to that given
func (o *StopsStopsForRouteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stops stops for route o k response
func (o *StopsStopsForRouteOK) Code() int {
	return 200
}

func (o *StopsStopsForRouteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v3/stops/route/{route_id}/route_type/{route_type}][%d] stopsStopsForRouteOK %s", 200, payload)
}

func (o *StopsStopsForRouteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v3/stops/route/{route_id}/route_type/{route_type}][%d] stopsStopsForRouteOK %s", 200, payload)
}

func (o *StopsStopsForRouteOK) GetPayload() *models.V3StopsOnRouteResponse {
	return o.Payload
}

func (o *StopsStopsForRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3StopsOnRouteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopsStopsForRouteBadRequest creates a StopsStopsForRouteBadRequest with default headers values
func NewStopsStopsForRouteBadRequest() *StopsStopsForRouteBadRequest {
	return &StopsStopsForRouteBadRequest{}
}

/*
StopsStopsForRouteBadRequest describes a response with status code 400, with default header values.

Invalid Request
*/
type StopsStopsForRouteBadRequest struct {
	Payload *models.V3ErrorResponse
}

// IsSuccess returns true when this stops stops for route bad request response has a 2xx status code
func (o *StopsStopsForRouteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stops stops for route bad request response has a 3xx status code
func (o *StopsStopsForRouteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stops stops for route bad request response has a 4xx status code
func (o *StopsStopsForRouteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stops stops for route bad request response has a 5xx status code
func (o *StopsStopsForRouteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stops stops for route bad request response a status code equal to that given
func (o *StopsStopsForRouteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stops stops for route bad request response
func (o *StopsStopsForRouteBadRequest) Code() int {
	return 400
}

func (o *StopsStopsForRouteBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v3/stops/route/{route_id}/route_type/{route_type}][%d] stopsStopsForRouteBadRequest %s", 400, payload)
}

func (o *StopsStopsForRouteBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v3/stops/route/{route_id}/route_type/{route_type}][%d] stopsStopsForRouteBadRequest %s", 400, payload)
}

func (o *StopsStopsForRouteBadRequest) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *StopsStopsForRouteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopsStopsForRouteForbidden creates a StopsStopsForRouteForbidden with default headers values
func NewStopsStopsForRouteForbidden() *StopsStopsForRouteForbidden {
	return &StopsStopsForRouteForbidden{}
}

/*
StopsStopsForRouteForbidden describes a response with status code 403, with default header values.

Access Denied
*/
type StopsStopsForRouteForbidden struct {
	Payload *models.V3ErrorResponse
}

// IsSuccess returns true when this stops stops for route forbidden response has a 2xx status code
func (o *StopsStopsForRouteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stops stops for route forbidden response has a 3xx status code
func (o *StopsStopsForRouteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stops stops for route forbidden response has a 4xx status code
func (o *StopsStopsForRouteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stops stops for route forbidden response has a 5xx status code
func (o *StopsStopsForRouteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stops stops for route forbidden response a status code equal to that given
func (o *StopsStopsForRouteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stops stops for route forbidden response
func (o *StopsStopsForRouteForbidden) Code() int {
	return 403
}

func (o *StopsStopsForRouteForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v3/stops/route/{route_id}/route_type/{route_type}][%d] stopsStopsForRouteForbidden %s", 403, payload)
}

func (o *StopsStopsForRouteForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v3/stops/route/{route_id}/route_type/{route_type}][%d] stopsStopsForRouteForbidden %s", 403, payload)
}

func (o *StopsStopsForRouteForbidden) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *StopsStopsForRouteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
