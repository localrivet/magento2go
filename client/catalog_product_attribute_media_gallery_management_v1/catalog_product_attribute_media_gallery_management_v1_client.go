// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_attribute_media_gallery_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new catalog product attribute media gallery management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog product attribute media gallery management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CatalogProductAttributeMediaGalleryManagementV1GetGet(params *CatalogProductAttributeMediaGalleryManagementV1GetGetParams, opts ...ClientOption) (*CatalogProductAttributeMediaGalleryManagementV1GetGetOK, error)

	CatalogProductAttributeMediaGalleryManagementV1GetListGet(params *CatalogProductAttributeMediaGalleryManagementV1GetListGetParams, opts ...ClientOption) (*CatalogProductAttributeMediaGalleryManagementV1GetListGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CatalogProductAttributeMediaGalleryManagementV1GetGet Return information about gallery entry
*/
func (a *Client) CatalogProductAttributeMediaGalleryManagementV1GetGet(params *CatalogProductAttributeMediaGalleryManagementV1GetGetParams, opts ...ClientOption) (*CatalogProductAttributeMediaGalleryManagementV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeMediaGalleryManagementV1GetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "catalogProductAttributeMediaGalleryManagementV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/products/{sku}/media/{entryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogProductAttributeMediaGalleryManagementV1GetGetReader{formats: a.formats},
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
	success, ok := result.(*CatalogProductAttributeMediaGalleryManagementV1GetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CatalogProductAttributeMediaGalleryManagementV1GetGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CatalogProductAttributeMediaGalleryManagementV1GetListGet Retrieve the list of gallery entries associated with given product
*/
func (a *Client) CatalogProductAttributeMediaGalleryManagementV1GetListGet(params *CatalogProductAttributeMediaGalleryManagementV1GetListGetParams, opts ...ClientOption) (*CatalogProductAttributeMediaGalleryManagementV1GetListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductAttributeMediaGalleryManagementV1GetListGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "catalogProductAttributeMediaGalleryManagementV1GetListGet",
		Method:             "GET",
		PathPattern:        "/V1/products/{sku}/media",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogProductAttributeMediaGalleryManagementV1GetListGetReader{formats: a.formats},
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
	success, ok := result.(*CatalogProductAttributeMediaGalleryManagementV1GetListGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CatalogProductAttributeMediaGalleryManagementV1GetListGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}