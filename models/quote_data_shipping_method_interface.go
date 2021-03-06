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

// QuoteDataShippingMethodInterface Interface ShippingMethodInterface
//
// swagger:model quote-data-shipping-method-interface
type QuoteDataShippingMethodInterface struct {

	// Shipping amount in store currency.
	// Required: true
	Amount *float64 `json:"amount"`

	// The value of the availability flag for the current shipping method.
	// Required: true
	Available *bool `json:"available"`

	// Shipping amount in base currency.
	// Required: true
	BaseAmount *float64 `json:"base_amount"`

	// Shipping carrier code.
	// Required: true
	CarrierCode *string `json:"carrier_code"`

	// Shipping carrier title. Otherwise, null.
	CarrierTitle string `json:"carrier_title,omitempty"`

	// Shipping Error message.
	// Required: true
	ErrorMessage *string `json:"error_message"`

	// extension attributes
	ExtensionAttributes QuoteDataShippingMethodExtensionInterface `json:"extension_attributes,omitempty"`

	// Shipping method code.
	// Required: true
	MethodCode *string `json:"method_code"`

	// Shipping method title. Otherwise, null.
	MethodTitle string `json:"method_title,omitempty"`

	// Shipping price excl tax.
	// Required: true
	PriceExclTax *float64 `json:"price_excl_tax"`

	// Shipping price incl tax.
	// Required: true
	PriceInclTax *float64 `json:"price_incl_tax"`
}

// Validate validates this quote data shipping method interface
func (m *QuoteDataShippingMethodInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCarrierCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethodCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceExclTax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceInclTax(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuoteDataShippingMethodInterface) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *QuoteDataShippingMethodInterface) validateAvailable(formats strfmt.Registry) error {

	if err := validate.Required("available", "body", m.Available); err != nil {
		return err
	}

	return nil
}

func (m *QuoteDataShippingMethodInterface) validateBaseAmount(formats strfmt.Registry) error {

	if err := validate.Required("base_amount", "body", m.BaseAmount); err != nil {
		return err
	}

	return nil
}

func (m *QuoteDataShippingMethodInterface) validateCarrierCode(formats strfmt.Registry) error {

	if err := validate.Required("carrier_code", "body", m.CarrierCode); err != nil {
		return err
	}

	return nil
}

func (m *QuoteDataShippingMethodInterface) validateErrorMessage(formats strfmt.Registry) error {

	if err := validate.Required("error_message", "body", m.ErrorMessage); err != nil {
		return err
	}

	return nil
}

func (m *QuoteDataShippingMethodInterface) validateMethodCode(formats strfmt.Registry) error {

	if err := validate.Required("method_code", "body", m.MethodCode); err != nil {
		return err
	}

	return nil
}

func (m *QuoteDataShippingMethodInterface) validatePriceExclTax(formats strfmt.Registry) error {

	if err := validate.Required("price_excl_tax", "body", m.PriceExclTax); err != nil {
		return err
	}

	return nil
}

func (m *QuoteDataShippingMethodInterface) validatePriceInclTax(formats strfmt.Registry) error {

	if err := validate.Required("price_incl_tax", "body", m.PriceInclTax); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this quote data shipping method interface based on context it is used
func (m *QuoteDataShippingMethodInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuoteDataShippingMethodInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuoteDataShippingMethodInterface) UnmarshalBinary(b []byte) error {
	var res QuoteDataShippingMethodInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
