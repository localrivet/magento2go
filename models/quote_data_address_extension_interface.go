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
)

// QuoteDataAddressExtensionInterface ExtensionInterface class for @see \Magento\Quote\Api\Data\AddressInterface
//
// swagger:model quote-data-address-extension-interface
type QuoteDataAddressExtensionInterface struct {

	// discounts
	Discounts []*SalesRuleDataRuleDiscountInterface `json:"discounts"`

	// pickup location code
	PickupLocationCode string `json:"pickup_location_code,omitempty"`
}

// Validate validates this quote data address extension interface
func (m *QuoteDataAddressExtensionInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscounts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuoteDataAddressExtensionInterface) validateDiscounts(formats strfmt.Registry) error {
	if swag.IsZero(m.Discounts) { // not required
		return nil
	}

	for i := 0; i < len(m.Discounts); i++ {
		if swag.IsZero(m.Discounts[i]) { // not required
			continue
		}

		if m.Discounts[i] != nil {
			if err := m.Discounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("discounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("discounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this quote data address extension interface based on the context it is used
func (m *QuoteDataAddressExtensionInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDiscounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuoteDataAddressExtensionInterface) contextValidateDiscounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Discounts); i++ {

		if m.Discounts[i] != nil {
			if err := m.Discounts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("discounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("discounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuoteDataAddressExtensionInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuoteDataAddressExtensionInterface) UnmarshalBinary(b []byte) error {
	var res QuoteDataAddressExtensionInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
