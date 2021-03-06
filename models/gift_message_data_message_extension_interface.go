// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GiftMessageDataMessageExtensionInterface ExtensionInterface class for @see \Magento\GiftMessage\Api\Data\MessageInterface
//
// swagger:model gift-message-data-message-extension-interface
type GiftMessageDataMessageExtensionInterface struct {

	// entity id
	EntityID string `json:"entity_id,omitempty"`

	// entity type
	EntityType string `json:"entity_type,omitempty"`
}

// Validate validates this gift message data message extension interface
func (m *GiftMessageDataMessageExtensionInterface) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gift message data message extension interface based on context it is used
func (m *GiftMessageDataMessageExtensionInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GiftMessageDataMessageExtensionInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GiftMessageDataMessageExtensionInterface) UnmarshalBinary(b []byte) error {
	var res GiftMessageDataMessageExtensionInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
