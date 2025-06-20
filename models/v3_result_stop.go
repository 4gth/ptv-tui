// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3ResultStop v3 result stop
//
// swagger:model V3.ResultStop
type V3ResultStop struct {

	// Transport mode identifier
	RouteType int32 `json:"route_type,omitempty"`

	// List of routes travelling through the stop
	Routes []*V3ResultRoute `json:"routes"`

	// Distance of stop from input location (in metres); returns 0 if no location is input
	StopDistance float32 `json:"stop_distance,omitempty"`

	// Stop identifier
	StopID int32 `json:"stop_id,omitempty"`

	// Landmark in proximity of stop
	StopLandmark string `json:"stop_landmark,omitempty"`

	// Geographic coordinate of latitude at stop
	StopLatitude float32 `json:"stop_latitude,omitempty"`

	// Geographic coordinate of longitude at stop
	StopLongitude float32 `json:"stop_longitude,omitempty"`

	// Name of stop
	StopName string `json:"stop_name,omitempty"`

	// Sequence of the stop on the route/run; return 0 when route_id or run_id not specified. Order ascendingly by this field (when non zero) to get physical order (earliest first) of stops on the route_id/run_id.
	StopSequence int32 `json:"stop_sequence,omitempty"`

	// suburb of stop
	StopSuburb string `json:"stop_suburb,omitempty"`
}

// Validate validates this v3 result stop
func (m *V3ResultStop) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoutes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3ResultStop) validateRoutes(formats strfmt.Registry) error {
	if swag.IsZero(m.Routes) { // not required
		return nil
	}

	for i := 0; i < len(m.Routes); i++ {
		if swag.IsZero(m.Routes[i]) { // not required
			continue
		}

		if m.Routes[i] != nil {
			if err := m.Routes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("routes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v3 result stop based on the context it is used
func (m *V3ResultStop) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoutes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3ResultStop) contextValidateRoutes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Routes); i++ {

		if m.Routes[i] != nil {

			if swag.IsZero(m.Routes[i]) { // not required
				return nil
			}

			if err := m.Routes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("routes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3ResultStop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3ResultStop) UnmarshalBinary(b []byte) error {
	var res V3ResultStop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
