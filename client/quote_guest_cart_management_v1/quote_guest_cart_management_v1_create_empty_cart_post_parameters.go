// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_cart_management_v1

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

// NewQuoteGuestCartManagementV1CreateEmptyCartPostParams creates a new QuoteGuestCartManagementV1CreateEmptyCartPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQuoteGuestCartManagementV1CreateEmptyCartPostParams() *QuoteGuestCartManagementV1CreateEmptyCartPostParams {
	return &QuoteGuestCartManagementV1CreateEmptyCartPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteGuestCartManagementV1CreateEmptyCartPostParamsWithTimeout creates a new QuoteGuestCartManagementV1CreateEmptyCartPostParams object
// with the ability to set a timeout on a request.
func NewQuoteGuestCartManagementV1CreateEmptyCartPostParamsWithTimeout(timeout time.Duration) *QuoteGuestCartManagementV1CreateEmptyCartPostParams {
	return &QuoteGuestCartManagementV1CreateEmptyCartPostParams{
		timeout: timeout,
	}
}

// NewQuoteGuestCartManagementV1CreateEmptyCartPostParamsWithContext creates a new QuoteGuestCartManagementV1CreateEmptyCartPostParams object
// with the ability to set a context for a request.
func NewQuoteGuestCartManagementV1CreateEmptyCartPostParamsWithContext(ctx context.Context) *QuoteGuestCartManagementV1CreateEmptyCartPostParams {
	return &QuoteGuestCartManagementV1CreateEmptyCartPostParams{
		Context: ctx,
	}
}

// NewQuoteGuestCartManagementV1CreateEmptyCartPostParamsWithHTTPClient creates a new QuoteGuestCartManagementV1CreateEmptyCartPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewQuoteGuestCartManagementV1CreateEmptyCartPostParamsWithHTTPClient(client *http.Client) *QuoteGuestCartManagementV1CreateEmptyCartPostParams {
	return &QuoteGuestCartManagementV1CreateEmptyCartPostParams{
		HTTPClient: client,
	}
}

/* QuoteGuestCartManagementV1CreateEmptyCartPostParams contains all the parameters to send to the API endpoint
   for the quote guest cart management v1 create empty cart post operation.

   Typically these are written to a http.Request.
*/
type QuoteGuestCartManagementV1CreateEmptyCartPostParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the quote guest cart management v1 create empty cart post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuoteGuestCartManagementV1CreateEmptyCartPostParams) WithDefaults() *QuoteGuestCartManagementV1CreateEmptyCartPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the quote guest cart management v1 create empty cart post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuoteGuestCartManagementV1CreateEmptyCartPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the quote guest cart management v1 create empty cart post params
func (o *QuoteGuestCartManagementV1CreateEmptyCartPostParams) WithTimeout(timeout time.Duration) *QuoteGuestCartManagementV1CreateEmptyCartPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote guest cart management v1 create empty cart post params
func (o *QuoteGuestCartManagementV1CreateEmptyCartPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote guest cart management v1 create empty cart post params
func (o *QuoteGuestCartManagementV1CreateEmptyCartPostParams) WithContext(ctx context.Context) *QuoteGuestCartManagementV1CreateEmptyCartPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote guest cart management v1 create empty cart post params
func (o *QuoteGuestCartManagementV1CreateEmptyCartPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote guest cart management v1 create empty cart post params
func (o *QuoteGuestCartManagementV1CreateEmptyCartPostParams) WithHTTPClient(client *http.Client) *QuoteGuestCartManagementV1CreateEmptyCartPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote guest cart management v1 create empty cart post params
func (o *QuoteGuestCartManagementV1CreateEmptyCartPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteGuestCartManagementV1CreateEmptyCartPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
