// Code generated by go-swagger; DO NOT EDIT.

package configurable_product_link_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new configurable product link management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for configurable product link management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ConfigurableProductLinkManagementV1GetChildrenGet(params *ConfigurableProductLinkManagementV1GetChildrenGetParams, opts ...ClientOption) (*ConfigurableProductLinkManagementV1GetChildrenGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ConfigurableProductLinkManagementV1GetChildrenGet Get all children for Configurable product
*/
func (a *Client) ConfigurableProductLinkManagementV1GetChildrenGet(params *ConfigurableProductLinkManagementV1GetChildrenGetParams, opts ...ClientOption) (*ConfigurableProductLinkManagementV1GetChildrenGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfigurableProductLinkManagementV1GetChildrenGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "configurableProductLinkManagementV1GetChildrenGet",
		Method:             "GET",
		PathPattern:        "/V1/configurable-products/{sku}/children",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConfigurableProductLinkManagementV1GetChildrenGetReader{formats: a.formats},
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
	success, ok := result.(*ConfigurableProductLinkManagementV1GetChildrenGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ConfigurableProductLinkManagementV1GetChildrenGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
