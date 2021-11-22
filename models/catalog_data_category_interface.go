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

// CatalogDataCategoryInterface Category data interface.
//
// swagger:model catalog-data-category-interface
type CatalogDataCategoryInterface struct {

	// Available sort by for category.
	AvailableSortBy []string `json:"available_sort_by"`

	// Children ids comma separated.
	Children string `json:"children,omitempty"`

	// Category creation date and time.
	CreatedAt string `json:"created_at,omitempty"`

	// Custom attributes values.
	CustomAttributes []*FrameworkAttributeInterface `json:"custom_attributes"`

	// extension attributes
	ExtensionAttributes CatalogDataCategoryExtensionInterface `json:"extension_attributes,omitempty"`

	// Category id.
	ID int64 `json:"id,omitempty"`

	// Category is included in menu.
	IncludeInMenu bool `json:"include_in_menu,omitempty"`

	// Whether category is active
	IsActive bool `json:"is_active,omitempty"`

	// Category level
	Level int64 `json:"level,omitempty"`

	// Category name
	Name string `json:"name,omitempty"`

	// Parent category ID
	ParentID int64 `json:"parent_id,omitempty"`

	// Category full path.
	Path string `json:"path,omitempty"`

	// Category position
	Position int64 `json:"position,omitempty"`

	// Category last update date and time.
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this catalog data category interface
func (m *CatalogDataCategoryInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogDataCategoryInterface) validateCustomAttributes(formats strfmt.Registry) error {
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

// ContextValidate validate this catalog data category interface based on the context it is used
func (m *CatalogDataCategoryInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogDataCategoryInterface) contextValidateCustomAttributes(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CatalogDataCategoryInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogDataCategoryInterface) UnmarshalBinary(b []byte) error {
	var res CatalogDataCategoryInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
