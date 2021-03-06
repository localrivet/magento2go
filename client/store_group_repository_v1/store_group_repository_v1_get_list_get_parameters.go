// Code generated by go-swagger; DO NOT EDIT.

package store_group_repository_v1

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

// NewStoreGroupRepositoryV1GetListGetParams creates a new StoreGroupRepositoryV1GetListGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStoreGroupRepositoryV1GetListGetParams() *StoreGroupRepositoryV1GetListGetParams {
	return &StoreGroupRepositoryV1GetListGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStoreGroupRepositoryV1GetListGetParamsWithTimeout creates a new StoreGroupRepositoryV1GetListGetParams object
// with the ability to set a timeout on a request.
func NewStoreGroupRepositoryV1GetListGetParamsWithTimeout(timeout time.Duration) *StoreGroupRepositoryV1GetListGetParams {
	return &StoreGroupRepositoryV1GetListGetParams{
		timeout: timeout,
	}
}

// NewStoreGroupRepositoryV1GetListGetParamsWithContext creates a new StoreGroupRepositoryV1GetListGetParams object
// with the ability to set a context for a request.
func NewStoreGroupRepositoryV1GetListGetParamsWithContext(ctx context.Context) *StoreGroupRepositoryV1GetListGetParams {
	return &StoreGroupRepositoryV1GetListGetParams{
		Context: ctx,
	}
}

// NewStoreGroupRepositoryV1GetListGetParamsWithHTTPClient creates a new StoreGroupRepositoryV1GetListGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewStoreGroupRepositoryV1GetListGetParamsWithHTTPClient(client *http.Client) *StoreGroupRepositoryV1GetListGetParams {
	return &StoreGroupRepositoryV1GetListGetParams{
		HTTPClient: client,
	}
}

/* StoreGroupRepositoryV1GetListGetParams contains all the parameters to send to the API endpoint
   for the store group repository v1 get list get operation.

   Typically these are written to a http.Request.
*/
type StoreGroupRepositoryV1GetListGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the store group repository v1 get list get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StoreGroupRepositoryV1GetListGetParams) WithDefaults() *StoreGroupRepositoryV1GetListGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the store group repository v1 get list get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StoreGroupRepositoryV1GetListGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the store group repository v1 get list get params
func (o *StoreGroupRepositoryV1GetListGetParams) WithTimeout(timeout time.Duration) *StoreGroupRepositoryV1GetListGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the store group repository v1 get list get params
func (o *StoreGroupRepositoryV1GetListGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the store group repository v1 get list get params
func (o *StoreGroupRepositoryV1GetListGetParams) WithContext(ctx context.Context) *StoreGroupRepositoryV1GetListGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the store group repository v1 get list get params
func (o *StoreGroupRepositoryV1GetListGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the store group repository v1 get list get params
func (o *StoreGroupRepositoryV1GetListGetParams) WithHTTPClient(client *http.Client) *StoreGroupRepositoryV1GetListGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the store group repository v1 get list get params
func (o *StoreGroupRepositoryV1GetListGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StoreGroupRepositoryV1GetListGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
