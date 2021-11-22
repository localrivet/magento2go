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

// VertexAddressValidationAPIDataCleansedAddressInterface What we want to output to the API contains more data than what we want to bring in.  Chiefly, we want to ensure we're sending Magento data out (instead of Vertex data). That means sending Region IDs, and 2-character ISO country codes.
//
// swagger:model vertex-address-validation-api-data-cleansed-address-interface
type VertexAddressValidationAPIDataCleansedAddressInterface struct {

	// city
	City string `json:"city,omitempty"`

	// country code
	CountryCode string `json:"country_code,omitempty"`

	// country name
	CountryName string `json:"country_name,omitempty"`

	// postal code
	PostalCode string `json:"postal_code,omitempty"`

	// region id
	RegionID int64 `json:"region_id,omitempty"`

	// region name
	RegionName string `json:"region_name,omitempty"`

	// street address
	// Required: true
	StreetAddress []string `json:"street_address"`

	// The regional sub division, such as a county or parish
	SubDivision string `json:"sub_division,omitempty"`
}

// Validate validates this vertex address validation api data cleansed address interface
func (m *VertexAddressValidationAPIDataCleansedAddressInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStreetAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VertexAddressValidationAPIDataCleansedAddressInterface) validateStreetAddress(formats strfmt.Registry) error {

	if err := validate.Required("street_address", "body", m.StreetAddress); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vertex address validation api data cleansed address interface based on context it is used
func (m *VertexAddressValidationAPIDataCleansedAddressInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VertexAddressValidationAPIDataCleansedAddressInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VertexAddressValidationAPIDataCleansedAddressInterface) UnmarshalBinary(b []byte) error {
	var res VertexAddressValidationAPIDataCleansedAddressInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}