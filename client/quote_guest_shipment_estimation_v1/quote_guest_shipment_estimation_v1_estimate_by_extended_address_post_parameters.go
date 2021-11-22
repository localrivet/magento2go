// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_shipment_estimation_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewQuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams creates a new QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams() *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams {
	return &QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParamsWithTimeout creates a new QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams object
// with the ability to set a timeout on a request.
func NewQuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParamsWithTimeout(timeout time.Duration) *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams {
	return &QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams{
		timeout: timeout,
	}
}

// NewQuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParamsWithContext creates a new QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams object
// with the ability to set a context for a request.
func NewQuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParamsWithContext(ctx context.Context) *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams {
	return &QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams{
		Context: ctx,
	}
}

// NewQuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParamsWithHTTPClient creates a new QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewQuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParamsWithHTTPClient(client *http.Client) *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams {
	return &QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams{
		HTTPClient: client,
	}
}

/* QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams contains all the parameters to send to the API endpoint
   for the quote guest shipment estimation v1 estimate by extended address post operation.

   Typically these are written to a http.Request.
*/
type QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams struct {

	// CartID.
	CartID string

	// QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody.
	QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the quote guest shipment estimation v1 estimate by extended address post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams) WithDefaults() *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the quote guest shipment estimation v1 estimate by extended address post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the quote guest shipment estimation v1 estimate by extended address post params
func (o *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams) WithTimeout(timeout time.Duration) *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote guest shipment estimation v1 estimate by extended address post params
func (o *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote guest shipment estimation v1 estimate by extended address post params
func (o *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams) WithContext(ctx context.Context) *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote guest shipment estimation v1 estimate by extended address post params
func (o *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote guest shipment estimation v1 estimate by extended address post params
func (o *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams) WithHTTPClient(client *http.Client) *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote guest shipment estimation v1 estimate by extended address post params
func (o *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the quote guest shipment estimation v1 estimate by extended address post params
func (o *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams) WithCartID(cartID string) *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the quote guest shipment estimation v1 estimate by extended address post params
func (o *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithQuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody adds the quoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody to the quote guest shipment estimation v1 estimate by extended address post params
func (o *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams) WithQuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody(quoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody) *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams {
	o.SetQuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody(quoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody)
	return o
}

// SetQuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody adds the quoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody to the quote guest shipment estimation v1 estimate by extended address post params
func (o *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams) SetQuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody(quoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody) {
	o.QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody = quoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.QuoteGuestShipmentEstimationV1EstimateByExtendedAddressPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
