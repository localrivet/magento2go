// Code generated by go-swagger; DO NOT EDIT.

package catalog_category_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magento2go/models"
)

// CatalogCategoryRepositoryV1GetGetReader is a Reader for the CatalogCategoryRepositoryV1GetGet structure.
type CatalogCategoryRepositoryV1GetGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogCategoryRepositoryV1GetGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogCategoryRepositoryV1GetGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogCategoryRepositoryV1GetGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCatalogCategoryRepositoryV1GetGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogCategoryRepositoryV1GetGetOK creates a CatalogCategoryRepositoryV1GetGetOK with default headers values
func NewCatalogCategoryRepositoryV1GetGetOK() *CatalogCategoryRepositoryV1GetGetOK {
	return &CatalogCategoryRepositoryV1GetGetOK{}
}

/* CatalogCategoryRepositoryV1GetGetOK describes a response with status code 200, with default header values.

200 Success.
*/
type CatalogCategoryRepositoryV1GetGetOK struct {
	Payload *models.CatalogDataCategoryInterface
}

func (o *CatalogCategoryRepositoryV1GetGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/categories/{categoryId}][%d] catalogCategoryRepositoryV1GetGetOK  %+v", 200, o.Payload)
}
func (o *CatalogCategoryRepositoryV1GetGetOK) GetPayload() *models.CatalogDataCategoryInterface {
	return o.Payload
}

func (o *CatalogCategoryRepositoryV1GetGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogDataCategoryInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCategoryRepositoryV1GetGetBadRequest creates a CatalogCategoryRepositoryV1GetGetBadRequest with default headers values
func NewCatalogCategoryRepositoryV1GetGetBadRequest() *CatalogCategoryRepositoryV1GetGetBadRequest {
	return &CatalogCategoryRepositoryV1GetGetBadRequest{}
}

/* CatalogCategoryRepositoryV1GetGetBadRequest describes a response with status code 400, with default header values.

400 Bad Request
*/
type CatalogCategoryRepositoryV1GetGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CatalogCategoryRepositoryV1GetGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/categories/{categoryId}][%d] catalogCategoryRepositoryV1GetGetBadRequest  %+v", 400, o.Payload)
}
func (o *CatalogCategoryRepositoryV1GetGetBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CatalogCategoryRepositoryV1GetGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCategoryRepositoryV1GetGetDefault creates a CatalogCategoryRepositoryV1GetGetDefault with default headers values
func NewCatalogCategoryRepositoryV1GetGetDefault(code int) *CatalogCategoryRepositoryV1GetGetDefault {
	return &CatalogCategoryRepositoryV1GetGetDefault{
		_statusCode: code,
	}
}

/* CatalogCategoryRepositoryV1GetGetDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type CatalogCategoryRepositoryV1GetGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog category repository v1 get get default response
func (o *CatalogCategoryRepositoryV1GetGetDefault) Code() int {
	return o._statusCode
}

func (o *CatalogCategoryRepositoryV1GetGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/categories/{categoryId}][%d] catalogCategoryRepositoryV1GetGet default  %+v", o._statusCode, o.Payload)
}
func (o *CatalogCategoryRepositoryV1GetGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CatalogCategoryRepositoryV1GetGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
