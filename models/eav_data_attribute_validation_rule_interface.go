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

// EavDataAttributeValidationRuleInterface Interface AttributeValidationRuleInterface
//
// swagger:model eav-data-attribute-validation-rule-interface
type EavDataAttributeValidationRuleInterface struct {

	// Object key
	// Required: true
	Key *string `json:"key"`

	// Object value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this eav data attribute validation rule interface
func (m *EavDataAttributeValidationRuleInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
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

func (m *EavDataAttributeValidationRuleInterface) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *EavDataAttributeValidationRuleInterface) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this eav data attribute validation rule interface based on context it is used
func (m *EavDataAttributeValidationRuleInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EavDataAttributeValidationRuleInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EavDataAttributeValidationRuleInterface) UnmarshalBinary(b []byte) error {
	var res EavDataAttributeValidationRuleInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
