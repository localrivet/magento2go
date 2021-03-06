// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/localrivet/magento2go/models"
)

// CatalogProductRepositoryV1GetGetReader is a Reader for the CatalogProductRepositoryV1GetGet structure.
type CatalogProductRepositoryV1GetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogProductRepositoryV1GetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogProductRepositoryV1GetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogProductRepositoryV1GetGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCatalogProductRepositoryV1GetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogProductRepositoryV1GetGetOK creates a CatalogProductRepositoryV1GetGetOK with default headers values
func NewCatalogProductRepositoryV1GetGetOK() *CatalogProductRepositoryV1GetGetOK {
	return &CatalogProductRepositoryV1GetGetOK{}
}

/* CatalogProductRepositoryV1GetGetOK describes a response with status code 200, with default header values.

200 Success.
*/
type CatalogProductRepositoryV1GetGetOK struct {
	Payload *models.CatalogDataProductInterface
}

func (o *CatalogProductRepositoryV1GetGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/products/{sku}][%d] catalogProductRepositoryV1GetGetOK  %+v", 200, o.Payload)
}
func (o *CatalogProductRepositoryV1GetGetOK) GetPayload() *models.CatalogDataProductInterface {
	return o.Payload
}

func (o *CatalogProductRepositoryV1GetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogDataProductInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductRepositoryV1GetGetBadRequest creates a CatalogProductRepositoryV1GetGetBadRequest with default headers values
func NewCatalogProductRepositoryV1GetGetBadRequest() *CatalogProductRepositoryV1GetGetBadRequest {
	return &CatalogProductRepositoryV1GetGetBadRequest{}
}

/* CatalogProductRepositoryV1GetGetBadRequest describes a response with status code 400, with default header values.

400 Bad Request
*/
type CatalogProductRepositoryV1GetGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CatalogProductRepositoryV1GetGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/products/{sku}][%d] catalogProductRepositoryV1GetGetBadRequest  %+v", 400, o.Payload)
}
func (o *CatalogProductRepositoryV1GetGetBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CatalogProductRepositoryV1GetGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductRepositoryV1GetGetDefault creates a CatalogProductRepositoryV1GetGetDefault with default headers values
func NewCatalogProductRepositoryV1GetGetDefault(code int) *CatalogProductRepositoryV1GetGetDefault {
	return &CatalogProductRepositoryV1GetGetDefault{
		_statusCode: code,
	}
}

/* CatalogProductRepositoryV1GetGetDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type CatalogProductRepositoryV1GetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog product repository v1 get get default response
func (o *CatalogProductRepositoryV1GetGetDefault) Code() int {
	return o._statusCode
}

func (o *CatalogProductRepositoryV1GetGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/products/{sku}][%d] catalogProductRepositoryV1GetGet default  %+v", o._statusCode, o.Payload)
}
func (o *CatalogProductRepositoryV1GetGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CatalogProductRepositoryV1GetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
