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

// CatalogDataProductRenderExtensionInterface ExtensionInterface class for @see \Magento\Catalog\Api\Data\ProductRenderInterface
//
// swagger:model catalog-data-product-render-extension-interface
type CatalogDataProductRenderExtensionInterface struct {

	// ddg brand
	DdgBrand string `json:"ddg_brand,omitempty"`

	// ddg categories
	DdgCategories []string `json:"ddg_categories"`

	// ddg description
	DdgDescription string `json:"ddg_description,omitempty"`

	// ddg image
	DdgImage string `json:"ddg_image,omitempty"`

	// ddg sku
	DdgSku string `json:"ddg_sku,omitempty"`

	// review html
	ReviewHTML string `json:"review_html,omitempty"`

	// wishlist button
	WishlistButton *CatalogDataProductRenderButtonInterface `json:"wishlist_button,omitempty"`
}

// Validate validates this catalog data product render extension interface
func (m *CatalogDataProductRenderExtensionInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWishlistButton(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogDataProductRenderExtensionInterface) validateWishlistButton(formats strfmt.Registry) error {
	if swag.IsZero(m.WishlistButton) { // not required
		return nil
	}

	if m.WishlistButton != nil {
		if err := m.WishlistButton.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wishlist_button")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wishlist_button")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this catalog data product render extension interface based on the context it is used
func (m *CatalogDataProductRenderExtensionInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWishlistButton(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogDataProductRenderExtensionInterface) contextValidateWishlistButton(ctx context.Context, formats strfmt.Registry) error {

	if m.WishlistButton != nil {
		if err := m.WishlistButton.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wishlist_button")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wishlist_button")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogDataProductRenderExtensionInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogDataProductRenderExtensionInterface) UnmarshalBinary(b []byte) error {
	var res CatalogDataProductRenderExtensionInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}