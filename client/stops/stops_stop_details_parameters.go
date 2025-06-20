// Code generated by go-swagger; DO NOT EDIT.

package stops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewStopsStopDetailsParams creates a new StopsStopDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopsStopDetailsParams() *StopsStopDetailsParams {
	return &StopsStopDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopsStopDetailsParamsWithTimeout creates a new StopsStopDetailsParams object
// with the ability to set a timeout on a request.
func NewStopsStopDetailsParamsWithTimeout(timeout time.Duration) *StopsStopDetailsParams {
	return &StopsStopDetailsParams{
		timeout: timeout,
	}
}

// NewStopsStopDetailsParamsWithContext creates a new StopsStopDetailsParams object
// with the ability to set a context for a request.
func NewStopsStopDetailsParamsWithContext(ctx context.Context) *StopsStopDetailsParams {
	return &StopsStopDetailsParams{
		Context: ctx,
	}
}

// NewStopsStopDetailsParamsWithHTTPClient creates a new StopsStopDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopsStopDetailsParamsWithHTTPClient(client *http.Client) *StopsStopDetailsParams {
	return &StopsStopDetailsParams{
		HTTPClient: client,
	}
}

/*
StopsStopDetailsParams contains all the parameters to send to the API endpoint

	for the stops stop details operation.

	Typically these are written to a http.Request.
*/
type StopsStopDetailsParams struct {

	/* Devid.

	   Your developer id
	*/
	Devid *string

	/* Gtfs.

	   Incdicates whether the stop_id is a GTFS ID or not
	*/
	Gtfs *bool

	/* RouteType.

	   Number identifying transport mode; values returned via RouteTypes API

	   Format: int32
	*/
	RouteType int32

	/* Signature.

	   Authentication signature for request
	*/
	Signature *string

	/* StopAccessibility.

	   Indicates if stop accessibility information will be returned (default = false)
	*/
	StopAccessibility *bool

	/* StopAmenities.

	   Indicates if stop amenity information will be returned (default = false)
	*/
	StopAmenities *bool

	/* StopContact.

	   Indicates if stop contact information will be returned (default = false)
	*/
	StopContact *bool

	/* StopDisruptions.

	   Indicates if stop disruption information will be returned (default = false)
	*/
	StopDisruptions *bool

	/* StopID.

	   Identifier of stop; values returned by Stops API

	   Format: int32
	*/
	StopID int32

	/* StopLocation.

	   Indicates if stop location information will be returned (default = false)
	*/
	StopLocation *bool

	/* StopStaffing.

	   Indicates if stop staffing information will be returned (default = false)
	*/
	StopStaffing *bool

	/* StopTicket.

	   Indicates if stop ticket information will be returned (default = false)
	*/
	StopTicket *bool

	/* Token.

	   Please ignore
	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stops stop details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopsStopDetailsParams) WithDefaults() *StopsStopDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stops stop details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopsStopDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stops stop details params
func (o *StopsStopDetailsParams) WithTimeout(timeout time.Duration) *StopsStopDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stops stop details params
func (o *StopsStopDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stops stop details params
func (o *StopsStopDetailsParams) WithContext(ctx context.Context) *StopsStopDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stops stop details params
func (o *StopsStopDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stops stop details params
func (o *StopsStopDetailsParams) WithHTTPClient(client *http.Client) *StopsStopDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stops stop details params
func (o *StopsStopDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevid adds the devid to the stops stop details params
func (o *StopsStopDetailsParams) WithDevid(devid *string) *StopsStopDetailsParams {
	o.SetDevid(devid)
	return o
}

// SetDevid adds the devid to the stops stop details params
func (o *StopsStopDetailsParams) SetDevid(devid *string) {
	o.Devid = devid
}

// WithGtfs adds the gtfs to the stops stop details params
func (o *StopsStopDetailsParams) WithGtfs(gtfs *bool) *StopsStopDetailsParams {
	o.SetGtfs(gtfs)
	return o
}

// SetGtfs adds the gtfs to the stops stop details params
func (o *StopsStopDetailsParams) SetGtfs(gtfs *bool) {
	o.Gtfs = gtfs
}

// WithRouteType adds the routeType to the stops stop details params
func (o *StopsStopDetailsParams) WithRouteType(routeType int32) *StopsStopDetailsParams {
	o.SetRouteType(routeType)
	return o
}

// SetRouteType adds the routeType to the stops stop details params
func (o *StopsStopDetailsParams) SetRouteType(routeType int32) {
	o.RouteType = routeType
}

// WithSignature adds the signature to the stops stop details params
func (o *StopsStopDetailsParams) WithSignature(signature *string) *StopsStopDetailsParams {
	o.SetSignature(signature)
	return o
}

// SetSignature adds the signature to the stops stop details params
func (o *StopsStopDetailsParams) SetSignature(signature *string) {
	o.Signature = signature
}

// WithStopAccessibility adds the stopAccessibility to the stops stop details params
func (o *StopsStopDetailsParams) WithStopAccessibility(stopAccessibility *bool) *StopsStopDetailsParams {
	o.SetStopAccessibility(stopAccessibility)
	return o
}

// SetStopAccessibility adds the stopAccessibility to the stops stop details params
func (o *StopsStopDetailsParams) SetStopAccessibility(stopAccessibility *bool) {
	o.StopAccessibility = stopAccessibility
}

// WithStopAmenities adds the stopAmenities to the stops stop details params
func (o *StopsStopDetailsParams) WithStopAmenities(stopAmenities *bool) *StopsStopDetailsParams {
	o.SetStopAmenities(stopAmenities)
	return o
}

// SetStopAmenities adds the stopAmenities to the stops stop details params
func (o *StopsStopDetailsParams) SetStopAmenities(stopAmenities *bool) {
	o.StopAmenities = stopAmenities
}

// WithStopContact adds the stopContact to the stops stop details params
func (o *StopsStopDetailsParams) WithStopContact(stopContact *bool) *StopsStopDetailsParams {
	o.SetStopContact(stopContact)
	return o
}

// SetStopContact adds the stopContact to the stops stop details params
func (o *StopsStopDetailsParams) SetStopContact(stopContact *bool) {
	o.StopContact = stopContact
}

// WithStopDisruptions adds the stopDisruptions to the stops stop details params
func (o *StopsStopDetailsParams) WithStopDisruptions(stopDisruptions *bool) *StopsStopDetailsParams {
	o.SetStopDisruptions(stopDisruptions)
	return o
}

// SetStopDisruptions adds the stopDisruptions to the stops stop details params
func (o *StopsStopDetailsParams) SetStopDisruptions(stopDisruptions *bool) {
	o.StopDisruptions = stopDisruptions
}

// WithStopID adds the stopID to the stops stop details params
func (o *StopsStopDetailsParams) WithStopID(stopID int32) *StopsStopDetailsParams {
	o.SetStopID(stopID)
	return o
}

// SetStopID adds the stopId to the stops stop details params
func (o *StopsStopDetailsParams) SetStopID(stopID int32) {
	o.StopID = stopID
}

// WithStopLocation adds the stopLocation to the stops stop details params
func (o *StopsStopDetailsParams) WithStopLocation(stopLocation *bool) *StopsStopDetailsParams {
	o.SetStopLocation(stopLocation)
	return o
}

// SetStopLocation adds the stopLocation to the stops stop details params
func (o *StopsStopDetailsParams) SetStopLocation(stopLocation *bool) {
	o.StopLocation = stopLocation
}

// WithStopStaffing adds the stopStaffing to the stops stop details params
func (o *StopsStopDetailsParams) WithStopStaffing(stopStaffing *bool) *StopsStopDetailsParams {
	o.SetStopStaffing(stopStaffing)
	return o
}

// SetStopStaffing adds the stopStaffing to the stops stop details params
func (o *StopsStopDetailsParams) SetStopStaffing(stopStaffing *bool) {
	o.StopStaffing = stopStaffing
}

// WithStopTicket adds the stopTicket to the stops stop details params
func (o *StopsStopDetailsParams) WithStopTicket(stopTicket *bool) *StopsStopDetailsParams {
	o.SetStopTicket(stopTicket)
	return o
}

// SetStopTicket adds the stopTicket to the stops stop details params
func (o *StopsStopDetailsParams) SetStopTicket(stopTicket *bool) {
	o.StopTicket = stopTicket
}

// WithToken adds the token to the stops stop details params
func (o *StopsStopDetailsParams) WithToken(token *string) *StopsStopDetailsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the stops stop details params
func (o *StopsStopDetailsParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *StopsStopDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Devid != nil {

		// query param devid
		var qrDevid string

		if o.Devid != nil {
			qrDevid = *o.Devid
		}
		qDevid := qrDevid
		if qDevid != "" {

			if err := r.SetQueryParam("devid", qDevid); err != nil {
				return err
			}
		}
	}

	if o.Gtfs != nil {

		// query param gtfs
		var qrGtfs bool

		if o.Gtfs != nil {
			qrGtfs = *o.Gtfs
		}
		qGtfs := swag.FormatBool(qrGtfs)
		if qGtfs != "" {

			if err := r.SetQueryParam("gtfs", qGtfs); err != nil {
				return err
			}
		}
	}

	// path param route_type
	if err := r.SetPathParam("route_type", swag.FormatInt32(o.RouteType)); err != nil {
		return err
	}

	if o.Signature != nil {

		// query param signature
		var qrSignature string

		if o.Signature != nil {
			qrSignature = *o.Signature
		}
		qSignature := qrSignature
		if qSignature != "" {

			if err := r.SetQueryParam("signature", qSignature); err != nil {
				return err
			}
		}
	}

	if o.StopAccessibility != nil {

		// query param stop_accessibility
		var qrStopAccessibility bool

		if o.StopAccessibility != nil {
			qrStopAccessibility = *o.StopAccessibility
		}
		qStopAccessibility := swag.FormatBool(qrStopAccessibility)
		if qStopAccessibility != "" {

			if err := r.SetQueryParam("stop_accessibility", qStopAccessibility); err != nil {
				return err
			}
		}
	}

	if o.StopAmenities != nil {

		// query param stop_amenities
		var qrStopAmenities bool

		if o.StopAmenities != nil {
			qrStopAmenities = *o.StopAmenities
		}
		qStopAmenities := swag.FormatBool(qrStopAmenities)
		if qStopAmenities != "" {

			if err := r.SetQueryParam("stop_amenities", qStopAmenities); err != nil {
				return err
			}
		}
	}

	if o.StopContact != nil {

		// query param stop_contact
		var qrStopContact bool

		if o.StopContact != nil {
			qrStopContact = *o.StopContact
		}
		qStopContact := swag.FormatBool(qrStopContact)
		if qStopContact != "" {

			if err := r.SetQueryParam("stop_contact", qStopContact); err != nil {
				return err
			}
		}
	}

	if o.StopDisruptions != nil {

		// query param stop_disruptions
		var qrStopDisruptions bool

		if o.StopDisruptions != nil {
			qrStopDisruptions = *o.StopDisruptions
		}
		qStopDisruptions := swag.FormatBool(qrStopDisruptions)
		if qStopDisruptions != "" {

			if err := r.SetQueryParam("stop_disruptions", qStopDisruptions); err != nil {
				return err
			}
		}
	}

	// path param stop_id
	if err := r.SetPathParam("stop_id", swag.FormatInt32(o.StopID)); err != nil {
		return err
	}

	if o.StopLocation != nil {

		// query param stop_location
		var qrStopLocation bool

		if o.StopLocation != nil {
			qrStopLocation = *o.StopLocation
		}
		qStopLocation := swag.FormatBool(qrStopLocation)
		if qStopLocation != "" {

			if err := r.SetQueryParam("stop_location", qStopLocation); err != nil {
				return err
			}
		}
	}

	if o.StopStaffing != nil {

		// query param stop_staffing
		var qrStopStaffing bool

		if o.StopStaffing != nil {
			qrStopStaffing = *o.StopStaffing
		}
		qStopStaffing := swag.FormatBool(qrStopStaffing)
		if qStopStaffing != "" {

			if err := r.SetQueryParam("stop_staffing", qStopStaffing); err != nil {
				return err
			}
		}
	}

	if o.StopTicket != nil {

		// query param stop_ticket
		var qrStopTicket bool

		if o.StopTicket != nil {
			qrStopTicket = *o.StopTicket
		}
		qStopTicket := swag.FormatBool(qrStopTicket)
		if qStopTicket != "" {

			if err := r.SetQueryParam("stop_ticket", qStopTicket); err != nil {
				return err
			}
		}
	}

	if o.Token != nil {

		// query param token
		var qrToken string

		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken
		if qToken != "" {

			if err := r.SetQueryParam("token", qToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
