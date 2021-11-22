// Code generated by go-swagger; DO NOT EDIT.

package customer_account_management_v1

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

// CustomerAccountManagementV1InitiatePasswordResetPutReader is a Reader for the CustomerAccountManagementV1InitiatePasswordResetPut structure.
type CustomerAccountManagementV1InitiatePasswordResetPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerAccountManagementV1InitiatePasswordResetPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerAccountManagementV1InitiatePasswordResetPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewCustomerAccountManagementV1InitiatePasswordResetPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCustomerAccountManagementV1InitiatePasswordResetPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerAccountManagementV1InitiatePasswordResetPutOK creates a CustomerAccountManagementV1InitiatePasswordResetPutOK with default headers values
func NewCustomerAccountManagementV1InitiatePasswordResetPutOK() *CustomerAccountManagementV1InitiatePasswordResetPutOK {
	return &CustomerAccountManagementV1InitiatePasswordResetPutOK{}
}

/* CustomerAccountManagementV1InitiatePasswordResetPutOK describes a response with status code 200, with default header values.

200 Success.
*/
type CustomerAccountManagementV1InitiatePasswordResetPutOK struct {
	Payload bool
}

func (o *CustomerAccountManagementV1InitiatePasswordResetPutOK) Error() string {
	return fmt.Sprintf("[PUT /V1/customers/password][%d] customerAccountManagementV1InitiatePasswordResetPutOK  %+v", 200, o.Payload)
}
func (o *CustomerAccountManagementV1InitiatePasswordResetPutOK) GetPayload() bool {
	return o.Payload
}

func (o *CustomerAccountManagementV1InitiatePasswordResetPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountManagementV1InitiatePasswordResetPutInternalServerError creates a CustomerAccountManagementV1InitiatePasswordResetPutInternalServerError with default headers values
func NewCustomerAccountManagementV1InitiatePasswordResetPutInternalServerError() *CustomerAccountManagementV1InitiatePasswordResetPutInternalServerError {
	return &CustomerAccountManagementV1InitiatePasswordResetPutInternalServerError{}
}

/* CustomerAccountManagementV1InitiatePasswordResetPutInternalServerError describes a response with status code 500, with default header values.

Internal Server error
*/
type CustomerAccountManagementV1InitiatePasswordResetPutInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CustomerAccountManagementV1InitiatePasswordResetPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /V1/customers/password][%d] customerAccountManagementV1InitiatePasswordResetPutInternalServerError  %+v", 500, o.Payload)
}
func (o *CustomerAccountManagementV1InitiatePasswordResetPutInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CustomerAccountManagementV1InitiatePasswordResetPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountManagementV1InitiatePasswordResetPutDefault creates a CustomerAccountManagementV1InitiatePasswordResetPutDefault with default headers values
func NewCustomerAccountManagementV1InitiatePasswordResetPutDefault(code int) *CustomerAccountManagementV1InitiatePasswordResetPutDefault {
	return &CustomerAccountManagementV1InitiatePasswordResetPutDefault{
		_statusCode: code,
	}
}

/* CustomerAccountManagementV1InitiatePasswordResetPutDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type CustomerAccountManagementV1InitiatePasswordResetPutDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the customer account management v1 initiate password reset put default response
func (o *CustomerAccountManagementV1InitiatePasswordResetPutDefault) Code() int {
	return o._statusCode
}

func (o *CustomerAccountManagementV1InitiatePasswordResetPutDefault) Error() string {
	return fmt.Sprintf("[PUT /V1/customers/password][%d] customerAccountManagementV1InitiatePasswordResetPut default  %+v", o._statusCode, o.Payload)
}
func (o *CustomerAccountManagementV1InitiatePasswordResetPutDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CustomerAccountManagementV1InitiatePasswordResetPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CustomerAccountManagementV1InitiatePasswordResetPutBody customer account management v1 initiate password reset put body
swagger:model CustomerAccountManagementV1InitiatePasswordResetPutBody
*/
type CustomerAccountManagementV1InitiatePasswordResetPutBody struct {

	// email
	// Required: true
	Email *string `json:"email"`

	// template
	// Required: true
	Template *string `json:"template"`

	// website Id
	WebsiteID int64 `json:"websiteId,omitempty"`
}

// Validate validates this customer account management v1 initiate password reset put body
func (o *CustomerAccountManagementV1InitiatePasswordResetPutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CustomerAccountManagementV1InitiatePasswordResetPutBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("customerAccountManagementV1InitiatePasswordResetPutBody"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *CustomerAccountManagementV1InitiatePasswordResetPutBody) validateTemplate(formats strfmt.Registry) error {

	if err := validate.Required("customerAccountManagementV1InitiatePasswordResetPutBody"+"."+"template", "body", o.Template); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this customer account management v1 initiate password reset put body based on context it is used
func (o *CustomerAccountManagementV1InitiatePasswordResetPutBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CustomerAccountManagementV1InitiatePasswordResetPutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CustomerAccountManagementV1InitiatePasswordResetPutBody) UnmarshalBinary(b []byte) error {
	var res CustomerAccountManagementV1InitiatePasswordResetPutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
