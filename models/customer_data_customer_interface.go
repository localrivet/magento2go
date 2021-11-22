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

// CustomerDataCustomerInterface Customer entity interface for API handling.
//
// swagger:model customer-data-customer-interface
type CustomerDataCustomerInterface struct {

	// Customer addresses.
	Addresses []*CustomerDataAddressInterface `json:"addresses"`

	// Confirmation
	Confirmation string `json:"confirmation,omitempty"`

	// Created at time
	CreatedAt string `json:"created_at,omitempty"`

	// Created in area
	CreatedIn string `json:"created_in,omitempty"`

	// Custom attributes values.
	CustomAttributes []*FrameworkAttributeInterface `json:"custom_attributes"`

	// Default billing address id
	DefaultBilling string `json:"default_billing,omitempty"`

	// Default shipping address id
	DefaultShipping string `json:"default_shipping,omitempty"`

	// Disable auto group change flag.
	DisableAutoGroupChange int64 `json:"disable_auto_group_change,omitempty"`

	// In keeping with current security and privacy best practices, be sure you are aware of any potential legal and security risks associated with the storage of customers’ full date of birth (month, day, year) along with other personal identifiers (e.g., full name) before collecting or processing such data.
	Dob string `json:"dob,omitempty"`

	// Email address
	// Required: true
	Email *string `json:"email"`

	// extension attributes
	ExtensionAttributes *CustomerDataCustomerExtensionInterface `json:"extension_attributes,omitempty"`

	// First name
	// Required: true
	Firstname *string `json:"firstname"`

	// Gender
	Gender int64 `json:"gender,omitempty"`

	// Group id
	GroupID int64 `json:"group_id,omitempty"`

	// Customer id
	ID int64 `json:"id,omitempty"`

	// Last name
	// Required: true
	Lastname *string `json:"lastname"`

	// Middle name
	Middlename string `json:"middlename,omitempty"`

	// Prefix
	Prefix string `json:"prefix,omitempty"`

	// Store id
	StoreID int64 `json:"store_id,omitempty"`

	// Suffix
	Suffix string `json:"suffix,omitempty"`

	// Tax Vat
	Taxvat string `json:"taxvat,omitempty"`

	// Updated at time
	UpdatedAt string `json:"updated_at,omitempty"`

	// Website id
	WebsiteID int64 `json:"website_id,omitempty"`
}

// Validate validates this customer data customer interface
func (m *CustomerDataCustomerInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensionAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerDataCustomerInterface) validateAddresses(formats strfmt.Registry) error {
	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.Addresses); i++ {
		if swag.IsZero(m.Addresses[i]) { // not required
			continue
		}

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomerDataCustomerInterface) validateCustomAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomAttributes) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomAttributes); i++ {
		if swag.IsZero(m.CustomAttributes[i]) { // not required
			continue
		}

		if m.CustomAttributes[i] != nil {
			if err := m.CustomAttributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("custom_attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomerDataCustomerInterface) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *CustomerDataCustomerInterface) validateExtensionAttributes(formats strfmt.Registry) error {
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

func (m *CustomerDataCustomerInterface) validateFirstname(formats strfmt.Registry) error {

	if err := validate.Required("firstname", "body", m.Firstname); err != nil {
		return err
	}

	return nil
}

func (m *CustomerDataCustomerInterface) validateLastname(formats strfmt.Registry) error {

	if err := validate.Required("lastname", "body", m.Lastname); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this customer data customer interface based on the context it is used
func (m *CustomerDataCustomerInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtensionAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerDataCustomerInterface) contextValidateAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Addresses); i++ {

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("addresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomerDataCustomerInterface) contextValidateCustomAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomAttributes); i++ {

		if m.CustomAttributes[i] != nil {
			if err := m.CustomAttributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("custom_attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomerDataCustomerInterface) contextValidateExtensionAttributes(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CustomerDataCustomerInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerDataCustomerInterface) UnmarshalBinary(b []byte) error {
	var res CustomerDataCustomerInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
