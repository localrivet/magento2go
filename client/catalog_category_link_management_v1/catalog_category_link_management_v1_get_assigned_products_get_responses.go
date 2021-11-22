// Code generated by go-swagger; DO NOT EDIT.

package catalog_category_link_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/localrivet/magento2go/models"
)

// CatalogCategoryLinkManagementV1GetAssignedProductsGetReader is a Reader for the CatalogCategoryLinkManagementV1GetAssignedProductsGet structure.
type CatalogCategoryLinkManagementV1GetAssignedProductsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogCategoryLinkManagementV1GetAssignedProductsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogCategoryLinkManagementV1GetAssignedProductsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCatalogCategoryLinkManagementV1GetAssignedProductsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogCategoryLinkManagementV1GetAssignedProductsGetOK creates a CatalogCategoryLinkManagementV1GetAssignedProductsGetOK with default headers values
func NewCatalogCategoryLinkManagementV1GetAssignedProductsGetOK() *CatalogCategoryLinkManagementV1GetAssignedProductsGetOK {
	return &CatalogCategoryLinkManagementV1GetAssignedProductsGetOK{}
}

/* CatalogCategoryLinkManagementV1GetAssignedProductsGetOK describes a response with status code 200, with default header values.

200 Success.
*/
type CatalogCategoryLinkManagementV1GetAssignedProductsGetOK struct {
	Payload []*models.CatalogDataCategoryProductLinkInterface
}

func (o *CatalogCategoryLinkManagementV1GetAssignedProductsGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/categories/{categoryId}/products][%d] catalogCategoryLinkManagementV1GetAssignedProductsGetOK  %+v", 200, o.Payload)
}
func (o *CatalogCategoryLinkManagementV1GetAssignedProductsGetOK) GetPayload() []*models.CatalogDataCategoryProductLinkInterface {
	return o.Payload
}

func (o *CatalogCategoryLinkManagementV1GetAssignedProductsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogCategoryLinkManagementV1GetAssignedProductsGetDefault creates a CatalogCategoryLinkManagementV1GetAssignedProductsGetDefault with default headers values
func NewCatalogCategoryLinkManagementV1GetAssignedProductsGetDefault(code int) *CatalogCategoryLinkManagementV1GetAssignedProductsGetDefault {
	return &CatalogCategoryLinkManagementV1GetAssignedProductsGetDefault{
		_statusCode: code,
	}
}

/* CatalogCategoryLinkManagementV1GetAssignedProductsGetDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type CatalogCategoryLinkManagementV1GetAssignedProductsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalog category link management v1 get assigned products get default response
func (o *CatalogCategoryLinkManagementV1GetAssignedProductsGetDefault) Code() int {
	return o._statusCode
}

func (o *CatalogCategoryLinkManagementV1GetAssignedProductsGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/categories/{categoryId}/products][%d] catalogCategoryLinkManagementV1GetAssignedProductsGet default  %+v", o._statusCode, o.Payload)
}
func (o *CatalogCategoryLinkManagementV1GetAssignedProductsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CatalogCategoryLinkManagementV1GetAssignedProductsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
