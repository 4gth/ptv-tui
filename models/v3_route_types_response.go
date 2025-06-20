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

// V3RouteTypesResponse v3 route types response
//
// swagger:model V3.RouteTypesResponse
type V3RouteTypesResponse struct {

	// Transport mode identifiers
	RouteTypes []*V3RouteType `json:"route_types"`

	// API Status / Metadata
	Status *V3Status `json:"status,omitempty"`
}

// Validate validates this v3 route types response
func (m *V3RouteTypesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRouteTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3RouteTypesResponse) validateRouteTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.RouteTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.RouteTypes); i++ {
		if swag.IsZero(m.RouteTypes[i]) { // not required
			continue
		}

		if m.RouteTypes[i] != nil {
			if err := m.RouteTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("route_types" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("route_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3RouteTypesResponse) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v3 route types response based on the context it is used
func (m *V3RouteTypesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRouteTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3RouteTypesResponse) contextValidateRouteTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RouteTypes); i++ {

		if m.RouteTypes[i] != nil {

			if swag.IsZero(m.RouteTypes[i]) { // not required
				return nil
			}

			if err := m.RouteTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("route_types" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("route_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3RouteTypesResponse) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3RouteTypesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3RouteTypesResponse) UnmarshalBinary(b []byte) error {
	var res V3RouteTypesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
