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

// WeeeDataProductRenderWeeeAdjustmentAttributeInterface List of all weee attributes, their amounts, etc.., that product has
//
// swagger:model weee-data-product-render-weee-adjustment-attribute-interface
type WeeeDataProductRenderWeeeAdjustmentAttributeInterface struct {

	// Weee attribute amount
	// Required: true
	Amount *string `json:"amount"`

	// Product amount exclude tax
	// Required: true
	AmountExclTax *string `json:"amount_excl_tax"`

	// Weee attribute code
	// Required: true
	AttributeCode *string `json:"attribute_code"`

	// extension attributes
	// Required: true
	ExtensionAttributes WeeeDataProductRenderWeeeAdjustmentAttributeExtensionInterface `json:"extension_attributes"`

	// Tax which is calculated to fixed product tax attribute
	// Required: true
	TaxAmount *string `json:"tax_amount"`

	// Tax amount of weee attribute
	// Required: true
	TaxAmountInclTax *string `json:"tax_amount_incl_tax"`
}

// Validate validates this weee data product render weee adjustment attribute interface
func (m *WeeeDataProductRenderWeeeAdjustmentAttributeInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmountExclTax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributeCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensionAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxAmountInclTax(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WeeeDataProductRenderWeeeAdjustmentAttributeInterface) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *WeeeDataProductRenderWeeeAdjustmentAttributeInterface) validateAmountExclTax(formats strfmt.Registry) error {

	if err := validate.Required("amount_excl_tax", "body", m.AmountExclTax); err != nil {
		return err
	}

	return nil
}

func (m *WeeeDataProductRenderWeeeAdjustmentAttributeInterface) validateAttributeCode(formats strfmt.Registry) error {

	if err := validate.Required("attribute_code", "body", m.AttributeCode); err != nil {
		return err
	}

	return nil
}

func (m *WeeeDataProductRenderWeeeAdjustmentAttributeInterface) validateExtensionAttributes(formats strfmt.Registry) error {

	if m.ExtensionAttributes == nil {
		return errors.Required("extension_attributes", "body", nil)
	}

	return nil
}

func (m *WeeeDataProductRenderWeeeAdjustmentAttributeInterface) validateTaxAmount(formats strfmt.Registry) error {

	if err := validate.Required("tax_amount", "body", m.TaxAmount); err != nil {
		return err
	}

	return nil
}

func (m *WeeeDataProductRenderWeeeAdjustmentAttributeInterface) validateTaxAmountInclTax(formats strfmt.Registry) error {

	if err := validate.Required("tax_amount_incl_tax", "body", m.TaxAmountInclTax); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this weee data product render weee adjustment attribute interface based on context it is used
func (m *WeeeDataProductRenderWeeeAdjustmentAttributeInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WeeeDataProductRenderWeeeAdjustmentAttributeInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WeeeDataProductRenderWeeeAdjustmentAttributeInterface) UnmarshalBinary(b []byte) error {
	var res WeeeDataProductRenderWeeeAdjustmentAttributeInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
