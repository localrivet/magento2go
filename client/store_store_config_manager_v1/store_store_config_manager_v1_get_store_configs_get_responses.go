// Code generated by go-swagger; DO NOT EDIT.

package store_store_config_manager_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/localrivet/magento2go/models"
)

// StoreStoreConfigManagerV1GetStoreConfigsGetReader is a Reader for the StoreStoreConfigManagerV1GetStoreConfigsGet structure.
type StoreStoreConfigManagerV1GetStoreConfigsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StoreStoreConfigManagerV1GetStoreConfigsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStoreStoreConfigManagerV1GetStoreConfigsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStoreStoreConfigManagerV1GetStoreConfigsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStoreStoreConfigManagerV1GetStoreConfigsGetOK creates a StoreStoreConfigManagerV1GetStoreConfigsGetOK with default headers values
func NewStoreStoreConfigManagerV1GetStoreConfigsGetOK() *StoreStoreConfigManagerV1GetStoreConfigsGetOK {
	return &StoreStoreConfigManagerV1GetStoreConfigsGetOK{}
}

/* StoreStoreConfigManagerV1GetStoreConfigsGetOK describes a response with status code 200, with default header values.

200 Success.
*/
type StoreStoreConfigManagerV1GetStoreConfigsGetOK struct {
	Payload []*models.StoreDataStoreConfigInterface
}

func (o *StoreStoreConfigManagerV1GetStoreConfigsGetOK) Error() string {
	return fmt.Sprintf("[GET /V1/store/storeConfigs][%d] storeStoreConfigManagerV1GetStoreConfigsGetOK  %+v", 200, o.Payload)
}
func (o *StoreStoreConfigManagerV1GetStoreConfigsGetOK) GetPayload() []*models.StoreDataStoreConfigInterface {
	return o.Payload
}

func (o *StoreStoreConfigManagerV1GetStoreConfigsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStoreStoreConfigManagerV1GetStoreConfigsGetDefault creates a StoreStoreConfigManagerV1GetStoreConfigsGetDefault with default headers values
func NewStoreStoreConfigManagerV1GetStoreConfigsGetDefault(code int) *StoreStoreConfigManagerV1GetStoreConfigsGetDefault {
	return &StoreStoreConfigManagerV1GetStoreConfigsGetDefault{
		_statusCode: code,
	}
}

/* StoreStoreConfigManagerV1GetStoreConfigsGetDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type StoreStoreConfigManagerV1GetStoreConfigsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the store store config manager v1 get store configs get default response
func (o *StoreStoreConfigManagerV1GetStoreConfigsGetDefault) Code() int {
	return o._statusCode
}

func (o *StoreStoreConfigManagerV1GetStoreConfigsGetDefault) Error() string {
	return fmt.Sprintf("[GET /V1/store/storeConfigs][%d] storeStoreConfigManagerV1GetStoreConfigsGet default  %+v", o._statusCode, o.Payload)
}
func (o *StoreStoreConfigManagerV1GetStoreConfigsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *StoreStoreConfigManagerV1GetStoreConfigsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
