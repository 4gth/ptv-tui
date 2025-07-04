// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V3JourneyPlannerLocation v3 journey planner location
//
// swagger:model V3.JourneyPlannerLocation
type V3JourneyPlannerLocation struct {

	// disruption ids
	DisruptionIds []int64 `json:"DisruptionIds"`

	// lat
	Lat float32 `json:"Lat,omitempty"`

	// locality
	Locality string `json:"Locality,omitempty"`

	// location name
	LocationName string `json:"LocationName,omitempty"`

	// lon
	Lon float32 `json:"Lon,omitempty"`

	// place Id
	PlaceID int32 `json:"PlaceId,omitempty"`

	// platform
	Platform string `json:"Platform,omitempty"`

	// route types
	RouteTypes []int32 `json:"RouteTypes"`

	// stop Id
	StopID int32 `json:"StopId,omitempty"`

	// stop ticket
	StopTicket *V3StopTicket `json:"StopTicket,omitempty"`
}

// Validate validates this v3 journey planner location
func (m *V3JourneyPlannerLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRouteTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopTicket(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v3JourneyPlannerLocationRouteTypesItemsEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[0,1,2,3,4]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v3JourneyPlannerLocationRouteTypesItemsEnum = append(v3JourneyPlannerLocationRouteTypesItemsEnum, v)
	}
}

func (m *V3JourneyPlannerLocation) validateRouteTypesItemsEnum(path, location string, value int32) error {
	if err := validate.EnumCase(path, location, value, v3JourneyPlannerLocationRouteTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *V3JourneyPlannerLocation) validateRouteTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.RouteTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.RouteTypes); i++ {

		// value enum
		if err := m.validateRouteTypesItemsEnum("RouteTypes"+"."+strconv.Itoa(i), "body", m.RouteTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *V3JourneyPlannerLocation) validateStopTicket(formats strfmt.Registry) error {
	if swag.IsZero(m.StopTicket) { // not required
		return nil
	}

	if m.StopTicket != nil {
		if err := m.StopTicket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StopTicket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("StopTicket")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v3 journey planner location based on the context it is used
func (m *V3JourneyPlannerLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStopTicket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3JourneyPlannerLocation) contextValidateStopTicket(ctx context.Context, formats strfmt.Registry) error {

	if m.StopTicket != nil {

		if swag.IsZero(m.StopTicket) { // not required
			return nil
		}

		if err := m.StopTicket.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StopTicket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("StopTicket")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3JourneyPlannerLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3JourneyPlannerLocation) UnmarshalBinary(b []byte) error {
	var res V3JourneyPlannerLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
