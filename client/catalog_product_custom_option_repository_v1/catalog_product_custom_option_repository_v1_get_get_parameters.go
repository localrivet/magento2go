// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_custom_option_repository_v1

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
	"github.com/go-openapi/swag"
)

// NewCatalogProductCustomOptionRepositoryV1GetGetParams creates a new CatalogProductCustomOptionRepositoryV1GetGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogProductCustomOptionRepositoryV1GetGetParams() *CatalogProductCustomOptionRepositoryV1GetGetParams {
	return &CatalogProductCustomOptionRepositoryV1GetGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogProductCustomOptionRepositoryV1GetGetParamsWithTimeout creates a new CatalogProductCustomOptionRepositoryV1GetGetParams object
// with the ability to set a timeout on a request.
func NewCatalogProductCustomOptionRepositoryV1GetGetParamsWithTimeout(timeout time.Duration) *CatalogProductCustomOptionRepositoryV1GetGetParams {
	return &CatalogProductCustomOptionRepositoryV1GetGetParams{
		timeout: timeout,
	}
}

// NewCatalogProductCustomOptionRepositoryV1GetGetParamsWithContext creates a new CatalogProductCustomOptionRepositoryV1GetGetParams object
// with the ability to set a context for a request.
func NewCatalogProductCustomOptionRepositoryV1GetGetParamsWithContext(ctx context.Context) *CatalogProductCustomOptionRepositoryV1GetGetParams {
	return &CatalogProductCustomOptionRepositoryV1GetGetParams{
		Context: ctx,
	}
}

// NewCatalogProductCustomOptionRepositoryV1GetGetParamsWithHTTPClient creates a new CatalogProductCustomOptionRepositoryV1GetGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogProductCustomOptionRepositoryV1GetGetParamsWithHTTPClient(client *http.Client) *CatalogProductCustomOptionRepositoryV1GetGetParams {
	return &CatalogProductCustomOptionRepositoryV1GetGetParams{
		HTTPClient: client,
	}
}

/* CatalogProductCustomOptionRepositoryV1GetGetParams contains all the parameters to send to the API endpoint
   for the catalog product custom option repository v1 get get operation.

   Typically these are written to a http.Request.
*/
type CatalogProductCustomOptionRepositoryV1GetGetParams struct {

	// OptionID.
	OptionID int64

	// Sku.
	Sku string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the catalog product custom option repository v1 get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogProductCustomOptionRepositoryV1GetGetParams) WithDefaults() *CatalogProductCustomOptionRepositoryV1GetGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog product custom option repository v1 get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogProductCustomOptionRepositoryV1GetGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog product custom option repository v1 get get params
func (o *CatalogProductCustomOptionRepositoryV1GetGetParams) WithTimeout(timeout time.Duration) *CatalogProductCustomOptionRepositoryV1GetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog product custom option repository v1 get get params
func (o *CatalogProductCustomOptionRepositoryV1GetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog product custom option repository v1 get get params
func (o *CatalogProductCustomOptionRepositoryV1GetGetParams) WithContext(ctx context.Context) *CatalogProductCustomOptionRepositoryV1GetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog product custom option repository v1 get get params
func (o *CatalogProductCustomOptionRepositoryV1GetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog product custom option repository v1 get get params
func (o *CatalogProductCustomOptionRepositoryV1GetGetParams) WithHTTPClient(client *http.Client) *CatalogProductCustomOptionRepositoryV1GetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog product custom option repository v1 get get params
func (o *CatalogProductCustomOptionRepositoryV1GetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOptionID adds the optionID to the catalog product custom option repository v1 get get params
func (o *CatalogProductCustomOptionRepositoryV1GetGetParams) WithOptionID(optionID int64) *CatalogProductCustomOptionRepositoryV1GetGetParams {
	o.SetOptionID(optionID)
	return o
}

// SetOptionID adds the optionId to the catalog product custom option repository v1 get get params
func (o *CatalogProductCustomOptionRepositoryV1GetGetParams) SetOptionID(optionID int64) {
	o.OptionID = optionID
}

// WithSku adds the sku to the catalog product custom option repository v1 get get params
func (o *CatalogProductCustomOptionRepositoryV1GetGetParams) WithSku(sku string) *CatalogProductCustomOptionRepositoryV1GetGetParams {
	o.SetSku(sku)
	return o
}

// SetSku adds the sku to the catalog product custom option repository v1 get get params
func (o *CatalogProductCustomOptionRepositoryV1GetGetParams) SetSku(sku string) {
	o.Sku = sku
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogProductCustomOptionRepositoryV1GetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param optionId
	if err := r.SetPathParam("optionId", swag.FormatInt64(o.OptionID)); err != nil {
		return err
	}

	// path param sku
	if err := r.SetPathParam("sku", o.Sku); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
