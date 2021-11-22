// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_cart_total_management_v1

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

// QuoteGuestCartTotalManagementV1CollectTotalsPutReader is a Reader for the QuoteGuestCartTotalManagementV1CollectTotalsPut structure.
type QuoteGuestCartTotalManagementV1CollectTotalsPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuoteGuestCartTotalManagementV1CollectTotalsPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQuoteGuestCartTotalManagementV1CollectTotalsPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteGuestCartTotalManagementV1CollectTotalsPutOK creates a QuoteGuestCartTotalManagementV1CollectTotalsPutOK with default headers values
func NewQuoteGuestCartTotalManagementV1CollectTotalsPutOK() *QuoteGuestCartTotalManagementV1CollectTotalsPutOK {
	return &QuoteGuestCartTotalManagementV1CollectTotalsPutOK{}
}

/* QuoteGuestCartTotalManagementV1CollectTotalsPutOK describes a response with status code 200, with default header values.

200 Success.
*/
type QuoteGuestCartTotalManagementV1CollectTotalsPutOK struct {
	Payload *models.QuoteDataTotalsInterface
}

func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutOK) Error() string {
	return fmt.Sprintf("[PUT /V1/guest-carts/{cartId}/collect-totals][%d] quoteGuestCartTotalManagementV1CollectTotalsPutOK  %+v", 200, o.Payload)
}
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutOK) GetPayload() *models.QuoteDataTotalsInterface {
	return o.Payload
}

func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuoteDataTotalsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestCartTotalManagementV1CollectTotalsPutDefault creates a QuoteGuestCartTotalManagementV1CollectTotalsPutDefault with default headers values
func NewQuoteGuestCartTotalManagementV1CollectTotalsPutDefault(code int) *QuoteGuestCartTotalManagementV1CollectTotalsPutDefault {
	return &QuoteGuestCartTotalManagementV1CollectTotalsPutDefault{
		_statusCode: code,
	}
}

/* QuoteGuestCartTotalManagementV1CollectTotalsPutDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type QuoteGuestCartTotalManagementV1CollectTotalsPutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote guest cart total management v1 collect totals put default response
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutDefault) Code() int {
	return o._statusCode
}

func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/guest-carts/{cartId}/collect-totals][%d] quoteGuestCartTotalManagementV1CollectTotalsPut default  %+v", o._statusCode, o.Payload)
}
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*QuoteGuestCartTotalManagementV1CollectTotalsPutBody quote guest cart total management v1 collect totals put body
swagger:model QuoteGuestCartTotalManagementV1CollectTotalsPutBody
*/
type QuoteGuestCartTotalManagementV1CollectTotalsPutBody struct {

	// additional data
	AdditionalData *models.QuoteDataTotalsAdditionalDataInterface `json:"additionalData,omitempty"`

	// payment method
	// Required: true
	PaymentMethod *models.QuoteDataPaymentInterface `json:"paymentMethod"`

	// The carrier code.
	ShippingCarrierCode string `json:"shippingCarrierCode,omitempty"`

	// The shipping method code.
	ShippingMethodCode string `json:"shippingMethodCode,omitempty"`
}

// Validate validates this quote guest cart total management v1 collect totals put body
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAdditionalData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePaymentMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutBody) validateAdditionalData(formats strfmt.Registry) error {
	if swag.IsZero(o.AdditionalData) { // not required
		return nil
	}

	if o.AdditionalData != nil {
		if err := o.AdditionalData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quoteGuestCartTotalManagementV1CollectTotalsPutBody" + "." + "additionalData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quoteGuestCartTotalManagementV1CollectTotalsPutBody" + "." + "additionalData")
			}
			return err
		}
	}

	return nil
}

func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutBody) validatePaymentMethod(formats strfmt.Registry) error {

	if err := validate.Required("quoteGuestCartTotalManagementV1CollectTotalsPutBody"+"."+"paymentMethod", "body", o.PaymentMethod); err != nil {
		return err
	}

	if o.PaymentMethod != nil {
		if err := o.PaymentMethod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quoteGuestCartTotalManagementV1CollectTotalsPutBody" + "." + "paymentMethod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quoteGuestCartTotalManagementV1CollectTotalsPutBody" + "." + "paymentMethod")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this quote guest cart total management v1 collect totals put body based on the context it is used
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAdditionalData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePaymentMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutBody) contextValidateAdditionalData(ctx context.Context, formats strfmt.Registry) error {

	if o.AdditionalData != nil {
		if err := o.AdditionalData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quoteGuestCartTotalManagementV1CollectTotalsPutBody" + "." + "additionalData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quoteGuestCartTotalManagementV1CollectTotalsPutBody" + "." + "additionalData")
			}
			return err
		}
	}

	return nil
}

func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutBody) contextValidatePaymentMethod(ctx context.Context, formats strfmt.Registry) error {

	if o.PaymentMethod != nil {
		if err := o.PaymentMethod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quoteGuestCartTotalManagementV1CollectTotalsPutBody" + "." + "paymentMethod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quoteGuestCartTotalManagementV1CollectTotalsPutBody" + "." + "paymentMethod")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QuoteGuestCartTotalManagementV1CollectTotalsPutBody) UnmarshalBinary(b []byte) error {
	var res QuoteGuestCartTotalManagementV1CollectTotalsPutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
