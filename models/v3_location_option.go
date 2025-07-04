// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3LocationOption v3 location option
//
// swagger:model V3.LocationOption
type V3LocationOption struct {

	// name
	Name string `json:"Name,omitempty"`

	// Url
	URL string `json:"Url,omitempty"`
}

// Validate validates this v3 location option
func (m *V3LocationOption) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v3 location option based on context it is used
func (m *V3LocationOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V3LocationOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3LocationOption) UnmarshalBinary(b []byte) error {
	var res V3LocationOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
