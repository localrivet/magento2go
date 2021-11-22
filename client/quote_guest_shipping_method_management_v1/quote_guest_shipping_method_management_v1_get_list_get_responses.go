// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_shipping_method_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/localrivet/magento2go/models"
)

// QuoteGuestShippingMethodManagementV1GetListGetReader is a Reader for the QuoteGuestShippingMethodManagementV1GetListGet structure.
type QuoteGuestShippingMethodManagementV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteGuestShippingMethodManagementV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuoteGuestShippingMethodManagementV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQuoteGuestShippingMethodManagementV1GetListGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQuoteGuestShippingMethodManagementV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteGuestShippingMethodManagementV1GetListGetOK creates a QuoteGuestShippingMethodManagementV1GetListGetOK with default headers values
func NewQuoteGuestShippingMethodManagementV1GetListGetOK() *QuoteGuestShippingMethodManagementV1GetListGetOK {
	return &QuoteGuestShippingMethodManagementV1GetListGetOK{}
}

/* QuoteGuestShippingMethodManagementV1GetListGetOK describes a response with status code 200, with default header values.

200 Success.
*/
type QuoteGuestShippingMethodManagementV1GetListGetOK struct {
	Payload []*models.QuoteDataShippingMethodInterface
}

func (o *QuoteGuestShippingMethodManagementV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/guest-carts/{cartId}/shipping-methods][%d] quoteGuestShippingMethodManagementV1GetListGetOK  %+v", 200, o.Payload)
}
func (o *QuoteGuestShippingMethodManagementV1GetListGetOK) GetPayload() []*models.QuoteDataShippingMethodInterface {
	return o.Payload
}

func (o *QuoteGuestShippingMethodManagementV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestShippingMethodManagementV1GetListGetBadRequest creates a QuoteGuestShippingMethodManagementV1GetListGetBadRequest with default headers values
func NewQuoteGuestShippingMethodManagementV1GetListGetBadRequest() *QuoteGuestShippingMethodManagementV1GetListGetBadRequest {
	return &QuoteGuestShippingMethodManagementV1GetListGetBadRequest{}
}

/* QuoteGuestShippingMethodManagementV1GetListGetBadRequest describes a response with status code 400, with default header values.

400 Bad Request
*/
type QuoteGuestShippingMethodManagementV1GetListGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteGuestShippingMethodManagementV1GetListGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/guest-carts/{cartId}/shipping-methods][%d] quoteGuestShippingMethodManagementV1GetListGetBadRequest  %+v", 400, o.Payload)
}
func (o *QuoteGuestShippingMethodManagementV1GetListGetBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QuoteGuestShippingMethodManagementV1GetListGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestShippingMethodManagementV1GetListGetDefault creates a QuoteGuestShippingMethodManagementV1GetListGetDefault with default headers values
func NewQuoteGuestShippingMethodManagementV1GetListGetDefault(code int) *QuoteGuestShippingMethodManagementV1GetListGetDefault {
	return &QuoteGuestShippingMethodManagementV1GetListGetDefault{
		_statusCode: code,
	}
}

/* QuoteGuestShippingMethodManagementV1GetListGetDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type QuoteGuestShippingMethodManagementV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote guest shipping method management v1 get list get default response
func (o *QuoteGuestShippingMethodManagementV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *QuoteGuestShippingMethodManagementV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/guest-carts/{cartId}/shipping-methods][%d] quoteGuestShippingMethodManagementV1GetListGet default  %+v", o._statusCode, o.Payload)
}
func (o *QuoteGuestShippingMethodManagementV1GetListGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QuoteGuestShippingMethodManagementV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
