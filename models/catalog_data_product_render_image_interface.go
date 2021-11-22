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

// CatalogDataProductRenderImageInterface Product Render image interface. Represents physical characteristics of image, that can be used in product listing or product view
//
// swagger:model catalog-data-product-render-image-interface
type CatalogDataProductRenderImageInterface struct {

	// Image code
	// Required: true
	Code *string `json:"code"`

	// extension attributes
	ExtensionAttributes CatalogDataProductRenderImageExtensionInterface `json:"extension_attributes,omitempty"`

	// Image height
	// Required: true
	Height *float64 `json:"height"`

	// Image label
	// Required: true
	Label *string `json:"label"`

	// Resize height
	// Required: true
	ResizedHeight *float64 `json:"resized_height"`

	// Resize width
	// Required: true
	ResizedWidth *float64 `json:"resized_width"`

	// Image url
	// Required: true
	URL *string `json:"url"`

	// Image width in px
	// Required: true
	Width *float64 `json:"width"`
}

// Validate validates this catalog data product render image interface
func (m *CatalogDataProductRenderImageInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResizedHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResizedWidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWidth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogDataProductRenderImageInterface) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductRenderImageInterface) validateHeight(formats strfmt.Registry) error {

	if err := validate.Required("height", "body", m.Height); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductRenderImageInterface) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductRenderImageInterface) validateResizedHeight(formats strfmt.Registry) error {

	if err := validate.Required("resized_height", "body", m.ResizedHeight); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductRenderImageInterface) validateResizedWidth(formats strfmt.Registry) error {

	if err := validate.Required("resized_width", "body", m.ResizedWidth); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductRenderImageInterface) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

func (m *CatalogDataProductRenderImageInterface) validateWidth(formats strfmt.Registry) error {

	if err := validate.Required("width", "body", m.Width); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this catalog data product render image interface based on context it is used
func (m *CatalogDataProductRenderImageInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogDataProductRenderImageInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogDataProductRenderImageInterface) UnmarshalBinary(b []byte) error {
	var res CatalogDataProductRenderImageInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}