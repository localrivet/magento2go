// Code generated by go-swagger; DO NOT EDIT.

package checkout_guest_totals_information_management_v1

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

// CheckoutGuestTotalsInformationManagementV1CalculatePostReader is a Reader for the CheckoutGuestTotalsInformationManagementV1CalculatePost structure.
type CheckoutGuestTotalsInformationManagementV1CalculatePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckoutGuestTotalsInformationManagementV1CalculatePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCheckoutGuestTotalsInformationManagementV1CalculatePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCheckoutGuestTotalsInformationManagementV1CalculatePostOK creates a CheckoutGuestTotalsInformationManagementV1CalculatePostOK with default headers values
func NewCheckoutGuestTotalsInformationManagementV1CalculatePostOK() *CheckoutGuestTotalsInformationManagementV1CalculatePostOK {
	return &CheckoutGuestTotalsInformationManagementV1CalculatePostOK{}
}

/* CheckoutGuestTotalsInformationManagementV1CalculatePostOK describes a response with status code 200, with default header values.

200 Success.
*/
type CheckoutGuestTotalsInformationManagementV1CalculatePostOK struct {
	Payload *models.QuoteDataTotalsInterface
}

func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostOK) Error() string {
	return fmt.Sprintf("[POST /V1/guest-carts/{cartId}/totals-information][%d] checkoutGuestTotalsInformationManagementV1CalculatePostOK  %+v", 200, o.Payload)
}
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostOK) GetPayload() *models.QuoteDataTotalsInterface {
	return o.Payload
}

func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuoteDataTotalsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckoutGuestTotalsInformationManagementV1CalculatePostDefault creates a CheckoutGuestTotalsInformationManagementV1CalculatePostDefault with default headers values
func NewCheckoutGuestTotalsInformationManagementV1CalculatePostDefault(code int) *CheckoutGuestTotalsInformationManagementV1CalculatePostDefault {
	return &CheckoutGuestTotalsInformationManagementV1CalculatePostDefault{
		_statusCode: code,
	}
}

/* CheckoutGuestTotalsInformationManagementV1CalculatePostDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type CheckoutGuestTotalsInformationManagementV1CalculatePostDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the checkout guest totals information management v1 calculate post default response
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostDefault) Code() int {
	return o._statusCode
}

func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostDefault) Error() string {
	return fmt.Sprintf("[POST /V1/guest-carts/{cartId}/totals-information][%d] checkoutGuestTotalsInformationManagementV1CalculatePost default  %+v", o._statusCode, o.Payload)
}
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CheckoutGuestTotalsInformationManagementV1CalculatePostBody checkout guest totals information management v1 calculate post body
swagger:model CheckoutGuestTotalsInformationManagementV1CalculatePostBody
*/
type CheckoutGuestTotalsInformationManagementV1CalculatePostBody struct {

	// address information
	// Required: true
	AddressInformation *models.CheckoutDataTotalsInformationInterface `json:"addressInformation"`
}

// Validate validates this checkout guest totals information management v1 calculate post body
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAddressInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostBody) validateAddressInformation(formats strfmt.Registry) error {

	if err := validate.Required("checkoutGuestTotalsInformationManagementV1CalculatePostBody"+"."+"addressInformation", "body", o.AddressInformation); err != nil {
		return err
	}

	if o.AddressInformation != nil {
		if err := o.AddressInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkoutGuestTotalsInformationManagementV1CalculatePostBody" + "." + "addressInformation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checkoutGuestTotalsInformationManagementV1CalculatePostBody" + "." + "addressInformation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this checkout guest totals information management v1 calculate post body based on the context it is used
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAddressInformation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostBody) contextValidateAddressInformation(ctx context.Context, formats strfmt.Registry) error {

	if o.AddressInformation != nil {
		if err := o.AddressInformation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkoutGuestTotalsInformationManagementV1CalculatePostBody" + "." + "addressInformation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checkoutGuestTotalsInformationManagementV1CalculatePostBody" + "." + "addressInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckoutGuestTotalsInformationManagementV1CalculatePostBody) UnmarshalBinary(b []byte) error {
	var res CheckoutGuestTotalsInformationManagementV1CalculatePostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
