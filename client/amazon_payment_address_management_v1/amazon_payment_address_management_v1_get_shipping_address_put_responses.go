// Code generated by go-swagger; DO NOT EDIT.

package amazon_payment_address_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"magento2go/models"
)

// AmazonPaymentAddressManagementV1GetShippingAddressPutReader is a Reader for the AmazonPaymentAddressManagementV1GetShippingAddressPut structure.
type AmazonPaymentAddressManagementV1GetShippingAddressPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AmazonPaymentAddressManagementV1GetShippingAddressPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAmazonPaymentAddressManagementV1GetShippingAddressPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAmazonPaymentAddressManagementV1GetShippingAddressPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAmazonPaymentAddressManagementV1GetShippingAddressPutOK creates a AmazonPaymentAddressManagementV1GetShippingAddressPutOK with default headers values
func NewAmazonPaymentAddressManagementV1GetShippingAddressPutOK() *AmazonPaymentAddressManagementV1GetShippingAddressPutOK {
	return &AmazonPaymentAddressManagementV1GetShippingAddressPutOK{}
}

/* AmazonPaymentAddressManagementV1GetShippingAddressPutOK describes a response with status code 200, with default header values.

200 Success.
*/
type AmazonPaymentAddressManagementV1GetShippingAddressPutOK struct {
	Payload string
}

func (o *AmazonPaymentAddressManagementV1GetShippingAddressPutOK) Error() string {
	return fmt.Sprintf("[PUT /V1/amazon-shipping-address/{amazonOrderReferenceId}][%d] amazonPaymentAddressManagementV1GetShippingAddressPutOK  %+v", 200, o.Payload)
}
func (o *AmazonPaymentAddressManagementV1GetShippingAddressPutOK) GetPayload() string {
	return o.Payload
}

func (o *AmazonPaymentAddressManagementV1GetShippingAddressPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAmazonPaymentAddressManagementV1GetShippingAddressPutDefault creates a AmazonPaymentAddressManagementV1GetShippingAddressPutDefault with default headers values
func NewAmazonPaymentAddressManagementV1GetShippingAddressPutDefault(code int) *AmazonPaymentAddressManagementV1GetShippingAddressPutDefault {
	return &AmazonPaymentAddressManagementV1GetShippingAddressPutDefault{
		_statusCode: code,
	}
}

/* AmazonPaymentAddressManagementV1GetShippingAddressPutDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type AmazonPaymentAddressManagementV1GetShippingAddressPutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the amazon payment address management v1 get shipping address put default response
func (o *AmazonPaymentAddressManagementV1GetShippingAddressPutDefault) Code() int {
	return o._statusCode
}

func (o *AmazonPaymentAddressManagementV1GetShippingAddressPutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/amazon-shipping-address/{amazonOrderReferenceId}][%d] amazonPaymentAddressManagementV1GetShippingAddressPut default  %+v", o._statusCode, o.Payload)
}
func (o *AmazonPaymentAddressManagementV1GetShippingAddressPutDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AmazonPaymentAddressManagementV1GetShippingAddressPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AmazonPaymentAddressManagementV1GetShippingAddressPutBody amazon payment address management v1 get shipping address put body
swagger:model AmazonPaymentAddressManagementV1GetShippingAddressPutBody
*/
type AmazonPaymentAddressManagementV1GetShippingAddressPutBody struct {

	// address consent token
	// Required: true
	AddressConsentToken *string `json:"addressConsentToken"`
}

// Validate validates this amazon payment address management v1 get shipping address put body
func (o *AmazonPaymentAddressManagementV1GetShippingAddressPutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAddressConsentToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AmazonPaymentAddressManagementV1GetShippingAddressPutBody) validateAddressConsentToken(formats strfmt.Registry) error {

	if err := validate.Required("amazonPaymentAddressManagementV1GetShippingAddressPutBody"+"."+"addressConsentToken", "body", o.AddressConsentToken); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this amazon payment address management v1 get shipping address put body based on context it is used
func (o *AmazonPaymentAddressManagementV1GetShippingAddressPutBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AmazonPaymentAddressManagementV1GetShippingAddressPutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AmazonPaymentAddressManagementV1GetShippingAddressPutBody) UnmarshalBinary(b []byte) error {
	var res AmazonPaymentAddressManagementV1GetShippingAddressPutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}