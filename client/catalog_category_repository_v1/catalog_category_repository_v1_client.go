// Code generated by go-swagger; DO NOT EDIT.

package catalog_category_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new catalog category repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog category repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CatalogCategoryRepositoryV1GetGet(params *CatalogCategoryRepositoryV1GetGetParams, opts ...ClientOption) (*CatalogCategoryRepositoryV1GetGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CatalogCategoryRepositoryV1GetGet Get info about category by category id
*/
func (a *Client) CatalogCategoryRepositoryV1GetGet(params *CatalogCategoryRepositoryV1GetGetParams, opts ...ClientOption) (*CatalogCategoryRepositoryV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogCategoryRepositoryV1GetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "catalogCategoryRepositoryV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/categories/{categoryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogCategoryRepositoryV1GetGetReader{formats: a.formats},
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
	success, ok := result.(*CatalogCategoryRepositoryV1GetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CatalogCategoryRepositoryV1GetGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
