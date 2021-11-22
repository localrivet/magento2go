// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_render_list_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magento2go/models"
)

// CatalogProductRenderListV1GetListGetReader is a Reader for the CatalogProductRenderListV1GetListGet structure.
type CatalogProductRenderListV1GetListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogProductRenderListV1GetListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogProductRenderListV1GetListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCatalogProductRenderListV1GetListGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogProductRenderListV1GetListGetOK creates a CatalogProductRenderListV1GetListGetOK with default headers values
func NewCatalogProductRenderListV1GetListGetOK() *CatalogProductRenderListV1GetListGetOK {
	return &CatalogProductRenderListV1GetListGetOK{}
}

/* CatalogProductRenderListV1GetListGetOK describes a response with status code 200, with default header values.

200 Success.
*/
type CatalogProductRenderListV1GetListGetOK struct {
	Payload *models.CatalogDataProductRenderSearchResultsInterface
}

func (o *CatalogProductRenderListV1GetListGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/products-render-info][%d] catalogProductRenderListV1GetListGetOK  %+v", 200, o.Payload)
}
func (o *CatalogProductRenderListV1GetListGetOK) GetPayload() *models.CatalogDataProductRenderSearchResultsInterface {
	return o.Payload
}

func (o *CatalogProductRenderListV1GetListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogDataProductRenderSearchResultsInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogProductRenderListV1GetListGetDefault creates a CatalogProductRenderListV1GetListGetDefault with default headers values
func NewCatalogProductRenderListV1GetListGetDefault(code int) *CatalogProductRenderListV1GetListGetDefault {
	return &CatalogProductRenderListV1GetListGetDefault{
		_statusCode: code,
	}
}

/* CatalogProductRenderListV1GetListGetDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type CatalogProductRenderListV1GetListGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog product render list v1 get list get default response
func (o *CatalogProductRenderListV1GetListGetDefault) Code() int {
	return o._statusCode
}

func (o *CatalogProductRenderListV1GetListGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/products-render-info][%d] catalogProductRenderListV1GetListGet default  %+v", o._statusCode, o.Payload)
}
func (o *CatalogProductRenderListV1GetListGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CatalogProductRenderListV1GetListGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
