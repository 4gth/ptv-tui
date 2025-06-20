// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3StopAmenityDetails v3 stop amenity details
//
// swagger:model V3.StopAmenityDetails
type V3StopAmenityDetails struct {

	// The number of free car parking spots at the stop
	CarParking string `json:"car_parking,omitempty"`

	// Indicates if there are CCTV (i.e. closed circuit television) cameras at the stop
	Cctv bool `json:"cctv,omitempty"`

	// Indicates if there is a taxi rank at or near the stop
	TaxiRank bool `json:"taxi_rank,omitempty"`

	// Indicates if there is a public toilet at or near the stop
	Toilet bool `json:"toilet,omitempty"`
}

// Validate validates this v3 stop amenity details
func (m *V3StopAmenityDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v3 stop amenity details based on context it is used
func (m *V3StopAmenityDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V3StopAmenityDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3StopAmenityDetails) UnmarshalBinary(b []byte) error {
	var res V3StopAmenityDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
