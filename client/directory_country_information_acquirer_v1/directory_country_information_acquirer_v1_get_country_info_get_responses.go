// Code generated by go-swagger; DO NOT EDIT.

package directory_country_information_acquirer_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"magento2go/models"
)

// DirectoryCountryInformationAcquirerV1GetCountryInfoGetReader is a Reader for the DirectoryCountryInformationAcquirerV1GetCountryInfoGet structure.
type DirectoryCountryInformationAcquirerV1GetCountryInfoGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetOK creates a DirectoryCountryInformationAcquirerV1GetCountryInfoGetOK with default headers values
func NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetOK() *DirectoryCountryInformationAcquirerV1GetCountryInfoGetOK {
	return &DirectoryCountryInformationAcquirerV1GetCountryInfoGetOK{}
}

/* DirectoryCountryInformationAcquirerV1GetCountryInfoGetOK describes a response with status code 200, with default header values.

200 Success.
*/
type DirectoryCountryInformationAcquirerV1GetCountryInfoGetOK struct {
	Payload *models.DirectoryDataCountryInformationInterface
}

func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/directory/countries/{countryId}][%d] directoryCountryInformationAcquirerV1GetCountryInfoGetOK  %+v", 200, o.Payload)
}
func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetOK) GetPayload() *models.DirectoryDataCountryInformationInterface {
	return o.Payload
}

func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectoryDataCountryInformationInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetBadRequest creates a DirectoryCountryInformationAcquirerV1GetCountryInfoGetBadRequest with default headers values
func NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetBadRequest() *DirectoryCountryInformationAcquirerV1GetCountryInfoGetBadRequest {
	return &DirectoryCountryInformationAcquirerV1GetCountryInfoGetBadRequest{}
}

/* DirectoryCountryInformationAcquirerV1GetCountryInfoGetBadRequest describes a response with status code 400, with default header values.

400 Bad Request
*/
type DirectoryCountryInformationAcquirerV1GetCountryInfoGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /V1/directory/countries/{countryId}][%d] directoryCountryInformationAcquirerV1GetCountryInfoGetBadRequest  %+v", 400, o.Payload)
}
func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetDefault creates a DirectoryCountryInformationAcquirerV1GetCountryInfoGetDefault with default headers values
func NewDirectoryCountryInformationAcquirerV1GetCountryInfoGetDefault(code int) *DirectoryCountryInformationAcquirerV1GetCountryInfoGetDefault {
	return &DirectoryCountryInformationAcquirerV1GetCountryInfoGetDefault{
		_statusCode: code,
	}
}

/* DirectoryCountryInformationAcquirerV1GetCountryInfoGetDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type DirectoryCountryInformationAcquirerV1GetCountryInfoGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the directory country information acquirer v1 get country info get default response
func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetDefault) Code() int {
	return o._statusCode
}

func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/directory/countries/{countryId}][%d] directoryCountryInformationAcquirerV1GetCountryInfoGet default  %+v", o._statusCode, o.Payload)
}
func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DirectoryCountryInformationAcquirerV1GetCountryInfoGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
