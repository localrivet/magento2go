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

// TaxDataGrandTotalRatesInterface Interface GrandTotalRatesInterface
//
// swagger:model tax-data-grand-total-rates-interface
type TaxDataGrandTotalRatesInterface struct {

	// Tax percentage value
	// Required: true
	Percent *string `json:"percent"`

	// Rate title
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this tax data grand total rates interface
func (m *TaxDataGrandTotalRatesInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePercent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxDataGrandTotalRatesInterface) validatePercent(formats strfmt.Registry) error {

	if err := validate.Required("percent", "body", m.Percent); err != nil {
		return err
	}

	return nil
}

func (m *TaxDataGrandTotalRatesInterface) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this tax data grand total rates interface based on context it is used
func (m *TaxDataGrandTotalRatesInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaxDataGrandTotalRatesInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaxDataGrandTotalRatesInterface) UnmarshalBinary(b []byte) error {
	var res TaxDataGrandTotalRatesInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}