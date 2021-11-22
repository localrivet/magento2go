// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_cart_total_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magento2go/models"
)

// QuoteGuestCartTotalRepositoryV1GetGetReader is a Reader for the QuoteGuestCartTotalRepositoryV1GetGet structure.
type QuoteGuestCartTotalRepositoryV1GetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteGuestCartTotalRepositoryV1GetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuoteGuestCartTotalRepositoryV1GetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQuoteGuestCartTotalRepositoryV1GetGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQuoteGuestCartTotalRepositoryV1GetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteGuestCartTotalRepositoryV1GetGetOK creates a QuoteGuestCartTotalRepositoryV1GetGetOK with default headers values
func NewQuoteGuestCartTotalRepositoryV1GetGetOK() *QuoteGuestCartTotalRepositoryV1GetGetOK {
	return &QuoteGuestCartTotalRepositoryV1GetGetOK{}
}

/* QuoteGuestCartTotalRepositoryV1GetGetOK describes a response with status code 200, with default header values.

200 Success.
*/
type QuoteGuestCartTotalRepositoryV1GetGetOK struct {
	Payload *models.QuoteDataTotalsInterface
}

func (o *QuoteGuestCartTotalRepositoryV1GetGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/guest-carts/{cartId}/totals][%d] quoteGuestCartTotalRepositoryV1GetGetOK  %+v", 200, o.Payload)
}
func (o *QuoteGuestCartTotalRepositoryV1GetGetOK) GetPayload() *models.QuoteDataTotalsInterface {
	return o.Payload
}

func (o *QuoteGuestCartTotalRepositoryV1GetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuoteDataTotalsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestCartTotalRepositoryV1GetGetBadRequest creates a QuoteGuestCartTotalRepositoryV1GetGetBadRequest with default headers values
func NewQuoteGuestCartTotalRepositoryV1GetGetBadRequest() *QuoteGuestCartTotalRepositoryV1GetGetBadRequest {
	return &QuoteGuestCartTotalRepositoryV1GetGetBadRequest{}
}

/* QuoteGuestCartTotalRepositoryV1GetGetBadRequest describes a response with status code 400, with default header values.

400 Bad Request
*/
type QuoteGuestCartTotalRepositoryV1GetGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteGuestCartTotalRepositoryV1GetGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/guest-carts/{cartId}/totals][%d] quoteGuestCartTotalRepositoryV1GetGetBadRequest  %+v", 400, o.Payload)
}
func (o *QuoteGuestCartTotalRepositoryV1GetGetBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QuoteGuestCartTotalRepositoryV1GetGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestCartTotalRepositoryV1GetGetDefault creates a QuoteGuestCartTotalRepositoryV1GetGetDefault with default headers values
func NewQuoteGuestCartTotalRepositoryV1GetGetDefault(code int) *QuoteGuestCartTotalRepositoryV1GetGetDefault {
	return &QuoteGuestCartTotalRepositoryV1GetGetDefault{
		_statusCode: code,
	}
}

/* QuoteGuestCartTotalRepositoryV1GetGetDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type QuoteGuestCartTotalRepositoryV1GetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote guest cart total repository v1 get get default response
func (o *QuoteGuestCartTotalRepositoryV1GetGetDefault) Code() int {
	return o._statusCode
}

func (o *QuoteGuestCartTotalRepositoryV1GetGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/guest-carts/{cartId}/totals][%d] quoteGuestCartTotalRepositoryV1GetGet default  %+v", o._statusCode, o.Payload)
}
func (o *QuoteGuestCartTotalRepositoryV1GetGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QuoteGuestCartTotalRepositoryV1GetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
