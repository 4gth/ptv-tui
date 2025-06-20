// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V3StopResponse v3 stop response
//
// swagger:model V3.StopResponse
type V3StopResponse struct {

	// Disruption information applicable to relevant routes or stops
	Disruptions map[string]V3Disruption `json:"disruptions,omitempty"`

	// API Status / Metadata
	Status *V3Status `json:"status,omitempty"`

	// A metropolitan or V/Line train station
	Stop *V3StopDetails `json:"stop,omitempty"`
}

// Validate validates this v3 stop response
func (m *V3StopResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisruptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStop(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3StopResponse) validateDisruptions(formats strfmt.Registry) error {
	if swag.IsZero(m.Disruptions) { // not required
		return nil
	}

	for k := range m.Disruptions {

		if err := validate.Required("disruptions"+"."+k, "body", m.Disruptions[k]); err != nil {
			return err
		}
		if val, ok := m.Disruptions[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disruptions" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disruptions" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3StopResponse) validateStatus(formats strfmt.Registry) error {
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

func (m *V3StopResponse) validateStop(formats strfmt.Registry) error {
	if swag.IsZero(m.Stop) { // not required
		return nil
	}

	if m.Stop != nil {
		if err := m.Stop.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stop")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stop")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v3 stop response based on the context it is used
func (m *V3StopResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisruptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStop(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3StopResponse) contextValidateDisruptions(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Disruptions {

		if val, ok := m.Disruptions[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *V3StopResponse) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V3StopResponse) contextValidateStop(ctx context.Context, formats strfmt.Registry) error {

	if m.Stop != nil {

		if swag.IsZero(m.Stop) { // not required
			return nil
		}

		if err := m.Stop.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stop")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stop")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3StopResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3StopResponse) UnmarshalBinary(b []byte) error {
	var res V3StopResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
