// Code generated by go-swagger; DO NOT EDIT.

package catalog_inventory_stock_registry_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new catalog inventory stock registry v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog inventory stock registry v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CatalogInventoryStockRegistryV1GetStockStatusBySkuGet(params *CatalogInventoryStockRegistryV1GetStockStatusBySkuGetParams, opts ...ClientOption) (*CatalogInventoryStockRegistryV1GetStockStatusBySkuGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CatalogInventoryStockRegistryV1GetStockStatusBySkuGet catalog inventory stock registry v1 get stock status by sku get API
*/
func (a *Client) CatalogInventoryStockRegistryV1GetStockStatusBySkuGet(params *CatalogInventoryStockRegistryV1GetStockStatusBySkuGetParams, opts ...ClientOption) (*CatalogInventoryStockRegistryV1GetStockStatusBySkuGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogInventoryStockRegistryV1GetStockStatusBySkuGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "catalogInventoryStockRegistryV1GetStockStatusBySkuGet",
		Method:             "GET",
		PathPattern:        "/V1/stockStatuses/{productSku}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogInventoryStockRegistryV1GetStockStatusBySkuGetReader{formats: a.formats},
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
	success, ok := result.(*CatalogInventoryStockRegistryV1GetStockStatusBySkuGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CatalogInventoryStockRegistryV1GetStockStatusBySkuGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
