// Code generated by go-swagger; DO NOT EDIT.

package directory_currency_information_acquirer_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new directory currency information acquirer v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for directory currency information acquirer v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGet(params *DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetParams, opts ...ClientOption) (*DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGet Get currency information for the store.
*/
func (a *Client) DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGet(params *DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetParams, opts ...ClientOption) (*DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "directoryCurrencyInformationAcquirerV1GetCurrencyInfoGet",
		Method:             "GET",
		PathPattern:        "/V1/directory/currency",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DirectoryCurrencyInformationAcquirerV1GetCurrencyInfoGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
