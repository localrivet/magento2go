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

// VertexAddressValidationAPIDataAddressInterface This is the interface for submission to the API.  For ease the API only takes interfaces the same as the SDK, but the SDK model cannot implement the interface.  Since the SDK model technically isn't API, we re-implement a simple DTO to handle the scenario.
//
// swagger:model vertex-address-validation-api-data-address-interface
type VertexAddressValidationAPIDataAddressInterface struct {

	// The proper name of the city
	City string `json:"city,omitempty"`

	// ISO 3166-1 Alpha-3 country code
	Country string `json:"country,omitempty"`

	// The proper name or the postal abbreviation of the state, province, or territory
	MainDivision string `json:"main_division,omitempty"`

	// The Postal Code
	PostalCode string `json:"postal_code,omitempty"`

	// The street address
	// Required: true
	StreetAddress []string `json:"street_address"`

	// The name of the county
	SubDivision string `json:"sub_division,omitempty"`
}

// Validate validates this vertex address validation api data address interface
func (m *VertexAddressValidationAPIDataAddressInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStreetAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VertexAddressValidationAPIDataAddressInterface) validateStreetAddress(formats strfmt.Registry) error {

	if err := validate.Required("street_address", "body", m.StreetAddress); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vertex address validation api data address interface based on context it is used
func (m *VertexAddressValidationAPIDataAddressInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VertexAddressValidationAPIDataAddressInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VertexAddressValidationAPIDataAddressInterface) UnmarshalBinary(b []byte) error {
	var res VertexAddressValidationAPIDataAddressInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
