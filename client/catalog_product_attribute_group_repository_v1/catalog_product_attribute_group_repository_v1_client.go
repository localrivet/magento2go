// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_attribute_group_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new catalog product attribute group repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog product attribute group repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CatalogProductAttributeGroupRepositoryV1GetListGet(params *CatalogProductAttributeGroupRepositoryV1GetListGetParams, opts ...ClientOption) (*CatalogProductAttributeGroupRepositoryV1GetListGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CatalogProductAttributeGroupRepositoryV1GetListGet Retrieve list of attribute groups
*/
func (a *Client) CatalogProductAttributeGroupRepositoryV1GetListGet(params *CatalogProductAttributeGroupRepositoryV1GetListGetParams, opts ...ClientOption) (*CatalogProductAttributeGroupRepositoryV1GetListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeGroupRepositoryV1GetListGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "catalogProductAttributeGroupRepositoryV1GetListGet",
		Method:             "GET",
		PathPattern:        "/V1/products/attribute-sets/groups/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogProductAttributeGroupRepositoryV1GetListGetReader{formats: a.formats},
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
	success, ok := result.(*CatalogProductAttributeGroupRepositoryV1GetListGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CatalogProductAttributeGroupRepositoryV1GetListGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
