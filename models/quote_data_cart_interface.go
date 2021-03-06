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

// QuoteDataCartInterface Interface CartInterface
//
// swagger:model quote-data-cart-interface
type QuoteDataCartInterface struct {

	// billing address
	BillingAddress *QuoteDataAddressInterface `json:"billing_address,omitempty"`

	// Cart conversion date and time. Otherwise, null.
	ConvertedAt string `json:"converted_at,omitempty"`

	// Cart creation date and time. Otherwise, null.
	CreatedAt string `json:"created_at,omitempty"`

	// currency
	Currency *QuoteDataCurrencyInterface `json:"currency,omitempty"`

	// customer
	// Required: true
	Customer *CustomerDataCustomerInterface `json:"customer"`

	// For guest customers, false for logged in customers
	CustomerIsGuest bool `json:"customer_is_guest,omitempty"`

	// Notice text
	CustomerNote string `json:"customer_note,omitempty"`

	// Customer notification flag
	CustomerNoteNotify bool `json:"customer_note_notify,omitempty"`

	// Customer tax class ID.
	CustomerTaxClassID int64 `json:"customer_tax_class_id,omitempty"`

	// extension attributes
	ExtensionAttributes *QuoteDataCartExtensionInterface `json:"extension_attributes,omitempty"`

	// Cart/quote ID.
	// Required: true
	ID *int64 `json:"id"`

	// Active status flag value. Otherwise, null.
	IsActive bool `json:"is_active,omitempty"`

	// Virtual flag value. Otherwise, null.
	IsVirtual bool `json:"is_virtual,omitempty"`

	// Array of items. Otherwise, null.
	Items []*QuoteDataCartItemInterface `json:"items"`

	// Number of different items or products in the cart. Otherwise, null.
	ItemsCount int64 `json:"items_count,omitempty"`

	// Total quantity of all cart items. Otherwise, null.
	ItemsQty float64 `json:"items_qty,omitempty"`

	// Original order ID. Otherwise, null.
	OrigOrderID int64 `json:"orig_order_id,omitempty"`

	// Reserved order ID. Otherwise, null.
	ReservedOrderID string `json:"reserved_order_id,omitempty"`

	// Store identifier
	// Required: true
	StoreID *int64 `json:"store_id"`

	// Cart last update date and time. Otherwise, null.
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this quote data cart interface
func (m *QuoteDataCartInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensionAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoreID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuoteDataCartInterface) validateBillingAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingAddress) { // not required
		return nil
	}

	if m.BillingAddress != nil {
		if err := m.BillingAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billing_address")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteDataCartInterface) validateCurrency(formats strfmt.Registry) error {
	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	if m.Currency != nil {
		if err := m.Currency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("currency")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteDataCartInterface) validateCustomer(formats strfmt.Registry) error {

	if err := validate.Required("customer", "body", m.Customer); err != nil {
		return err
	}

	if m.Customer != nil {
		if err := m.Customer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customer")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteDataCartInterface) validateExtensionAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtensionAttributes) { // not required
		return nil
	}

	if m.ExtensionAttributes != nil {
		if err := m.ExtensionAttributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extension_attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extension_attributes")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteDataCartInterface) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *QuoteDataCartInterface) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuoteDataCartInterface) validateStoreID(formats strfmt.Registry) error {

	if err := validate.Required("store_id", "body", m.StoreID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this quote data cart interface based on the context it is used
func (m *QuoteDataCartInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBillingAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCurrency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtensionAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuoteDataCartInterface) contextValidateBillingAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.BillingAddress != nil {
		if err := m.BillingAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billing_address")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteDataCartInterface) contextValidateCurrency(ctx context.Context, formats strfmt.Registry) error {

	if m.Currency != nil {
		if err := m.Currency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currency")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("currency")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteDataCartInterface) contextValidateCustomer(ctx context.Context, formats strfmt.Registry) error {

	if m.Customer != nil {
		if err := m.Customer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customer")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteDataCartInterface) contextValidateExtensionAttributes(ctx context.Context, formats strfmt.Registry) error {

	if m.ExtensionAttributes != nil {
		if err := m.ExtensionAttributes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extension_attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extension_attributes")
			}
			return err
		}
	}

	return nil
}

func (m *QuoteDataCartInterface) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuoteDataCartInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuoteDataCartInterface) UnmarshalBinary(b []byte) error {
	var res QuoteDataCartInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
