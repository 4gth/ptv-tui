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

// V3StopToStopDisruptionsResponse v3 stop to stop disruptions response
//
// swagger:model V3.StopToStopDisruptionsResponse
type V3StopToStopDisruptionsResponse struct {

	// disruptions
	Disruptions []*V3StopToStopDisruption `json:"disruptions"`

	// API Status / Metadata
	Status *V3Status `json:"status,omitempty"`
}

// Validate validates this v3 stop to stop disruptions response
func (m *V3StopToStopDisruptionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisruptions(formats); err != nil {
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

func (m *V3StopToStopDisruptionsResponse) validateDisruptions(formats strfmt.Registry) error {
	if swag.IsZero(m.Disruptions) { // not required
		return nil
	}

	for i := 0; i < len(m.Disruptions); i++ {
		if swag.IsZero(m.Disruptions[i]) { // not required
			continue
		}

		if m.Disruptions[i] != nil {
			if err := m.Disruptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disruptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disruptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3StopToStopDisruptionsResponse) validateStatus(formats strfmt.Registry) error {
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

// ContextValidate validate this v3 stop to stop disruptions response based on the context it is used
func (m *V3StopToStopDisruptionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisruptions(ctx, formats); err != nil {
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

func (m *V3StopToStopDisruptionsResponse) contextValidateDisruptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Disruptions); i++ {

		if m.Disruptions[i] != nil {

			if swag.IsZero(m.Disruptions[i]) { // not required
				return nil
			}

			if err := m.Disruptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disruptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disruptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3StopToStopDisruptionsResponse) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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
func (m *V3StopToStopDisruptionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3StopToStopDisruptionsResponse) UnmarshalBinary(b []byte) error {
	var res V3StopToStopDisruptionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
