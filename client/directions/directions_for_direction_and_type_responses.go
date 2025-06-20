// Code generated by go-swagger; DO NOT EDIT.

package directions

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

// DirectionsForDirectionAndTypeReader is a Reader for the DirectionsForDirectionAndType structure.
type DirectionsForDirectionAndTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DirectionsForDirectionAndTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDirectionsForDirectionAndTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDirectionsForDirectionAndTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDirectionsForDirectionAndTypeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v3/directions/{direction_id}/route_type/{route_type}] Directions_ForDirectionAndType", response, response.Code())
	}
}

// NewDirectionsForDirectionAndTypeOK creates a DirectionsForDirectionAndTypeOK with default headers values
func NewDirectionsForDirectionAndTypeOK() *DirectionsForDirectionAndTypeOK {
	return &DirectionsForDirectionAndTypeOK{}
}

/*
DirectionsForDirectionAndTypeOK describes a response with status code 200, with default header values.

All routes of the specified route type that travel in the specified direction.
*/
type DirectionsForDirectionAndTypeOK struct {
	Payload *models.V3DirectionsResponse
}

// IsSuccess returns true when this directions for direction and type o k response has a 2xx status code
func (o *DirectionsForDirectionAndTypeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this directions for direction and type o k response has a 3xx status code
func (o *DirectionsForDirectionAndTypeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this directions for direction and type o k response has a 4xx status code
func (o *DirectionsForDirectionAndTypeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this directions for direction and type o k response has a 5xx status code
func (o *DirectionsForDirectionAndTypeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this directions for direction and type o k response a status code equal to that given
func (o *DirectionsForDirectionAndTypeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the directions for direction and type o k response
func (o *DirectionsForDirectionAndTypeOK) Code() int {
	return 200
}

func (o *DirectionsForDirectionAndTypeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v3/directions/{direction_id}/route_type/{route_type}][%d] directionsForDirectionAndTypeOK %s", 200, payload)
}

func (o *DirectionsForDirectionAndTypeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v3/directions/{direction_id}/route_type/{route_type}][%d] directionsForDirectionAndTypeOK %s", 200, payload)
}

func (o *DirectionsForDirectionAndTypeOK) GetPayload() *models.V3DirectionsResponse {
	return o.Payload
}

func (o *DirectionsForDirectionAndTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3DirectionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectionsForDirectionAndTypeBadRequest creates a DirectionsForDirectionAndTypeBadRequest with default headers values
func NewDirectionsForDirectionAndTypeBadRequest() *DirectionsForDirectionAndTypeBadRequest {
	return &DirectionsForDirectionAndTypeBadRequest{}
}

/*
DirectionsForDirectionAndTypeBadRequest describes a response with status code 400, with default header values.

Invalid Request
*/
type DirectionsForDirectionAndTypeBadRequest struct {
	Payload *models.V3ErrorResponse
}

// IsSuccess returns true when this directions for direction and type bad request response has a 2xx status code
func (o *DirectionsForDirectionAndTypeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this directions for direction and type bad request response has a 3xx status code
func (o *DirectionsForDirectionAndTypeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this directions for direction and type bad request response has a 4xx status code
func (o *DirectionsForDirectionAndTypeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this directions for direction and type bad request response has a 5xx status code
func (o *DirectionsForDirectionAndTypeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this directions for direction and type bad request response a status code equal to that given
func (o *DirectionsForDirectionAndTypeBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the directions for direction and type bad request response
func (o *DirectionsForDirectionAndTypeBadRequest) Code() int {
	return 400
}

func (o *DirectionsForDirectionAndTypeBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v3/directions/{direction_id}/route_type/{route_type}][%d] directionsForDirectionAndTypeBadRequest %s", 400, payload)
}

func (o *DirectionsForDirectionAndTypeBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v3/directions/{direction_id}/route_type/{route_type}][%d] directionsForDirectionAndTypeBadRequest %s", 400, payload)
}

func (o *DirectionsForDirectionAndTypeBadRequest) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *DirectionsForDirectionAndTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectionsForDirectionAndTypeForbidden creates a DirectionsForDirectionAndTypeForbidden with default headers values
func NewDirectionsForDirectionAndTypeForbidden() *DirectionsForDirectionAndTypeForbidden {
	return &DirectionsForDirectionAndTypeForbidden{}
}

/*
DirectionsForDirectionAndTypeForbidden describes a response with status code 403, with default header values.

Access Denied
*/
type DirectionsForDirectionAndTypeForbidden struct {
	Payload *models.V3ErrorResponse
}

// IsSuccess returns true when this directions for direction and type forbidden response has a 2xx status code
func (o *DirectionsForDirectionAndTypeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this directions for direction and type forbidden response has a 3xx status code
func (o *DirectionsForDirectionAndTypeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this directions for direction and type forbidden response has a 4xx status code
func (o *DirectionsForDirectionAndTypeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this directions for direction and type forbidden response has a 5xx status code
func (o *DirectionsForDirectionAndTypeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this directions for direction and type forbidden response a status code equal to that given
func (o *DirectionsForDirectionAndTypeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the directions for direction and type forbidden response
func (o *DirectionsForDirectionAndTypeForbidden) Code() int {
	return 403
}

func (o *DirectionsForDirectionAndTypeForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v3/directions/{direction_id}/route_type/{route_type}][%d] directionsForDirectionAndTypeForbidden %s", 403, payload)
}

func (o *DirectionsForDirectionAndTypeForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v3/directions/{direction_id}/route_type/{route_type}][%d] directionsForDirectionAndTypeForbidden %s", 403, payload)
}

func (o *DirectionsForDirectionAndTypeForbidden) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *DirectionsForDirectionAndTypeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
