// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_cart_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new quote guest cart repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for quote guest cart repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	QuoteGuestCartRepositoryV1GetGet(params *QuoteGuestCartRepositoryV1GetGetParams, opts ...ClientOption) (*QuoteGuestCartRepositoryV1GetGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  QuoteGuestCartRepositoryV1GetGet Enable a guest user to return information for a specified cart.
*/
func (a *Client) QuoteGuestCartRepositoryV1GetGet(params *QuoteGuestCartRepositoryV1GetGetParams, opts ...ClientOption) (*QuoteGuestCartRepositoryV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteGuestCartRepositoryV1GetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "quoteGuestCartRepositoryV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/guest-carts/{cartId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuoteGuestCartRepositoryV1GetGetReader{formats: a.formats},
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
	success, ok := result.(*QuoteGuestCartRepositoryV1GetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QuoteGuestCartRepositoryV1GetGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}