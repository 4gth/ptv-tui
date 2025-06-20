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

// V3SearchResult v3 search result
//
// swagger:model V3.SearchResult
type V3SearchResult struct {

	// myki ticket outlets
	Outlets []*V3ResultOutlet `json:"outlets"`

	// Train lines, tram routes, bus routes, regional coach routes, Night Bus routes
	Routes []*V3ResultRoute `json:"routes"`

	// API Status / Metadata
	Status *V3Status `json:"status,omitempty"`

	// Train stations, tram stops, bus stops, regional coach stops or Night Bus stops
	Stops []*V3ResultStop `json:"stops"`
}

// Validate validates this v3 search result
func (m *V3SearchResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutlets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStops(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3SearchResult) validateOutlets(formats strfmt.Registry) error {
	if swag.IsZero(m.Outlets) { // not required
		return nil
	}

	for i := 0; i < len(m.Outlets); i++ {
		if swag.IsZero(m.Outlets[i]) { // not required
			continue
		}

		if m.Outlets[i] != nil {
			if err := m.Outlets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outlets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("outlets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3SearchResult) validateRoutes(formats strfmt.Registry) error {
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

func (m *V3SearchResult) validateStatus(formats strfmt.Registry) error {
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

func (m *V3SearchResult) validateStops(formats strfmt.Registry) error {
	if swag.IsZero(m.Stops) { // not required
		return nil
	}

	for i := 0; i < len(m.Stops); i++ {
		if swag.IsZero(m.Stops[i]) { // not required
			continue
		}

		if m.Stops[i] != nil {
			if err := m.Stops[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stops" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stops" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v3 search result based on the context it is used
func (m *V3SearchResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOutlets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoutes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3SearchResult) contextValidateOutlets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Outlets); i++ {

		if m.Outlets[i] != nil {

			if swag.IsZero(m.Outlets[i]) { // not required
				return nil
			}

			if err := m.Outlets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outlets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("outlets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3SearchResult) contextValidateRoutes(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V3SearchResult) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V3SearchResult) contextValidateStops(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Stops); i++ {

		if m.Stops[i] != nil {

			if swag.IsZero(m.Stops[i]) { // not required
				return nil
			}

			if err := m.Stops[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stops" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stops" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3SearchResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3SearchResult) UnmarshalBinary(b []byte) error {
	var res V3SearchResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
