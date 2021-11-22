// Code generated by go-swagger; DO NOT EDIT.

package gift_message_guest_item_repository_v1

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

// GiftMessageGuestItemRepositoryV1SavePostReader is a Reader for the GiftMessageGuestItemRepositoryV1SavePost structure.
type GiftMessageGuestItemRepositoryV1SavePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GiftMessageGuestItemRepositoryV1SavePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGiftMessageGuestItemRepositoryV1SavePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGiftMessageGuestItemRepositoryV1SavePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGiftMessageGuestItemRepositoryV1SavePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGiftMessageGuestItemRepositoryV1SavePostOK creates a GiftMessageGuestItemRepositoryV1SavePostOK with default headers values
func NewGiftMessageGuestItemRepositoryV1SavePostOK() *GiftMessageGuestItemRepositoryV1SavePostOK {
	return &GiftMessageGuestItemRepositoryV1SavePostOK{}
}

/* GiftMessageGuestItemRepositoryV1SavePostOK describes a response with status code 200, with default header values.

200 Success.
*/
type GiftMessageGuestItemRepositoryV1SavePostOK struct {
	Payload bool
}

func (o *GiftMessageGuestItemRepositoryV1SavePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/guest-carts/{cartId}/gift-message/{itemId}][%d] giftMessageGuestItemRepositoryV1SavePostOK  %+v", 200, o.Payload)
}
func (o *GiftMessageGuestItemRepositoryV1SavePostOK) GetPayload() bool {
	return o.Payload
}

func (o *GiftMessageGuestItemRepositoryV1SavePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftMessageGuestItemRepositoryV1SavePostBadRequest creates a GiftMessageGuestItemRepositoryV1SavePostBadRequest with default headers values
func NewGiftMessageGuestItemRepositoryV1SavePostBadRequest() *GiftMessageGuestItemRepositoryV1SavePostBadRequest {
	return &GiftMessageGuestItemRepositoryV1SavePostBadRequest{}
}

/* GiftMessageGuestItemRepositoryV1SavePostBadRequest describes a response with status code 400, with default header values.

400 Bad Request
*/
type GiftMessageGuestItemRepositoryV1SavePostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GiftMessageGuestItemRepositoryV1SavePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /V1/guest-carts/{cartId}/gift-message/{itemId}][%d] giftMessageGuestItemRepositoryV1SavePostBadRequest  %+v", 400, o.Payload)
}
func (o *GiftMessageGuestItemRepositoryV1SavePostBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GiftMessageGuestItemRepositoryV1SavePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGiftMessageGuestItemRepositoryV1SavePostDefault creates a GiftMessageGuestItemRepositoryV1SavePostDefault with default headers values
func NewGiftMessageGuestItemRepositoryV1SavePostDefault(code int) *GiftMessageGuestItemRepositoryV1SavePostDefault {
	return &GiftMessageGuestItemRepositoryV1SavePostDefault{
		_statusCode: code,
	}
}

/* GiftMessageGuestItemRepositoryV1SavePostDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type GiftMessageGuestItemRepositoryV1SavePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the gift message guest item repository v1 save post default response
func (o *GiftMessageGuestItemRepositoryV1SavePostDefault) Code() int {
	return o._statusCode
}

func (o *GiftMessageGuestItemRepositoryV1SavePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/guest-carts/{cartId}/gift-message/{itemId}][%d] giftMessageGuestItemRepositoryV1SavePost default  %+v", o._statusCode, o.Payload)
}
func (o *GiftMessageGuestItemRepositoryV1SavePostDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GiftMessageGuestItemRepositoryV1SavePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GiftMessageGuestItemRepositoryV1SavePostBody gift message guest item repository v1 save post body
swagger:model GiftMessageGuestItemRepositoryV1SavePostBody
*/
type GiftMessageGuestItemRepositoryV1SavePostBody struct {

	// gift message
	// Required: true
	GiftMessage *models.GiftMessageDataMessageInterface `json:"giftMessage"`
}

// Validate validates this gift message guest item repository v1 save post body
func (o *GiftMessageGuestItemRepositoryV1SavePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGiftMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GiftMessageGuestItemRepositoryV1SavePostBody) validateGiftMessage(formats strfmt.Registry) error {

	if err := validate.Required("giftMessageGuestItemRepositoryV1SavePostBody"+"."+"giftMessage", "body", o.GiftMessage); err != nil {
		return err
	}

	if o.GiftMessage != nil {
		if err := o.GiftMessage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("giftMessageGuestItemRepositoryV1SavePostBody" + "." + "giftMessage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("giftMessageGuestItemRepositoryV1SavePostBody" + "." + "giftMessage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this gift message guest item repository v1 save post body based on the context it is used
func (o *GiftMessageGuestItemRepositoryV1SavePostBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateGiftMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GiftMessageGuestItemRepositoryV1SavePostBody) contextValidateGiftMessage(ctx context.Context, formats strfmt.Registry) error {

	if o.GiftMessage != nil {
		if err := o.GiftMessage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("giftMessageGuestItemRepositoryV1SavePostBody" + "." + "giftMessage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("giftMessageGuestItemRepositoryV1SavePostBody" + "." + "giftMessage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GiftMessageGuestItemRepositoryV1SavePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GiftMessageGuestItemRepositoryV1SavePostBody) UnmarshalBinary(b []byte) error {
	var res GiftMessageGuestItemRepositoryV1SavePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
