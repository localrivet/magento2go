// Code generated by go-swagger; DO NOT EDIT.

package pay_pal_braintree_auth_v1

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

// NewPayPalBraintreeAuthV1GetGetParams creates a new PayPalBraintreeAuthV1GetGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPayPalBraintreeAuthV1GetGetParams() *PayPalBraintreeAuthV1GetGetParams {
	return &PayPalBraintreeAuthV1GetGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPayPalBraintreeAuthV1GetGetParamsWithTimeout creates a new PayPalBraintreeAuthV1GetGetParams object
// with the ability to set a timeout on a request.
func NewPayPalBraintreeAuthV1GetGetParamsWithTimeout(timeout time.Duration) *PayPalBraintreeAuthV1GetGetParams {
	return &PayPalBraintreeAuthV1GetGetParams{
		timeout: timeout,
	}
}

// NewPayPalBraintreeAuthV1GetGetParamsWithContext creates a new PayPalBraintreeAuthV1GetGetParams object
// with the ability to set a context for a request.
func NewPayPalBraintreeAuthV1GetGetParamsWithContext(ctx context.Context) *PayPalBraintreeAuthV1GetGetParams {
	return &PayPalBraintreeAuthV1GetGetParams{
		Context: ctx,
	}
}

// NewPayPalBraintreeAuthV1GetGetParamsWithHTTPClient creates a new PayPalBraintreeAuthV1GetGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPayPalBraintreeAuthV1GetGetParamsWithHTTPClient(client *http.Client) *PayPalBraintreeAuthV1GetGetParams {
	return &PayPalBraintreeAuthV1GetGetParams{
		HTTPClient: client,
	}
}

/* PayPalBraintreeAuthV1GetGetParams contains all the parameters to send to the API endpoint
   for the pay pal braintree auth v1 get get operation.

   Typically these are written to a http.Request.
*/
type PayPalBraintreeAuthV1GetGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pay pal braintree auth v1 get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PayPalBraintreeAuthV1GetGetParams) WithDefaults() *PayPalBraintreeAuthV1GetGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pay pal braintree auth v1 get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PayPalBraintreeAuthV1GetGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pay pal braintree auth v1 get get params
func (o *PayPalBraintreeAuthV1GetGetParams) WithTimeout(timeout time.Duration) *PayPalBraintreeAuthV1GetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pay pal braintree auth v1 get get params
func (o *PayPalBraintreeAuthV1GetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pay pal braintree auth v1 get get params
func (o *PayPalBraintreeAuthV1GetGetParams) WithContext(ctx context.Context) *PayPalBraintreeAuthV1GetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pay pal braintree auth v1 get get params
func (o *PayPalBraintreeAuthV1GetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pay pal braintree auth v1 get get params
func (o *PayPalBraintreeAuthV1GetGetParams) WithHTTPClient(client *http.Client) *PayPalBraintreeAuthV1GetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pay pal braintree auth v1 get get params
func (o *PayPalBraintreeAuthV1GetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PayPalBraintreeAuthV1GetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}