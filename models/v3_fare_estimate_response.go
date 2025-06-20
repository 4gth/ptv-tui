// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3FareEstimateResponse v3 fare estimate response
//
// swagger:model V3.FareEstimateResponse
type V3FareEstimateResponse struct {

	// fare estimate result
	FareEstimateResult *V3FareEstimateResult `json:"FareEstimateResult,omitempty"`

	// fare estimate result status
	FareEstimateResultStatus *V3FareEstimateResultStatus `json:"FareEstimateResultStatus,omitempty"`
}

// Validate validates this v3 fare estimate response
func (m *V3FareEstimateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFareEstimateResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFareEstimateResultStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3FareEstimateResponse) validateFareEstimateResult(formats strfmt.Registry) error {
	if swag.IsZero(m.FareEstimateResult) { // not required
		return nil
	}

	if m.FareEstimateResult != nil {
		if err := m.FareEstimateResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FareEstimateResult")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FareEstimateResult")
			}
			return err
		}
	}

	return nil
}

func (m *V3FareEstimateResponse) validateFareEstimateResultStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.FareEstimateResultStatus) { // not required
		return nil
	}

	if m.FareEstimateResultStatus != nil {
		if err := m.FareEstimateResultStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FareEstimateResultStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FareEstimateResultStatus")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v3 fare estimate response based on the context it is used
func (m *V3FareEstimateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFareEstimateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFareEstimateResultStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3FareEstimateResponse) contextValidateFareEstimateResult(ctx context.Context, formats strfmt.Registry) error {

	if m.FareEstimateResult != nil {

		if swag.IsZero(m.FareEstimateResult) { // not required
			return nil
		}

		if err := m.FareEstimateResult.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FareEstimateResult")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FareEstimateResult")
			}
			return err
		}
	}

	return nil
}

func (m *V3FareEstimateResponse) contextValidateFareEstimateResultStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.FareEstimateResultStatus != nil {

		if swag.IsZero(m.FareEstimateResultStatus) { // not required
			return nil
		}

		if err := m.FareEstimateResultStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FareEstimateResultStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FareEstimateResultStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3FareEstimateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3FareEstimateResponse) UnmarshalBinary(b []byte) error {
	var res V3FareEstimateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
