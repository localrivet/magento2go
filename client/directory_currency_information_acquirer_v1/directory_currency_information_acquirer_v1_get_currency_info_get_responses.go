// Code generated by go-swagger; DO NOT EDIT.

package directory_currency_information_acquirer_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/localrivet/magento2go/models"
)

// DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetReader is a Reader for the DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGet structure.
type DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetOK creates a DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetOK with default headers values
func NewDirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetOK() *DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetOK {
	return &DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetOK{}
}

/* DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetOK describes a response with status code 200, with default header values.

200 Success.
*/
type DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetOK struct {
	Payload *models.DirectoryDataCurrencyInformationInterface
}

func (o *DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/directory/currency][%d] directoryCurrencyInformationAcquirerV1GetCurrencyInfoGetOK  %+v", 200, o.Payload)
}
func (o *DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetOK) GetPayload() *models.DirectoryDataCurrencyInformationInterface {
	return o.Payload
}

func (o *DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectoryDataCurrencyInformationInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetDefault creates a DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetDefault with default headers values
func NewDirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetDefault(code int) *DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetDefault {
	return &DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetDefault{
		_statusCode: code,
	}
}

/* DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the directory currency information acquirer v1 get currency info get default response
func (o *DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetDefault) Code() int {
	return o._statusCode
}

func (o *DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/directory/currency][%d] directoryCurrencyInformationAcquirerV1GetCurrencyInfoGet default  %+v", o._statusCode, o.Payload)
}
func (o *DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
