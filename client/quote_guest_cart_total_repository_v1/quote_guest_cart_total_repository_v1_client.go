// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_cart_total_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new quote guest cart total repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for quote guest cart total repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	QuoteGuestCartTotalRepositoryV1GetGet(params *QuoteGuestCartTotalRepositoryV1GetGetParams, opts ...ClientOption) (*QuoteGuestCartTotalRepositoryV1GetGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  QuoteGuestCartTotalRepositoryV1GetGet Return quote totals data for a specified cart.
*/
func (a *Client) QuoteGuestCartTotalRepositoryV1GetGet(params *QuoteGuestCartTotalRepositoryV1GetGetParams, opts ...ClientOption) (*QuoteGuestCartTotalRepositoryV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteGuestCartTotalRepositoryV1GetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "quoteGuestCartTotalRepositoryV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/guest-carts/{cartId}/totals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuoteGuestCartTotalRepositoryV1GetGetReader{formats: a.formats},
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
	success, ok := result.(*QuoteGuestCartTotalRepositoryV1GetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QuoteGuestCartTotalRepositoryV1GetGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
