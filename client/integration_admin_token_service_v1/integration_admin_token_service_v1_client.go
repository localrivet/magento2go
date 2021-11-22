// Code generated by go-swagger; DO NOT EDIT.

package integration_admin_token_service_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new integration admin token service v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for integration admin token service v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	IntegrationAdminTokenServiceV1CreateAdminAccessTokenPost(params *IntegrationAdminTokenServiceV1CreateAdminAccessTokenPostParams, opts ...ClientOption) (*IntegrationAdminTokenServiceV1CreateAdminAccessTokenPostOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  IntegrationAdminTokenServiceV1CreateAdminAccessTokenPost Create access token for admin given the admin credentials.
*/
func (a *Client) IntegrationAdminTokenServiceV1CreateAdminAccessTokenPost(params *IntegrationAdminTokenServiceV1CreateAdminAccessTokenPostParams, opts ...ClientOption) (*IntegrationAdminTokenServiceV1CreateAdminAccessTokenPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIntegrationAdminTokenServiceV1CreateAdminAccessTokenPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "integrationAdminTokenServiceV1CreateAdminAccessTokenPost",
		Method:             "POST",
		PathPattern:        "/V1/integration/admin/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IntegrationAdminTokenServiceV1CreateAdminAccessTokenPostReader{formats: a.formats},
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
	success, ok := result.(*IntegrationAdminTokenServiceV1CreateAdminAccessTokenPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IntegrationAdminTokenServiceV1CreateAdminAccessTokenPostDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}