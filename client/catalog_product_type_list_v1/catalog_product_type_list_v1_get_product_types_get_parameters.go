// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_type_list_v1

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

// NewCatalogProductTypeListV1GetProductTypesGetParams creates a new CatalogProductTypeListV1GetProductTypesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogProductTypeListV1GetProductTypesGetParams() *CatalogProductTypeListV1GetProductTypesGetParams {
	return &CatalogProductTypeListV1GetProductTypesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogProductTypeListV1GetProductTypesGetParamsWithTimeout creates a new CatalogProductTypeListV1GetProductTypesGetParams object
// with the ability to set a timeout on a request.
func NewCatalogProductTypeListV1GetProductTypesGetParamsWithTimeout(timeout time.Duration) *CatalogProductTypeListV1GetProductTypesGetParams {
	return &CatalogProductTypeListV1GetProductTypesGetParams{
		timeout: timeout,
	}
}

// NewCatalogProductTypeListV1GetProductTypesGetParamsWithContext creates a new CatalogProductTypeListV1GetProductTypesGetParams object
// with the ability to set a context for a request.
func NewCatalogProductTypeListV1GetProductTypesGetParamsWithContext(ctx context.Context) *CatalogProductTypeListV1GetProductTypesGetParams {
	return &CatalogProductTypeListV1GetProductTypesGetParams{
		Context: ctx,
	}
}

// NewCatalogProductTypeListV1GetProductTypesGetParamsWithHTTPClient creates a new CatalogProductTypeListV1GetProductTypesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogProductTypeListV1GetProductTypesGetParamsWithHTTPClient(client *http.Client) *CatalogProductTypeListV1GetProductTypesGetParams {
	return &CatalogProductTypeListV1GetProductTypesGetParams{
		HTTPClient: client,
	}
}

/* CatalogProductTypeListV1GetProductTypesGetParams contains all the parameters to send to the API endpoint
   for the catalog product type list v1 get product types get operation.

   Typically these are written to a http.Request.
*/
type CatalogProductTypeListV1GetProductTypesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the catalog product type list v1 get product types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogProductTypeListV1GetProductTypesGetParams) WithDefaults() *CatalogProductTypeListV1GetProductTypesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog product type list v1 get product types get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogProductTypeListV1GetProductTypesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog product type list v1 get product types get params
func (o *CatalogProductTypeListV1GetProductTypesGetParams) WithTimeout(timeout time.Duration) *CatalogProductTypeListV1GetProductTypesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog product type list v1 get product types get params
func (o *CatalogProductTypeListV1GetProductTypesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog product type list v1 get product types get params
func (o *CatalogProductTypeListV1GetProductTypesGetParams) WithContext(ctx context.Context) *CatalogProductTypeListV1GetProductTypesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog product type list v1 get product types get params
func (o *CatalogProductTypeListV1GetProductTypesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog product type list v1 get product types get params
func (o *CatalogProductTypeListV1GetProductTypesGetParams) WithHTTPClient(client *http.Client) *CatalogProductTypeListV1GetProductTypesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog product type list v1 get product types get params
func (o *CatalogProductTypeListV1GetProductTypesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogProductTypeListV1GetProductTypesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
