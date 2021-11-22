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

// CmsDataPageInterface CMS page interface.
//
// swagger:model cms-data-page-interface
type CmsDataPageInterface struct {

	// Active
	Active bool `json:"active,omitempty"`

	// Content
	Content string `json:"content,omitempty"`

	// Content heading
	ContentHeading string `json:"content_heading,omitempty"`

	// Creation time
	CreationTime string `json:"creation_time,omitempty"`

	// Custom layout update xml
	CustomLayoutUpdateXML string `json:"custom_layout_update_xml,omitempty"`

	// Custom root template
	CustomRootTemplate string `json:"custom_root_template,omitempty"`

	// Custom theme
	CustomTheme string `json:"custom_theme,omitempty"`

	// Custom theme from
	CustomThemeFrom string `json:"custom_theme_from,omitempty"`

	// Custom theme to
	CustomThemeTo string `json:"custom_theme_to,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// Identifier
	// Required: true
	Identifier *string `json:"identifier"`

	// Layout update xml
	LayoutUpdateXML string `json:"layout_update_xml,omitempty"`

	// Meta description
	MetaDescription string `json:"meta_description,omitempty"`

	// Meta keywords
	MetaKeywords string `json:"meta_keywords,omitempty"`

	// Meta title
	MetaTitle string `json:"meta_title,omitempty"`

	// Page layout
	PageLayout string `json:"page_layout,omitempty"`

	// Sort order
	SortOrder string `json:"sort_order,omitempty"`

	// Title
	Title string `json:"title,omitempty"`

	// Update time
	UpdateTime string `json:"update_time,omitempty"`
}

// Validate validates this cms data page interface
func (m *CmsDataPageInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CmsDataPageInterface) validateIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("identifier", "body", m.Identifier); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cms data page interface based on context it is used
func (m *CmsDataPageInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CmsDataPageInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CmsDataPageInterface) UnmarshalBinary(b []byte) error {
	var res CmsDataPageInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
