// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3FareEstimateResultStatus v3 fare estimate result status
//
// swagger:model V3.FareEstimateResultStatus
type V3FareEstimateResultStatus struct {

	// message
	Message string `json:"Message,omitempty"`

	// status code
	StatusCode int32 `json:"StatusCode,omitempty"`
}

// Validate validates this v3 fare estimate result status
func (m *V3FareEstimateResultStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v3 fare estimate result status based on context it is used
func (m *V3FareEstimateResultStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V3FareEstimateResultStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3FareEstimateResultStatus) UnmarshalBinary(b []byte) error {
	var res V3FareEstimateResultStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
