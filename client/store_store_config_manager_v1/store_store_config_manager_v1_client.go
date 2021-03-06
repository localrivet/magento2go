// Code generated by go-swagger; DO NOT EDIT.

package store_store_config_manager_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new store store config manager v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for store store config manager v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	StoreStoreConfigManagerV1GetStoreConfigsGet(params *StoreStoreConfigManagerV1GetStoreConfigsGetParams, opts ...ClientOption) (*StoreStoreConfigManagerV1GetStoreConfigsGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  StoreStoreConfigManagerV1GetStoreConfigsGet store store config manager v1 get store configs get API
*/
func (a *Client) StoreStoreConfigManagerV1GetStoreConfigsGet(params *StoreStoreConfigManagerV1GetStoreConfigsGetParams, opts ...ClientOption) (*StoreStoreConfigManagerV1GetStoreConfigsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStoreStoreConfigManagerV1GetStoreConfigsGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "storeStoreConfigManagerV1GetStoreConfigsGet",
		Method:             "GET",
		PathPattern:        "/V1/store/storeConfigs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StoreStoreConfigManagerV1GetStoreConfigsGetReader{formats: a.formats},
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
	success, ok := result.(*StoreStoreConfigManagerV1GetStoreConfigsGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StoreStoreConfigManagerV1GetStoreConfigsGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
