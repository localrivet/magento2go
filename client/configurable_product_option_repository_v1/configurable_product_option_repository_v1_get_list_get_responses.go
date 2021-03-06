// Code generated by go-swagger; DO NOT EDIT.

package configurable_product_option_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/localrivet/magento2go/models"
)

// ConfigurableProductOptionRepositoryV1GetListGetReader is a Reader for the ConfigurableProductOptionRepositoryV1GetListGet structure.
type ConfigurableProductOptionRepositoryV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigurableProductOptionRepositoryV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigurableProductOptionRepositoryV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConfigurableProductOptionRepositoryV1GetListGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewConfigurableProductOptionRepositoryV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConfigurableProductOptionRepositoryV1GetListGetOK creates a ConfigurableProductOptionRepositoryV1GetListGetOK with default headers values
func NewConfigurableProductOptionRepositoryV1GetListGetOK() *ConfigurableProductOptionRepositoryV1GetListGetOK {
	return &ConfigurableProductOptionRepositoryV1GetListGetOK{}
}

/* ConfigurableProductOptionRepositoryV1GetListGetOK describes a response with status code 200, with default header values.

200 Success.
*/
type ConfigurableProductOptionRepositoryV1GetListGetOK struct {
	Payload []*models.ConfigurableProductDataOptionInterface
}

func (o *ConfigurableProductOptionRepositoryV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/configurable-products/{sku}/options/all][%d] configurableProductOptionRepositoryV1GetListGetOK  %+v", 200, o.Payload)
}
func (o *ConfigurableProductOptionRepositoryV1GetListGetOK) GetPayload() []*models.ConfigurableProductDataOptionInterface {
	return o.Payload
}

func (o *ConfigurableProductOptionRepositoryV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigurableProductOptionRepositoryV1GetListGetBadRequest creates a ConfigurableProductOptionRepositoryV1GetListGetBadRequest with default headers values
func NewConfigurableProductOptionRepositoryV1GetListGetBadRequest() *ConfigurableProductOptionRepositoryV1GetListGetBadRequest {
	return &ConfigurableProductOptionRepositoryV1GetListGetBadRequest{}
}

/* ConfigurableProductOptionRepositoryV1GetListGetBadRequest describes a response with status code 400, with default header values.

400 Bad Request
*/
type ConfigurableProductOptionRepositoryV1GetListGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ConfigurableProductOptionRepositoryV1GetListGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/configurable-products/{sku}/options/all][%d] configurableProductOptionRepositoryV1GetListGetBadRequest  %+v", 400, o.Payload)
}
func (o *ConfigurableProductOptionRepositoryV1GetListGetBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConfigurableProductOptionRepositoryV1GetListGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigurableProductOptionRepositoryV1GetListGetDefault creates a ConfigurableProductOptionRepositoryV1GetListGetDefault with default headers values
func NewConfigurableProductOptionRepositoryV1GetListGetDefault(code int) *ConfigurableProductOptionRepositoryV1GetListGetDefault {
	return &ConfigurableProductOptionRepositoryV1GetListGetDefault{
		_statusCode: code,
	}
}

/* ConfigurableProductOptionRepositoryV1GetListGetDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type ConfigurableProductOptionRepositoryV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the configurable product option repository v1 get list get default response
func (o *ConfigurableProductOptionRepositoryV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *ConfigurableProductOptionRepositoryV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/configurable-products/{sku}/options/all][%d] configurableProductOptionRepositoryV1GetListGet default  %+v", o._statusCode, o.Payload)
}
func (o *ConfigurableProductOptionRepositoryV1GetListGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConfigurableProductOptionRepositoryV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
