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

// CatalogDataProductRenderPriceInfoExtensionInterface ExtensionInterface class for @see \Magento\Catalog\Api\Data\ProductRender\PriceInfoInterface
//
// swagger:model catalog-data-product-render-price-info-extension-interface
type CatalogDataProductRenderPriceInfoExtensionInterface struct {

	// msrp
	Msrp *MsrpDataProductRenderMsrpPriceInfoInterface `json:"msrp,omitempty"`

	// tax adjustments
	TaxAdjustments *CatalogDataProductRenderPriceInfoInterface `json:"tax_adjustments,omitempty"`

	// weee adjustment
	WeeeAdjustment string `json:"weee_adjustment,omitempty"`

	// weee attributes
	WeeeAttributes []*WeeeDataProductRenderWeeeAdjustmentAttributeInterface `json:"weee_attributes"`
}

// Validate validates this catalog data product render price info extension interface
func (m *CatalogDataProductRenderPriceInfoExtensionInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMsrp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxAdjustments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeeeAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogDataProductRenderPriceInfoExtensionInterface) validateMsrp(formats strfmt.Registry) error {
	if swag.IsZero(m.Msrp) { // not required
		return nil
	}

	if m.Msrp != nil {
		if err := m.Msrp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("msrp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("msrp")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogDataProductRenderPriceInfoExtensionInterface) validateTaxAdjustments(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxAdjustments) { // not required
		return nil
	}

	if m.TaxAdjustments != nil {
		if err := m.TaxAdjustments.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tax_adjustments")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tax_adjustments")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogDataProductRenderPriceInfoExtensionInterface) validateWeeeAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.WeeeAttributes) { // not required
		return nil
	}

	for i := 0; i < len(m.WeeeAttributes); i++ {
		if swag.IsZero(m.WeeeAttributes[i]) { // not required
			continue
		}

		if m.WeeeAttributes[i] != nil {
			if err := m.WeeeAttributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("weee_attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("weee_attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this catalog data product render price info extension interface based on the context it is used
func (m *CatalogDataProductRenderPriceInfoExtensionInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMsrp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxAdjustments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWeeeAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogDataProductRenderPriceInfoExtensionInterface) contextValidateMsrp(ctx context.Context, formats strfmt.Registry) error {

	if m.Msrp != nil {
		if err := m.Msrp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("msrp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("msrp")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogDataProductRenderPriceInfoExtensionInterface) contextValidateTaxAdjustments(ctx context.Context, formats strfmt.Registry) error {

	if m.TaxAdjustments != nil {
		if err := m.TaxAdjustments.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tax_adjustments")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tax_adjustments")
			}
			return err
		}
	}

	return nil
}

func (m *CatalogDataProductRenderPriceInfoExtensionInterface) contextValidateWeeeAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WeeeAttributes); i++ {

		if m.WeeeAttributes[i] != nil {
			if err := m.WeeeAttributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("weee_attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("weee_attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogDataProductRenderPriceInfoExtensionInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogDataProductRenderPriceInfoExtensionInterface) UnmarshalBinary(b []byte) error {
	var res CatalogDataProductRenderPriceInfoExtensionInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
