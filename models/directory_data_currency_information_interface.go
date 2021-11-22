// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DirectoryDataCurrencyInformationInterface Currency Information interface.
//
// swagger:model directory-data-currency-information-interface
type DirectoryDataCurrencyInformationInterface struct {

	// The list of allowed currency codes for the store.
	// Required: true
	AvailableCurrencyCodes []string `json:"available_currency_codes"`

	// The base currency code for the store.
	// Required: true
	BaseCurrencyCode *string `json:"base_currency_code"`

	// The currency symbol of the base currency for the store.
	// Required: true
	BaseCurrencySymbol *string `json:"base_currency_symbol"`

	// The default display currency code for the store.
	// Required: true
	DefaultDisplayCurrencyCode *string `json:"default_display_currency_code"`

	// The currency symbol of the default display currency for the store.
	// Required: true
	DefaultDisplayCurrencySymbol *string `json:"default_display_currency_symbol"`

	// The list of exchange rate information for the store.
	// Required: true
	ExchangeRates []*DirectoryDataExchangeRateInterface `json:"exchange_rates"`

	// extension attributes
	ExtensionAttributes DirectoryDataCurrencyInformationExtensionInterface `json:"extension_attributes,omitempty"`
}

// Validate validates this directory data currency information interface
func (m *DirectoryDataCurrencyInformationInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableCurrencyCodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseCurrencyCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseCurrencySymbol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultDisplayCurrencyCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultDisplayCurrencySymbol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExchangeRates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectoryDataCurrencyInformationInterface) validateAvailableCurrencyCodes(formats strfmt.Registry) error {

	if err := validate.Required("available_currency_codes", "body", m.AvailableCurrencyCodes); err != nil {
		return err
	}

	return nil
}

func (m *DirectoryDataCurrencyInformationInterface) validateBaseCurrencyCode(formats strfmt.Registry) error {

	if err := validate.Required("base_currency_code", "body", m.BaseCurrencyCode); err != nil {
		return err
	}

	return nil
}

func (m *DirectoryDataCurrencyInformationInterface) validateBaseCurrencySymbol(formats strfmt.Registry) error {

	if err := validate.Required("base_currency_symbol", "body", m.BaseCurrencySymbol); err != nil {
		return err
	}

	return nil
}

func (m *DirectoryDataCurrencyInformationInterface) validateDefaultDisplayCurrencyCode(formats strfmt.Registry) error {

	if err := validate.Required("default_display_currency_code", "body", m.DefaultDisplayCurrencyCode); err != nil {
		return err
	}

	return nil
}

func (m *DirectoryDataCurrencyInformationInterface) validateDefaultDisplayCurrencySymbol(formats strfmt.Registry) error {

	if err := validate.Required("default_display_currency_symbol", "body", m.DefaultDisplayCurrencySymbol); err != nil {
		return err
	}

	return nil
}

func (m *DirectoryDataCurrencyInformationInterface) validateExchangeRates(formats strfmt.Registry) error {

	if err := validate.Required("exchange_rates", "body", m.ExchangeRates); err != nil {
		return err
	}

	for i := 0; i < len(m.ExchangeRates); i++ {
		if swag.IsZero(m.ExchangeRates[i]) { // not required
			continue
		}

		if m.ExchangeRates[i] != nil {
			if err := m.ExchangeRates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exchange_rates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("exchange_rates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this directory data currency information interface based on the context it is used
func (m *DirectoryDataCurrencyInformationInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExchangeRates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectoryDataCurrencyInformationInterface) contextValidateExchangeRates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExchangeRates); i++ {

		if m.ExchangeRates[i] != nil {
			if err := m.ExchangeRates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exchange_rates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("exchange_rates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectoryDataCurrencyInformationInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectoryDataCurrencyInformationInterface) UnmarshalBinary(b []byte) error {
	var res DirectoryDataCurrencyInformationInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
