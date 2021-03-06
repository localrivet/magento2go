// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_cart_item_repository_v1

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

	"github.com/localrivet/magento2go/models"
)

// QuoteGuestCartItemRepositoryV1SavePostReader is a Reader for the QuoteGuestCartItemRepositoryV1SavePost structure.
type QuoteGuestCartItemRepositoryV1SavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteGuestCartItemRepositoryV1SavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuoteGuestCartItemRepositoryV1SavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQuoteGuestCartItemRepositoryV1SavePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQuoteGuestCartItemRepositoryV1SavePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuoteGuestCartItemRepositoryV1SavePostOK creates a QuoteGuestCartItemRepositoryV1SavePostOK with default headers values
func NewQuoteGuestCartItemRepositoryV1SavePostOK() *QuoteGuestCartItemRepositoryV1SavePostOK {
	return &QuoteGuestCartItemRepositoryV1SavePostOK{}
}

/* QuoteGuestCartItemRepositoryV1SavePostOK describes a response with status code 200, with default header values.

200 Success.
*/
type QuoteGuestCartItemRepositoryV1SavePostOK struct {
	Payload *models.QuoteDataCartItemInterface
}

func (o *QuoteGuestCartItemRepositoryV1SavePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/guest-carts/{cartId}/items][%d] quoteGuestCartItemRepositoryV1SavePostOK  %+v", 200, o.Payload)
}
func (o *QuoteGuestCartItemRepositoryV1SavePostOK) GetPayload() *models.QuoteDataCartItemInterface {
	return o.Payload
}

func (o *QuoteGuestCartItemRepositoryV1SavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuoteDataCartItemInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestCartItemRepositoryV1SavePostBadRequest creates a QuoteGuestCartItemRepositoryV1SavePostBadRequest with default headers values
func NewQuoteGuestCartItemRepositoryV1SavePostBadRequest() *QuoteGuestCartItemRepositoryV1SavePostBadRequest {
	return &QuoteGuestCartItemRepositoryV1SavePostBadRequest{}
}

/* QuoteGuestCartItemRepositoryV1SavePostBadRequest describes a response with status code 400, with default header values.

400 Bad Request
*/
type QuoteGuestCartItemRepositoryV1SavePostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QuoteGuestCartItemRepositoryV1SavePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/guest-carts/{cartId}/items][%d] quoteGuestCartItemRepositoryV1SavePostBadRequest  %+v", 400, o.Payload)
}
func (o *QuoteGuestCartItemRepositoryV1SavePostBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QuoteGuestCartItemRepositoryV1SavePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGuestCartItemRepositoryV1SavePostDefault creates a QuoteGuestCartItemRepositoryV1SavePostDefault with default headers values
func NewQuoteGuestCartItemRepositoryV1SavePostDefault(code int) *QuoteGuestCartItemRepositoryV1SavePostDefault {
	return &QuoteGuestCartItemRepositoryV1SavePostDefault{
		_statusCode: code,
	}
}

/* QuoteGuestCartItemRepositoryV1SavePostDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type QuoteGuestCartItemRepositoryV1SavePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quote guest cart item repository v1 save post default response
func (o *QuoteGuestCartItemRepositoryV1SavePostDefault) Code() int {
	return o._statusCode
}

func (o *QuoteGuestCartItemRepositoryV1SavePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/guest-carts/{cartId}/items][%d] quoteGuestCartItemRepositoryV1SavePost default  %+v", o._statusCode, o.Payload)
}
func (o *QuoteGuestCartItemRepositoryV1SavePostDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QuoteGuestCartItemRepositoryV1SavePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*QuoteGuestCartItemRepositoryV1SavePostBody quote guest cart item repository v1 save post body
swagger:model QuoteGuestCartItemRepositoryV1SavePostBody
*/
type QuoteGuestCartItemRepositoryV1SavePostBody struct {

	// cart item
	// Required: true
	CartItem *models.QuoteDataCartItemInterface `json:"cartItem"`
}

// Validate validates this quote guest cart item repository v1 save post body
func (o *QuoteGuestCartItemRepositoryV1SavePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCartItem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QuoteGuestCartItemRepositoryV1SavePostBody) validateCartItem(formats strfmt.Registry) error {

	if err := validate.Required("quoteGuestCartItemRepositoryV1SavePostBody"+"."+"cartItem", "body", o.CartItem); err != nil {
		return err
	}

	if o.CartItem != nil {
		if err := o.CartItem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quoteGuestCartItemRepositoryV1SavePostBody" + "." + "cartItem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quoteGuestCartItemRepositoryV1SavePostBody" + "." + "cartItem")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this quote guest cart item repository v1 save post body based on the context it is used
func (o *QuoteGuestCartItemRepositoryV1SavePostBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCartItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QuoteGuestCartItemRepositoryV1SavePostBody) contextValidateCartItem(ctx context.Context, formats strfmt.Registry) error {

	if o.CartItem != nil {
		if err := o.CartItem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quoteGuestCartItemRepositoryV1SavePostBody" + "." + "cartItem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quoteGuestCartItemRepositoryV1SavePostBody" + "." + "cartItem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *QuoteGuestCartItemRepositoryV1SavePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QuoteGuestCartItemRepositoryV1SavePostBody) UnmarshalBinary(b []byte) error {
	var res QuoteGuestCartItemRepositoryV1SavePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
