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

// FrameworkAttributeInterface Interface for custom attribute value.
//
// swagger:model framework-attribute-interface
type FrameworkAttributeInterface struct {

	// Attribute code
	// Required: true
	AttributeCode *string `json:"attribute_code"`

	// Attribute value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this framework attribute interface
func (m *FrameworkAttributeInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FrameworkAttributeInterface) validateAttributeCode(formats strfmt.Registry) error {

	if err := validate.Required("attribute_code", "body", m.AttributeCode); err != nil {
		return err
	}

	return nil
}

func (m *FrameworkAttributeInterface) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this framework attribute interface based on context it is used
func (m *FrameworkAttributeInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FrameworkAttributeInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrameworkAttributeInterface) UnmarshalBinary(b []byte) error {
	var res FrameworkAttributeInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
