// Code generated by go-swagger; DO NOT EDIT.

package configurable_product_option_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magento2go/models"
)

// ConfigurableProductOptionRepositoryV1GetGetReader is a Reader for the ConfigurableProductOptionRepositoryV1GetGet structure.
type ConfigurableProductOptionRepositoryV1GetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigurableProductOptionRepositoryV1GetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigurableProductOptionRepositoryV1GetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConfigurableProductOptionRepositoryV1GetGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewConfigurableProductOptionRepositoryV1GetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConfigurableProductOptionRepositoryV1GetGetOK creates a ConfigurableProductOptionRepositoryV1GetGetOK with default headers values
func NewConfigurableProductOptionRepositoryV1GetGetOK() *ConfigurableProductOptionRepositoryV1GetGetOK {
	return &ConfigurableProductOptionRepositoryV1GetGetOK{}
}

/* ConfigurableProductOptionRepositoryV1GetGetOK describes a response with status code 200, with default header values.

200 Success.
*/
type ConfigurableProductOptionRepositoryV1GetGetOK struct {
	Payload *models.ConfigurableProductDataOptionInterface
}

func (o *ConfigurableProductOptionRepositoryV1GetGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/configurable-products/{sku}/options/{id}][%d] configurableProductOptionRepositoryV1GetGetOK  %+v", 200, o.Payload)
}
func (o *ConfigurableProductOptionRepositoryV1GetGetOK) GetPayload() *models.ConfigurableProductDataOptionInterface {
	return o.Payload
}

func (o *ConfigurableProductOptionRepositoryV1GetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigurableProductDataOptionInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigurableProductOptionRepositoryV1GetGetBadRequest creates a ConfigurableProductOptionRepositoryV1GetGetBadRequest with default headers values
func NewConfigurableProductOptionRepositoryV1GetGetBadRequest() *ConfigurableProductOptionRepositoryV1GetGetBadRequest {
	return &ConfigurableProductOptionRepositoryV1GetGetBadRequest{}
}

/* ConfigurableProductOptionRepositoryV1GetGetBadRequest describes a response with status code 400, with default header values.

400 Bad Request
*/
type ConfigurableProductOptionRepositoryV1GetGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ConfigurableProductOptionRepositoryV1GetGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/configurable-products/{sku}/options/{id}][%d] configurableProductOptionRepositoryV1GetGetBadRequest  %+v", 400, o.Payload)
}
func (o *ConfigurableProductOptionRepositoryV1GetGetBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConfigurableProductOptionRepositoryV1GetGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigurableProductOptionRepositoryV1GetGetDefault creates a ConfigurableProductOptionRepositoryV1GetGetDefault with default headers values
func NewConfigurableProductOptionRepositoryV1GetGetDefault(code int) *ConfigurableProductOptionRepositoryV1GetGetDefault {
	return &ConfigurableProductOptionRepositoryV1GetGetDefault{
		_statusCode: code,
	}
}

/* ConfigurableProductOptionRepositoryV1GetGetDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type ConfigurableProductOptionRepositoryV1GetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the configurable product option repository v1 get get default response
func (o *ConfigurableProductOptionRepositoryV1GetGetDefault) Code() int {
	return o._statusCode
}

func (o *ConfigurableProductOptionRepositoryV1GetGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/configurable-products/{sku}/options/{id}][%d] configurableProductOptionRepositoryV1GetGet default  %+v", o._statusCode, o.Payload)
}
func (o *ConfigurableProductOptionRepositoryV1GetGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConfigurableProductOptionRepositoryV1GetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}