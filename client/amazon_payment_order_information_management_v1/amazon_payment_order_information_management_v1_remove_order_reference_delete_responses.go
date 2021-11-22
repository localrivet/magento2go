// Code generated by go-swagger; DO NOT EDIT.

package amazon_payment_order_information_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magento2go/models"
)

// AmazonPaymentOrderInformationManagementV1RemoveOrderReferenceDeleteReader is a Reader for the AmazonPaymentOrderInformationManagementV1RemoveOrderReferenceDelete structure.
type AmazonPaymentOrderInformationManagementV1RemoveOrderReferenceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AmazonPaymentOrderInformationManagementV1RemoveOrderReferenceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewAmazonPaymentOrderInformationManagementV1RemoveOrderReferenceDeleteDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewAmazonPaymentOrderInformationManagementV1RemoveOrderReferenceDeleteDefault creates a AmazonPaymentOrderInformationManagementV1RemoveOrderReferenceDeleteDefault with default headers values
func NewAmazonPaymentOrderInformationManagementV1RemoveOrderReferenceDeleteDefault(code int) *AmazonPaymentOrderInformationManagementV1RemoveOrderReferenceDeleteDefault {
	return &AmazonPaymentOrderInformationManagementV1RemoveOrderReferenceDeleteDefault{
		_statusCode: code,
	}
}

/* AmazonPaymentOrderInformationManagementV1RemoveOrderReferenceDeleteDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type AmazonPaymentOrderInformationManagementV1RemoveOrderReferenceDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the amazon payment order information management v1 remove order reference delete default response
func (o *AmazonPaymentOrderInformationManagementV1RemoveOrderReferenceDeleteDefault) Code() int {
	return o._statusCode
}

func (o *AmazonPaymentOrderInformationManagementV1RemoveOrderReferenceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /V1/amazon/order-ref][%d] amazonPaymentOrderInformationManagementV1RemoveOrderReferenceDelete default  %+v", o._statusCode, o.Payload)
}
func (o *AmazonPaymentOrderInformationManagementV1RemoveOrderReferenceDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AmazonPaymentOrderInformationManagementV1RemoveOrderReferenceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
