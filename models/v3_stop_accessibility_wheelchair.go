// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3StopAccessibilityWheelchair v3 stop accessibility wheelchair
//
// swagger:model V3.StopAccessibilityWheelchair
type V3StopAccessibilityWheelchair struct {

	// accessible ramp
	AccessibleRamp bool `json:"accessible_ramp,omitempty"`

	// Indicates if there is at least one low ticket counter at the stop that complies with the Disability Standards for Accessible Public Transport under the Disability Discrimination Act (1992)
	LowTicketCounter bool `json:"low_ticket_counter,omitempty"`

	// Indicates if there is a space for mobility device to board on or off a transport mode
	Manouvering bool `json:"manouvering,omitempty"`

	// Indicates if there is at least one accessible parking spot at the stop that complies with the Disability Standards for Accessible Public Transport under the Disability Discrimination Act (1992)
	Parking bool `json:"parking,omitempty"`

	// Indicates if there is a raised platform to board a train
	RaisedPlatform bool `json:"raised_platform,omitempty"`

	// Indicates if there is shelter near the raised platform
	RaisedPlatformShelther bool `json:"raised_platform_shelther,omitempty"`

	// Indicates if there are ramps (&lt;1:14) at the stop/platform
	Ramp bool `json:"ramp,omitempty"`

	// Indicates if there is a path beyond the stop which is accessible
	SecondaryPath bool `json:"secondary_path,omitempty"`

	// Indicates if there are ramps (&gt;1:14) at the stop/platform
	SteepRamp bool `json:"steep_ramp,omitempty"`

	// Indicates if there is at least one accessible telephone at the stop/platform that complies with the Disability Standards for Accessible Public Transport under the Disability Discrimination Act (1992)
	Telephone bool `json:"telephone,omitempty"`

	// Indicates if there is at least one accessible toilet at the stop/platform that complies with the Disability Standards for Accessible Public Transport under the Disability Discrimination Act (1992)
	Toilet bool `json:"toilet,omitempty"`
}

// Validate validates this v3 stop accessibility wheelchair
func (m *V3StopAccessibilityWheelchair) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v3 stop accessibility wheelchair based on context it is used
func (m *V3StopAccessibilityWheelchair) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V3StopAccessibilityWheelchair) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3StopAccessibilityWheelchair) UnmarshalBinary(b []byte) error {
	var res V3StopAccessibilityWheelchair
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
