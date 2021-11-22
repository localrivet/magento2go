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

// InventoryInStorePickupAPIDataSearchRequestInterface Endpoint used to search Pickup Locations by different parameters: - by attribute filters fields @see \Magento\InventoryInStorePickupApi\Api\Data\SearchRequest\FiltersInterface - by distance to the address @see \Magento\InventoryInStorePickupApi\Api\Data\SearchRequest\AreaInterface Also, endpoint supports paging and sort orders.
//
// swagger:model inventory-in-store-pickup-api-data-search-request-interface
type InventoryInStorePickupAPIDataSearchRequestInterface struct {

	// area
	Area *InventoryInStorePickupAPIDataSearchRequestAreaInterface `json:"area,omitempty"`

	// Current page.
	// Required: true
	CurrentPage *int64 `json:"current_page"`

	// extension attributes
	ExtensionAttributes *InventoryInStorePickupAPIDataSearchRequestExtensionInterface `json:"extension_attributes,omitempty"`

	// filters
	Filters *InventoryInStorePickupAPIDataSearchRequestFiltersInterface `json:"filters,omitempty"`

	// Page size.
	PageSize int64 `json:"page_size,omitempty"`

	// Sales Channel code.
	// Required: true
	ScopeCode *string `json:"scope_code"`

	// Sales Channel Type.
	// Required: true
	ScopeType *string `json:"scope_type"`

	// Sort Order.
	Sort []*FrameworkSortOrder `json:"sort"`
}

// Validate validates this inventory in store pickup api data search request interface
func (m *InventoryInStorePickupAPIDataSearchRequestInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArea(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentPage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensionAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScopeCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScopeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryInStorePickupAPIDataSearchRequestInterface) validateArea(formats strfmt.Registry) error {
	if swag.IsZero(m.Area) { // not required
		return nil
	}

	if m.Area != nil {
		if err := m.Area.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("area")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("area")
			}
			return err
		}
	}

	return nil
}

func (m *InventoryInStorePickupAPIDataSearchRequestInterface) validateCurrentPage(formats strfmt.Registry) error {

	if err := validate.Required("current_page", "body", m.CurrentPage); err != nil {
		return err
	}

	return nil
}

func (m *InventoryInStorePickupAPIDataSearchRequestInterface) validateExtensionAttributes(formats strfmt.Registry) error {
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

func (m *InventoryInStorePickupAPIDataSearchRequestInterface) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	if m.Filters != nil {
		if err := m.Filters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filters")
			}
			return err
		}
	}

	return nil
}

func (m *InventoryInStorePickupAPIDataSearchRequestInterface) validateScopeCode(formats strfmt.Registry) error {

	if err := validate.Required("scope_code", "body", m.ScopeCode); err != nil {
		return err
	}

	return nil
}

func (m *InventoryInStorePickupAPIDataSearchRequestInterface) validateScopeType(formats strfmt.Registry) error {

	if err := validate.Required("scope_type", "body", m.ScopeType); err != nil {
		return err
	}

	return nil
}

func (m *InventoryInStorePickupAPIDataSearchRequestInterface) validateSort(formats strfmt.Registry) error {
	if swag.IsZero(m.Sort) { // not required
		return nil
	}

	for i := 0; i < len(m.Sort); i++ {
		if swag.IsZero(m.Sort[i]) { // not required
			continue
		}

		if m.Sort[i] != nil {
			if err := m.Sort[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sort" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this inventory in store pickup api data search request interface based on the context it is used
func (m *InventoryInStorePickupAPIDataSearchRequestInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArea(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtensionAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryInStorePickupAPIDataSearchRequestInterface) contextValidateArea(ctx context.Context, formats strfmt.Registry) error {

	if m.Area != nil {
		if err := m.Area.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("area")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("area")
			}
			return err
		}
	}

	return nil
}

func (m *InventoryInStorePickupAPIDataSearchRequestInterface) contextValidateExtensionAttributes(ctx context.Context, formats strfmt.Registry) error {

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

func (m *InventoryInStorePickupAPIDataSearchRequestInterface) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	if m.Filters != nil {
		if err := m.Filters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filters")
			}
			return err
		}
	}

	return nil
}

func (m *InventoryInStorePickupAPIDataSearchRequestInterface) contextValidateSort(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sort); i++ {

		if m.Sort[i] != nil {
			if err := m.Sort[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sort" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InventoryInStorePickupAPIDataSearchRequestInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryInStorePickupAPIDataSearchRequestInterface) UnmarshalBinary(b []byte) error {
	var res InventoryInStorePickupAPIDataSearchRequestInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
