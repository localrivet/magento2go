// Code generated by go-swagger; DO NOT EDIT.

package cms_page_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/localrivet/magento2go/models"
)

// CmsPageRepositoryV1GetByIDGetReader is a Reader for the CmsPageRepositoryV1GetByIDGet structure.
type CmsPageRepositoryV1GetByIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CmsPageRepositoryV1GetByIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCmsPageRepositoryV1GetByIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewCmsPageRepositoryV1GetByIDGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCmsPageRepositoryV1GetByIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCmsPageRepositoryV1GetByIDGetOK creates a CmsPageRepositoryV1GetByIDGetOK with default headers values
func NewCmsPageRepositoryV1GetByIDGetOK() *CmsPageRepositoryV1GetByIDGetOK {
	return &CmsPageRepositoryV1GetByIDGetOK{}
}

/* CmsPageRepositoryV1GetByIDGetOK describes a response with status code 200, with default header values.

200 Success.
*/
type CmsPageRepositoryV1GetByIDGetOK struct {
	Payload *models.CmsDataPageInterface
}

func (o *CmsPageRepositoryV1GetByIDGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/cmsPage/{pageId}][%d] cmsPageRepositoryV1GetByIdGetOK  %+v", 200, o.Payload)
}
func (o *CmsPageRepositoryV1GetByIDGetOK) GetPayload() *models.CmsDataPageInterface {
	return o.Payload
}

func (o *CmsPageRepositoryV1GetByIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CmsDataPageInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCmsPageRepositoryV1GetByIDGetInternalServerError creates a CmsPageRepositoryV1GetByIDGetInternalServerError with default headers values
func NewCmsPageRepositoryV1GetByIDGetInternalServerError() *CmsPageRepositoryV1GetByIDGetInternalServerError {
	return &CmsPageRepositoryV1GetByIDGetInternalServerError{}
}

/* CmsPageRepositoryV1GetByIDGetInternalServerError describes a response with status code 500, with default header values.

Internal Server error
*/
type CmsPageRepositoryV1GetByIDGetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CmsPageRepositoryV1GetByIDGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /V1/cmsPage/{pageId}][%d] cmsPageRepositoryV1GetByIdGetInternalServerError  %+v", 500, o.Payload)
}
func (o *CmsPageRepositoryV1GetByIDGetInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CmsPageRepositoryV1GetByIDGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCmsPageRepositoryV1GetByIDGetDefault creates a CmsPageRepositoryV1GetByIDGetDefault with default headers values
func NewCmsPageRepositoryV1GetByIDGetDefault(code int) *CmsPageRepositoryV1GetByIDGetDefault {
	return &CmsPageRepositoryV1GetByIDGetDefault{
		_statusCode: code,
	}
}

/* CmsPageRepositoryV1GetByIDGetDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type CmsPageRepositoryV1GetByIDGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cms page repository v1 get by Id get default response
func (o *CmsPageRepositoryV1GetByIDGetDefault) Code() int {
	return o._statusCode
}

func (o *CmsPageRepositoryV1GetByIDGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/cmsPage/{pageId}][%d] cmsPageRepositoryV1GetByIdGet default  %+v", o._statusCode, o.Payload)
}
func (o *CmsPageRepositoryV1GetByIDGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CmsPageRepositoryV1GetByIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
