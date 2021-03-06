// Code generated by go-swagger; DO NOT EDIT.

package gift_message_guest_cart_repository_v1

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

// GiftMessageGuestCartRepositoryV1SavePostReader is a Reader for the GiftMessageGuestCartRepositoryV1SavePost structure.
type GiftMessageGuestCartRepositoryV1SavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GiftMessageGuestCartRepositoryV1SavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGiftMessageGuestCartRepositoryV1SavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGiftMessageGuestCartRepositoryV1SavePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGiftMessageGuestCartRepositoryV1SavePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGiftMessageGuestCartRepositoryV1SavePostOK creates a GiftMessageGuestCartRepositoryV1SavePostOK with default headers values
func NewGiftMessageGuestCartRepositoryV1SavePostOK() *GiftMessageGuestCartRepositoryV1SavePostOK {
	return &GiftMessageGuestCartRepositoryV1SavePostOK{}
}

/* GiftMessageGuestCartRepositoryV1SavePostOK describes a response with status code 200, with default header values.

200 Success.
*/
type GiftMessageGuestCartRepositoryV1SavePostOK struct {
	Payload bool
}

func (o *GiftMessageGuestCartRepositoryV1SavePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/guest-carts/{cartId}/gift-message][%d] giftMessageGuestCartRepositoryV1SavePostOK  %+v", 200, o.Payload)
}
func (o *GiftMessageGuestCartRepositoryV1SavePostOK) GetPayload() bool {
	return o.Payload
}

func (o *GiftMessageGuestCartRepositoryV1SavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftMessageGuestCartRepositoryV1SavePostBadRequest creates a GiftMessageGuestCartRepositoryV1SavePostBadRequest with default headers values
func NewGiftMessageGuestCartRepositoryV1SavePostBadRequest() *GiftMessageGuestCartRepositoryV1SavePostBadRequest {
	return &GiftMessageGuestCartRepositoryV1SavePostBadRequest{}
}

/* GiftMessageGuestCartRepositoryV1SavePostBadRequest describes a response with status code 400, with default header values.

400 Bad Request
*/
type GiftMessageGuestCartRepositoryV1SavePostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GiftMessageGuestCartRepositoryV1SavePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/guest-carts/{cartId}/gift-message][%d] giftMessageGuestCartRepositoryV1SavePostBadRequest  %+v", 400, o.Payload)
}
func (o *GiftMessageGuestCartRepositoryV1SavePostBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GiftMessageGuestCartRepositoryV1SavePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftMessageGuestCartRepositoryV1SavePostDefault creates a GiftMessageGuestCartRepositoryV1SavePostDefault with default headers values
func NewGiftMessageGuestCartRepositoryV1SavePostDefault(code int) *GiftMessageGuestCartRepositoryV1SavePostDefault {
	return &GiftMessageGuestCartRepositoryV1SavePostDefault{
		_statusCode: code,
	}
}

/* GiftMessageGuestCartRepositoryV1SavePostDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type GiftMessageGuestCartRepositoryV1SavePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the gift message guest cart repository v1 save post default response
func (o *GiftMessageGuestCartRepositoryV1SavePostDefault) Code() int {
	return o._statusCode
}

func (o *GiftMessageGuestCartRepositoryV1SavePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/guest-carts/{cartId}/gift-message][%d] giftMessageGuestCartRepositoryV1SavePost default  %+v", o._statusCode, o.Payload)
}
func (o *GiftMessageGuestCartRepositoryV1SavePostDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GiftMessageGuestCartRepositoryV1SavePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GiftMessageGuestCartRepositoryV1SavePostBody gift message guest cart repository v1 save post body
swagger:model GiftMessageGuestCartRepositoryV1SavePostBody
*/
type GiftMessageGuestCartRepositoryV1SavePostBody struct {

	// gift message
	// Required: true
	GiftMessage *models.GiftMessageDataMessageInterface `json:"giftMessage"`
}

// Validate validates this gift message guest cart repository v1 save post body
func (o *GiftMessageGuestCartRepositoryV1SavePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGiftMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GiftMessageGuestCartRepositoryV1SavePostBody) validateGiftMessage(formats strfmt.Registry) error {

	if err := validate.Required("giftMessageGuestCartRepositoryV1SavePostBody"+"."+"giftMessage", "body", o.GiftMessage); err != nil {
		return err
	}

	if o.GiftMessage != nil {
		if err := o.GiftMessage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("giftMessageGuestCartRepositoryV1SavePostBody" + "." + "giftMessage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("giftMessageGuestCartRepositoryV1SavePostBody" + "." + "giftMessage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this gift message guest cart repository v1 save post body based on the context it is used
func (o *GiftMessageGuestCartRepositoryV1SavePostBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateGiftMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GiftMessageGuestCartRepositoryV1SavePostBody) contextValidateGiftMessage(ctx context.Context, formats strfmt.Registry) error {

	if o.GiftMessage != nil {
		if err := o.GiftMessage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("giftMessageGuestCartRepositoryV1SavePostBody" + "." + "giftMessage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("giftMessageGuestCartRepositoryV1SavePostBody" + "." + "giftMessage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GiftMessageGuestCartRepositoryV1SavePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GiftMessageGuestCartRepositoryV1SavePostBody) UnmarshalBinary(b []byte) error {
	var res GiftMessageGuestCartRepositoryV1SavePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
