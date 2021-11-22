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

	"magento2go/models"
)

// CustomerAccountManagementV1CreateAccountPostReader is a Reader for the CustomerAccountManagementV1CreateAccountPost structure.
type CustomerAccountManagementV1CreateAccountPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerAccountManagementV1CreateAccountPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerAccountManagementV1CreateAccountPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewCustomerAccountManagementV1CreateAccountPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCustomerAccountManagementV1CreateAccountPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerAccountManagementV1CreateAccountPostOK creates a CustomerAccountManagementV1CreateAccountPostOK with default headers values
func NewCustomerAccountManagementV1CreateAccountPostOK() *CustomerAccountManagementV1CreateAccountPostOK {
	return &CustomerAccountManagementV1CreateAccountPostOK{}
}

/* CustomerAccountManagementV1CreateAccountPostOK describes a response with status code 200, with default header values.

200 Success.
*/
type CustomerAccountManagementV1CreateAccountPostOK struct {
	Payload *models.CustomerDataCustomerInterface
}

func (o *CustomerAccountManagementV1CreateAccountPostOK) Error() string {
	return fmt.Sprintf("[POST /V1/customers][%d] customerAccountManagementV1CreateAccountPostOK  %+v", 200, o.Payload)
}
func (o *CustomerAccountManagementV1CreateAccountPostOK) GetPayload() *models.CustomerDataCustomerInterface {
	return o.Payload
}

func (o *CustomerAccountManagementV1CreateAccountPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerDataCustomerInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountManagementV1CreateAccountPostInternalServerError creates a CustomerAccountManagementV1CreateAccountPostInternalServerError with default headers values
func NewCustomerAccountManagementV1CreateAccountPostInternalServerError() *CustomerAccountManagementV1CreateAccountPostInternalServerError {
	return &CustomerAccountManagementV1CreateAccountPostInternalServerError{}
}

/* CustomerAccountManagementV1CreateAccountPostInternalServerError describes a response with status code 500, with default header values.

Internal Server error
*/
type CustomerAccountManagementV1CreateAccountPostInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CustomerAccountManagementV1CreateAccountPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /V1/customers][%d] customerAccountManagementV1CreateAccountPostInternalServerError  %+v", 500, o.Payload)
}
func (o *CustomerAccountManagementV1CreateAccountPostInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CustomerAccountManagementV1CreateAccountPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerAccountManagementV1CreateAccountPostDefault creates a CustomerAccountManagementV1CreateAccountPostDefault with default headers values
func NewCustomerAccountManagementV1CreateAccountPostDefault(code int) *CustomerAccountManagementV1CreateAccountPostDefault {
	return &CustomerAccountManagementV1CreateAccountPostDefault{
		_statusCode: code,
	}
}

/* CustomerAccountManagementV1CreateAccountPostDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type CustomerAccountManagementV1CreateAccountPostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the customer account management v1 create account post default response
func (o *CustomerAccountManagementV1CreateAccountPostDefault) Code() int {
	return o._statusCode
}

func (o *CustomerAccountManagementV1CreateAccountPostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/customers][%d] customerAccountManagementV1CreateAccountPost default  %+v", o._statusCode, o.Payload)
}
func (o *CustomerAccountManagementV1CreateAccountPostDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CustomerAccountManagementV1CreateAccountPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CustomerAccountManagementV1CreateAccountPostBody customer account management v1 create account post body
swagger:model CustomerAccountManagementV1CreateAccountPostBody
*/
type CustomerAccountManagementV1CreateAccountPostBody struct {

	// customer
	// Required: true
	Customer *models.CustomerDataCustomerInterface `json:"customer"`

	// password
	Password string `json:"password,omitempty"`

	// redirect Url
	RedirectURL string `json:"redirectUrl,omitempty"`
}

// Validate validates this customer account management v1 create account post body
func (o *CustomerAccountManagementV1CreateAccountPostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CustomerAccountManagementV1CreateAccountPostBody) validateCustomer(formats strfmt.Registry) error {

	if err := validate.Required("customerAccountManagementV1CreateAccountPostBody"+"."+"customer", "body", o.Customer); err != nil {
		return err
	}

	if o.Customer != nil {
		if err := o.Customer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customerAccountManagementV1CreateAccountPostBody" + "." + "customer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customerAccountManagementV1CreateAccountPostBody" + "." + "customer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this customer account management v1 create account post body based on the context it is used
func (o *CustomerAccountManagementV1CreateAccountPostBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCustomer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CustomerAccountManagementV1CreateAccountPostBody) contextValidateCustomer(ctx context.Context, formats strfmt.Registry) error {

	if o.Customer != nil {
		if err := o.Customer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customerAccountManagementV1CreateAccountPostBody" + "." + "customer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customerAccountManagementV1CreateAccountPostBody" + "." + "customer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CustomerAccountManagementV1CreateAccountPostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CustomerAccountManagementV1CreateAccountPostBody) UnmarshalBinary(b []byte) error {
	var res CustomerAccountManagementV1CreateAccountPostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
