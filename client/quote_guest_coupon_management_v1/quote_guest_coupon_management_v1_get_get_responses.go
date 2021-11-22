// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_coupon_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/localrivet/magento2go/models"
)

// QuoteGuestCouponManagementV1GetGetReader is a Reader for the QuoteGuestCouponManagementV1GetGet structure.
type QuoteGuestCouponManagementV1GetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteGuestCouponManagementV1GetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuoteGuestCouponManagementV1GetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQuoteGuestCouponManagementV1GetGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQuoteGuestCouponManagementV1GetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteGuestCouponManagementV1GetGetOK creates a QuoteGuestCouponManagementV1GetGetOK with default headers values
func NewQuoteGuestCouponManagementV1GetGetOK() *QuoteGuestCouponManagementV1GetGetOK {
	return &QuoteGuestCouponManagementV1GetGetOK{}
}

/* QuoteGuestCouponManagementV1GetGetOK describes a response with status code 200, with default header values.

200 Success.
*/
type QuoteGuestCouponManagementV1GetGetOK struct {
	Payload string
}

func (o *QuoteGuestCouponManagementV1GetGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/guest-carts/{cartId}/coupons][%d] quoteGuestCouponManagementV1GetGetOK  %+v", 200, o.Payload)
}
func (o *QuoteGuestCouponManagementV1GetGetOK) GetPayload() string {
	return o.Payload
}

func (o *QuoteGuestCouponManagementV1GetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestCouponManagementV1GetGetBadRequest creates a QuoteGuestCouponManagementV1GetGetBadRequest with default headers values
func NewQuoteGuestCouponManagementV1GetGetBadRequest() *QuoteGuestCouponManagementV1GetGetBadRequest {
	return &QuoteGuestCouponManagementV1GetGetBadRequest{}
}

/* QuoteGuestCouponManagementV1GetGetBadRequest describes a response with status code 400, with default header values.

400 Bad Request
*/
type QuoteGuestCouponManagementV1GetGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteGuestCouponManagementV1GetGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/guest-carts/{cartId}/coupons][%d] quoteGuestCouponManagementV1GetGetBadRequest  %+v", 400, o.Payload)
}
func (o *QuoteGuestCouponManagementV1GetGetBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QuoteGuestCouponManagementV1GetGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestCouponManagementV1GetGetDefault creates a QuoteGuestCouponManagementV1GetGetDefault with default headers values
func NewQuoteGuestCouponManagementV1GetGetDefault(code int) *QuoteGuestCouponManagementV1GetGetDefault {
	return &QuoteGuestCouponManagementV1GetGetDefault{
		_statusCode: code,
	}
}

/* QuoteGuestCouponManagementV1GetGetDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type QuoteGuestCouponManagementV1GetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote guest coupon management v1 get get default response
func (o *QuoteGuestCouponManagementV1GetGetDefault) Code() int {
	return o._statusCode
}

func (o *QuoteGuestCouponManagementV1GetGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/guest-carts/{cartId}/coupons][%d] quoteGuestCouponManagementV1GetGet default  %+v", o._statusCode, o.Payload)
}
func (o *QuoteGuestCouponManagementV1GetGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QuoteGuestCouponManagementV1GetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
