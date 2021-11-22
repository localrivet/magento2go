// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_shipment_estimation_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new quote guest shipment estimation v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for quote guest shipment estimation v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPost(params *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams, opts ...ClientOption) (*QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPost Estimate shipping by address and return list of available shipping methods
*/
func (a *Client) QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPost(params *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams, opts ...ClientOption) (*QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "quoteGuestShipmentEstimationV1EstimateByExtendedAddressPost",
		Method:             "POST",
		PathPattern:        "/V1/guest-carts/{cartId}/estimate-shipping-methods",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostReader{formats: a.formats},
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
	success, ok := result.(*QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
