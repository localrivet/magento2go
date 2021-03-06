// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomerDataCustomerExtensionInterface ExtensionInterface class for @see \Magento\Customer\Api\Data\CustomerInterface
//
// swagger:model customer-data-customer-extension-interface
type CustomerDataCustomerExtensionInterface struct {

	// amazon id
	AmazonID string `json:"amazon_id,omitempty"`

	// assistance allowed
	AssistanceAllowed int64 `json:"assistance_allowed,omitempty"`

	// is subscribed
	IsSubscribed bool `json:"is_subscribed,omitempty"`

	// vertex customer code
	VertexCustomerCode string `json:"vertex_customer_code,omitempty"`

	// vertex customer country
	VertexCustomerCountry string `json:"vertex_customer_country,omitempty"`
}

// Validate validates this customer data customer extension interface
func (m *CustomerDataCustomerExtensionInterface) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this customer data customer extension interface based on context it is used
func (m *CustomerDataCustomerExtensionInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomerDataCustomerExtensionInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerDataCustomerExtensionInterface) UnmarshalBinary(b []byte) error {
	var res CustomerDataCustomerExtensionInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
