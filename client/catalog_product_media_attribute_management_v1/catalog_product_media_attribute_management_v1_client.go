// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_media_attribute_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new catalog product media attribute management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog product media attribute management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CatalogProductMediaAttributeManagementV1GetListGet(params *CatalogProductMediaAttributeManagementV1GetListGetParams, opts ...ClientOption) (*CatalogProductMediaAttributeManagementV1GetListGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CatalogProductMediaAttributeManagementV1GetListGet Retrieve the list of media attributes (fronted input type is media_image) assigned to the given attribute set.
*/
func (a *Client) CatalogProductMediaAttributeManagementV1GetListGet(params *CatalogProductMediaAttributeManagementV1GetListGetParams, opts ...ClientOption) (*CatalogProductMediaAttributeManagementV1GetListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogProductMediaAttributeManagementV1GetListGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "catalogProductMediaAttributeManagementV1GetListGet",
		Method:             "GET",
		PathPattern:        "/V1/products/media/types/{attributeSetName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogProductMediaAttributeManagementV1GetListGetReader{formats: a.formats},
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
	success, ok := result.(*CatalogProductMediaAttributeManagementV1GetListGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CatalogProductMediaAttributeManagementV1GetListGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
