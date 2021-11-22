// Code generated by go-swagger; DO NOT EDIT.

package catalog_category_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/localrivet/magento2go/models"
)

// CatalogCategoryManagementV1GetTreeGetReader is a Reader for the CatalogCategoryManagementV1GetTreeGet structure.
type CatalogCategoryManagementV1GetTreeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogCategoryManagementV1GetTreeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogCategoryManagementV1GetTreeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCatalogCategoryManagementV1GetTreeGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCatalogCategoryManagementV1GetTreeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogCategoryManagementV1GetTreeGetOK creates a CatalogCategoryManagementV1GetTreeGetOK with default headers values
func NewCatalogCategoryManagementV1GetTreeGetOK() *CatalogCategoryManagementV1GetTreeGetOK {
	return &CatalogCategoryManagementV1GetTreeGetOK{}
}

/* CatalogCategoryManagementV1GetTreeGetOK describes a response with status code 200, with default header values.

200 Success.
*/
type CatalogCategoryManagementV1GetTreeGetOK struct {
	Payload *models.CatalogDataCategoryTreeInterface
}

func (o *CatalogCategoryManagementV1GetTreeGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/categories][%d] catalogCategoryManagementV1GetTreeGetOK  %+v", 200, o.Payload)
}
func (o *CatalogCategoryManagementV1GetTreeGetOK) GetPayload() *models.CatalogDataCategoryTreeInterface {
	return o.Payload
}

func (o *CatalogCategoryManagementV1GetTreeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogDataCategoryTreeInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCategoryManagementV1GetTreeGetBadRequest creates a CatalogCategoryManagementV1GetTreeGetBadRequest with default headers values
func NewCatalogCategoryManagementV1GetTreeGetBadRequest() *CatalogCategoryManagementV1GetTreeGetBadRequest {
	return &CatalogCategoryManagementV1GetTreeGetBadRequest{}
}

/* CatalogCategoryManagementV1GetTreeGetBadRequest describes a response with status code 400, with default header values.

400 Bad Request
*/
type CatalogCategoryManagementV1GetTreeGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CatalogCategoryManagementV1GetTreeGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/categories][%d] catalogCategoryManagementV1GetTreeGetBadRequest  %+v", 400, o.Payload)
}
func (o *CatalogCategoryManagementV1GetTreeGetBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CatalogCategoryManagementV1GetTreeGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCategoryManagementV1GetTreeGetDefault creates a CatalogCategoryManagementV1GetTreeGetDefault with default headers values
func NewCatalogCategoryManagementV1GetTreeGetDefault(code int) *CatalogCategoryManagementV1GetTreeGetDefault {
	return &CatalogCategoryManagementV1GetTreeGetDefault{
		_statusCode: code,
	}
}

/* CatalogCategoryManagementV1GetTreeGetDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type CatalogCategoryManagementV1GetTreeGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog category management v1 get tree get default response
func (o *CatalogCategoryManagementV1GetTreeGetDefault) Code() int {
	return o._statusCode
}

func (o *CatalogCategoryManagementV1GetTreeGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/categories][%d] catalogCategoryManagementV1GetTreeGet default  %+v", o._statusCode, o.Payload)
}
func (o *CatalogCategoryManagementV1GetTreeGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CatalogCategoryManagementV1GetTreeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
