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

// EavDataAttributeOptionInterface Created from:
//
// swagger:model eav-data-attribute-option-interface
type EavDataAttributeOptionInterface struct {

	// Default
	IsDefault bool `json:"is_default,omitempty"`

	// Option label
	// Required: true
	Label *string `json:"label"`

	// Option order
	SortOrder int64 `json:"sort_order,omitempty"`

	// Option label for store scopes
	StoreLabels []*EavDataAttributeOptionLabelInterface `json:"store_labels"`

	// Option value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this eav data attribute option interface
func (m *EavDataAttributeOptionInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoreLabels(formats); err != nil {
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

func (m *EavDataAttributeOptionInterface) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *EavDataAttributeOptionInterface) validateStoreLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.StoreLabels) { // not required
		return nil
	}

	for i := 0; i < len(m.StoreLabels); i++ {
		if swag.IsZero(m.StoreLabels[i]) { // not required
			continue
		}

		if m.StoreLabels[i] != nil {
			if err := m.StoreLabels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("store_labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("store_labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EavDataAttributeOptionInterface) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this eav data attribute option interface based on the context it is used
func (m *EavDataAttributeOptionInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStoreLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EavDataAttributeOptionInterface) contextValidateStoreLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StoreLabels); i++ {

		if m.StoreLabels[i] != nil {
			if err := m.StoreLabels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("store_labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("store_labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EavDataAttributeOptionInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EavDataAttributeOptionInterface) UnmarshalBinary(b []byte) error {
	var res EavDataAttributeOptionInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
