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

// CatalogDataProductSearchResultsInterface catalog data product search results interface
//
// swagger:model catalog-data-product-search-results-interface
type CatalogDataProductSearchResultsInterface struct {

	// Attributes list.
	// Required: true
	Items []*CatalogDataProductInterface `json:"items"`

	// search criteria
	// Required: true
	SearchCriteria *FrameworkSearchCriteriaInterface `json:"search_criteria"`

	// Total count.
	// Required: true
	TotalCount *int64 `json:"total_count"`
}

// Validate validates this catalog data product search results interface
func (m *CatalogDataProductSearchResultsInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchCriteria(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogDataProductSearchResultsInterface) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CatalogDataProductSearchResultsInterface) validateSearchCriteria(formats strfmt.Registry) error {

	if err := validate.Required("search_criteria", "body", m.SearchCriteria); err != nil {
		return err
	}

	if m.SearchCriteria != nil {
		if err := m.SearchCriteria.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("search_criteria")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("search_criteria")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogDataProductSearchResultsInterface) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("total_count", "body", m.TotalCount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this catalog data product search results interface based on the context it is used
func (m *CatalogDataProductSearchResultsInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSearchCriteria(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogDataProductSearchResultsInterface) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CatalogDataProductSearchResultsInterface) contextValidateSearchCriteria(ctx context.Context, formats strfmt.Registry) error {

	if m.SearchCriteria != nil {
		if err := m.SearchCriteria.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("search_criteria")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("search_criteria")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogDataProductSearchResultsInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogDataProductSearchResultsInterface) UnmarshalBinary(b []byte) error {
	var res CatalogDataProductSearchResultsInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
