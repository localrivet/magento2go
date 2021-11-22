// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_custom_option_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magento2go/models"
)

// CatalogProductCustomOptionRepositoryV1GetGetReader is a Reader for the CatalogProductCustomOptionRepositoryV1GetGet structure.
type CatalogProductCustomOptionRepositoryV1GetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogProductCustomOptionRepositoryV1GetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogProductCustomOptionRepositoryV1GetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCatalogProductCustomOptionRepositoryV1GetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogProductCustomOptionRepositoryV1GetGetOK creates a CatalogProductCustomOptionRepositoryV1GetGetOK with default headers values
func NewCatalogProductCustomOptionRepositoryV1GetGetOK() *CatalogProductCustomOptionRepositoryV1GetGetOK {
	return &CatalogProductCustomOptionRepositoryV1GetGetOK{}
}

/* CatalogProductCustomOptionRepositoryV1GetGetOK describes a response with status code 200, with default header values.

200 Success.
*/
type CatalogProductCustomOptionRepositoryV1GetGetOK struct {
	Payload *models.CatalogDataProductCustomOptionInterface
}

func (o *CatalogProductCustomOptionRepositoryV1GetGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/products/{sku}/options/{optionId}][%d] catalogProductCustomOptionRepositoryV1GetGetOK  %+v", 200, o.Payload)
}
func (o *CatalogProductCustomOptionRepositoryV1GetGetOK) GetPayload() *models.CatalogDataProductCustomOptionInterface {
	return o.Payload
}

func (o *CatalogProductCustomOptionRepositoryV1GetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogDataProductCustomOptionInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductCustomOptionRepositoryV1GetGetDefault creates a CatalogProductCustomOptionRepositoryV1GetGetDefault with default headers values
func NewCatalogProductCustomOptionRepositoryV1GetGetDefault(code int) *CatalogProductCustomOptionRepositoryV1GetGetDefault {
	return &CatalogProductCustomOptionRepositoryV1GetGetDefault{
		_statusCode: code,
	}
}

/* CatalogProductCustomOptionRepositoryV1GetGetDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type CatalogProductCustomOptionRepositoryV1GetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog product custom option repository v1 get get default response
func (o *CatalogProductCustomOptionRepositoryV1GetGetDefault) Code() int {
	return o._statusCode
}

func (o *CatalogProductCustomOptionRepositoryV1GetGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/products/{sku}/options/{optionId}][%d] catalogProductCustomOptionRepositoryV1GetGet default  %+v", o._statusCode, o.Payload)
}
func (o *CatalogProductCustomOptionRepositoryV1GetGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CatalogProductCustomOptionRepositoryV1GetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
