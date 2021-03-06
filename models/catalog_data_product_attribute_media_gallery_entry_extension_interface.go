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

// CatalogDataProductAttributeMediaGalleryEntryExtensionInterface ExtensionInterface class for @see \Magento\Catalog\Api\Data\ProductAttributeMediaGalleryEntryInterface
//
// swagger:model catalog-data-product-attribute-media-gallery-entry-extension-interface
type CatalogDataProductAttributeMediaGalleryEntryExtensionInterface struct {

	// video content
	VideoContent *FrameworkDataVideoContentInterface `json:"video_content,omitempty"`
}

// Validate validates this catalog data product attribute media gallery entry extension interface
func (m *CatalogDataProductAttributeMediaGalleryEntryExtensionInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVideoContent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogDataProductAttributeMediaGalleryEntryExtensionInterface) validateVideoContent(formats strfmt.Registry) error {
	if swag.IsZero(m.VideoContent) { // not required
		return nil
	}

	if m.VideoContent != nil {
		if err := m.VideoContent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("video_content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("video_content")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this catalog data product attribute media gallery entry extension interface based on the context it is used
func (m *CatalogDataProductAttributeMediaGalleryEntryExtensionInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVideoContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogDataProductAttributeMediaGalleryEntryExtensionInterface) contextValidateVideoContent(ctx context.Context, formats strfmt.Registry) error {

	if m.VideoContent != nil {
		if err := m.VideoContent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("video_content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("video_content")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogDataProductAttributeMediaGalleryEntryExtensionInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogDataProductAttributeMediaGalleryEntryExtensionInterface) UnmarshalBinary(b []byte) error {
	var res CatalogDataProductAttributeMediaGalleryEntryExtensionInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
