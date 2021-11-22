// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_payment_method_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new quote guest payment method management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for quote guest payment method management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	QuoteGuestPaymentMethodManagementV1GetGet(params *QuoteGuestPaymentMethodManagementV1GetGetParams, opts ...ClientOption) (*QuoteGuestPaymentMethodManagementV1GetGetOK, error)

	QuoteGuestPaymentMethodManagementV1GetListGet(params *QuoteGuestPaymentMethodManagementV1GetListGetParams, opts ...ClientOption) (*QuoteGuestPaymentMethodManagementV1GetListGetOK, error)

	QuoteGuestPaymentMethodManagementV1SetPut(params *QuoteGuestPaymentMethodManagementV1SetPutParams, opts ...ClientOption) (*QuoteGuestPaymentMethodManagementV1SetPutOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  QuoteGuestPaymentMethodManagementV1GetGet Return the payment method for a specified shopping cart.
*/
func (a *Client) QuoteGuestPaymentMethodManagementV1GetGet(params *QuoteGuestPaymentMethodManagementV1GetGetParams, opts ...ClientOption) (*QuoteGuestPaymentMethodManagementV1GetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteGuestPaymentMethodManagementV1GetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "quoteGuestPaymentMethodManagementV1GetGet",
		Method:             "GET",
		PathPattern:        "/V1/guest-carts/{cartId}/selected-payment-method",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuoteGuestPaymentMethodManagementV1GetGetReader{formats: a.formats},
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
	success, ok := result.(*QuoteGuestPaymentMethodManagementV1GetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QuoteGuestPaymentMethodManagementV1GetGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QuoteGuestPaymentMethodManagementV1GetListGet List available payment methods for a specified shopping cart. This call returns an array of objects, but detailed information about each object’s attributes might not be included.  See https://devdocs.magento.com/codelinks/attributes.html#GuestPaymentMethodManagementInterface to determine which call to use to get detailed information about all attributes for an object.
*/
func (a *Client) QuoteGuestPaymentMethodManagementV1GetListGet(params *QuoteGuestPaymentMethodManagementV1GetListGetParams, opts ...ClientOption) (*QuoteGuestPaymentMethodManagementV1GetListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteGuestPaymentMethodManagementV1GetListGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "quoteGuestPaymentMethodManagementV1GetListGet",
		Method:             "GET",
		PathPattern:        "/V1/guest-carts/{cartId}/payment-methods",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuoteGuestPaymentMethodManagementV1GetListGetReader{formats: a.formats},
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
	success, ok := result.(*QuoteGuestPaymentMethodManagementV1GetListGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QuoteGuestPaymentMethodManagementV1GetListGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QuoteGuestPaymentMethodManagementV1SetPut Add a specified payment method to a specified shopping cart.
*/
func (a *Client) QuoteGuestPaymentMethodManagementV1SetPut(params *QuoteGuestPaymentMethodManagementV1SetPutParams, opts ...ClientOption) (*QuoteGuestPaymentMethodManagementV1SetPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteGuestPaymentMethodManagementV1SetPutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "quoteGuestPaymentMethodManagementV1SetPut",
		Method:             "PUT",
		PathPattern:        "/V1/guest-carts/{cartId}/selected-payment-method",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuoteGuestPaymentMethodManagementV1SetPutReader{formats: a.formats},
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
	success, ok := result.(*QuoteGuestPaymentMethodManagementV1SetPutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QuoteGuestPaymentMethodManagementV1SetPutDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}