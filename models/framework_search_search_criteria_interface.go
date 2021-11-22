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

// FrameworkSearchSearchCriteriaInterface Interface SearchCriteriaInterface
//
// swagger:model framework-search-search-criteria-interface
type FrameworkSearchSearchCriteriaInterface struct {

	// Current page.
	CurrentPage int64 `json:"current_page,omitempty"`

	// A list of filter groups.
	// Required: true
	FilterGroups []*FrameworkSearchFilterGroup `json:"filter_groups"`

	// Page size.
	PageSize int64 `json:"page_size,omitempty"`

	// request name
	// Required: true
	RequestName *string `json:"request_name"`

	// Sort order.
	SortOrders []*FrameworkSortOrder `json:"sort_orders"`
}

// Validate validates this framework search search criteria interface
func (m *FrameworkSearchSearchCriteriaInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilterGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSortOrders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FrameworkSearchSearchCriteriaInterface) validateFilterGroups(formats strfmt.Registry) error {

	if err := validate.Required("filter_groups", "body", m.FilterGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.FilterGroups); i++ {
		if swag.IsZero(m.FilterGroups[i]) { // not required
			continue
		}

		if m.FilterGroups[i] != nil {
			if err := m.FilterGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filter_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filter_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FrameworkSearchSearchCriteriaInterface) validateRequestName(formats strfmt.Registry) error {

	if err := validate.Required("request_name", "body", m.RequestName); err != nil {
		return err
	}

	return nil
}

func (m *FrameworkSearchSearchCriteriaInterface) validateSortOrders(formats strfmt.Registry) error {
	if swag.IsZero(m.SortOrders) { // not required
		return nil
	}

	for i := 0; i < len(m.SortOrders); i++ {
		if swag.IsZero(m.SortOrders[i]) { // not required
			continue
		}

		if m.SortOrders[i] != nil {
			if err := m.SortOrders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sort_orders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sort_orders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this framework search search criteria interface based on the context it is used
func (m *FrameworkSearchSearchCriteriaInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilterGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSortOrders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FrameworkSearchSearchCriteriaInterface) contextValidateFilterGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FilterGroups); i++ {

		if m.FilterGroups[i] != nil {
			if err := m.FilterGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filter_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filter_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FrameworkSearchSearchCriteriaInterface) contextValidateSortOrders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SortOrders); i++ {

		if m.SortOrders[i] != nil {
			if err := m.SortOrders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sort_orders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sort_orders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FrameworkSearchSearchCriteriaInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrameworkSearchSearchCriteriaInterface) UnmarshalBinary(b []byte) error {
	var res FrameworkSearchSearchCriteriaInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
