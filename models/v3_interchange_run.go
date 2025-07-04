// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3InterchangeRun Feeder / Distributor details
//
// swagger:model V3.InterchangeRun
type V3InterchangeRun struct {

	// Indicates whether the interchange information is shown to end users
	Advertised bool `json:"advertised,omitempty"`

	// Indicates the destination name
	DestinationName string `json:"destination_name,omitempty"`

	// Indicates whether the direction for this run
	DirectionID int32 `json:"direction_id,omitempty"`

	// Route identifier
	RouteID int32 `json:"route_id,omitempty"`

	// Run Identifier
	RunRef string `json:"run_ref,omitempty"`

	// Stop identifier
	StopID int32 `json:"stop_id,omitempty"`
}

// Validate validates this v3 interchange run
func (m *V3InterchangeRun) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v3 interchange run based on context it is used
func (m *V3InterchangeRun) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V3InterchangeRun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3InterchangeRun) UnmarshalBinary(b []byte) error {
	var res V3InterchangeRun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
