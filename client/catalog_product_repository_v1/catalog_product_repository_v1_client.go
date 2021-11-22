// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new catalog product repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog product repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CatalogProductRepositoryV1GetGet(params *CatalogProductRepositoryV1GetGetParams, opts ...ClientOption) (*CatalogProductRepositoryV1GetGetOK, error)

	CatalogProductRepositoryV1GetListGet(params *CatalogProductRepositoryV1GetListGetParams, opts ...ClientOption) (*CatalogProductRepositoryV1GetListGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CatalogProductRepositoryV1GetGet Get info about product by product SKU
*/
func (a *Client) CatalogProductRepositoryV1GetGet(params *CatalogProductRepositoryV1GetGetParams, opts ...ClientOption) (*CatalogProductRepositoryV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductRepositoryV1GetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "catalogProductRepositoryV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/products/{sku}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogProductRepositoryV1GetGetReader{formats: a.formats},
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
	success, ok := result.(*CatalogProductRepositoryV1GetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CatalogProductRepositoryV1GetGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CatalogProductRepositoryV1GetListGet Get product list
*/
func (a *Client) CatalogProductRepositoryV1GetListGet(params *CatalogProductRepositoryV1GetListGetParams, opts ...ClientOption) (*CatalogProductRepositoryV1GetListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductRepositoryV1GetListGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "catalogProductRepositoryV1GetListGet",
		Method:             "GET",
		PathPattern:        "/V1/products",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogProductRepositoryV1GetListGetReader{formats: a.formats},
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
	success, ok := result.(*CatalogProductRepositoryV1GetListGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CatalogProductRepositoryV1GetListGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}