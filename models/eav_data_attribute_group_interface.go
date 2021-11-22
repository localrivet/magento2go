// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EavDataAttributeGroupInterface Interface AttributeGroupInterface
//
// swagger:model eav-data-attribute-group-interface
type EavDataAttributeGroupInterface struct {

	// Id
	AttributeGroupID string `json:"attribute_group_id,omitempty"`

	// Name
	AttributeGroupName string `json:"attribute_group_name,omitempty"`

	// Attribute set id
	AttributeSetID int64 `json:"attribute_set_id,omitempty"`

	// extension attributes
	ExtensionAttributes *EavDataAttributeGroupExtensionInterface `json:"extension_attributes,omitempty"`
}

// Validate validates this eav data attribute group interface
func (m *EavDataAttributeGroupInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtensionAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EavDataAttributeGroupInterface) validateExtensionAttributes(formats strfmt.Registry) error {
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

// ContextValidate validate this eav data attribute group interface based on the context it is used
func (m *EavDataAttributeGroupInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExtensionAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EavDataAttributeGroupInterface) contextValidateExtensionAttributes(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EavDataAttributeGroupInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EavDataAttributeGroupInterface) UnmarshalBinary(b []byte) error {
	var res EavDataAttributeGroupInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
