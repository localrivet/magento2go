// Code generated by go-swagger; DO NOT EDIT.

package catalog_product_attribute_media_gallery_management_v1

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

// NewCatalogProductAttributeMediaGalleryManagementV1GetGetParams creates a new CatalogProductAttributeMediaGalleryManagementV1GetGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogProductAttributeMediaGalleryManagementV1GetGetParams() *CatalogProductAttributeMediaGalleryManagementV1GetGetParams {
	return &CatalogProductAttributeMediaGalleryManagementV1GetGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogProductAttributeMediaGalleryManagementV1GetGetParamsWithTimeout creates a new CatalogProductAttributeMediaGalleryManagementV1GetGetParams object
// with the ability to set a timeout on a request.
func NewCatalogProductAttributeMediaGalleryManagementV1GetGetParamsWithTimeout(timeout time.Duration) *CatalogProductAttributeMediaGalleryManagementV1GetGetParams {
	return &CatalogProductAttributeMediaGalleryManagementV1GetGetParams{
		timeout: timeout,
	}
}

// NewCatalogProductAttributeMediaGalleryManagementV1GetGetParamsWithContext creates a new CatalogProductAttributeMediaGalleryManagementV1GetGetParams object
// with the ability to set a context for a request.
func NewCatalogProductAttributeMediaGalleryManagementV1GetGetParamsWithContext(ctx context.Context) *CatalogProductAttributeMediaGalleryManagementV1GetGetParams {
	return &CatalogProductAttributeMediaGalleryManagementV1GetGetParams{
		Context: ctx,
	}
}

// NewCatalogProductAttributeMediaGalleryManagementV1GetGetParamsWithHTTPClient creates a new CatalogProductAttributeMediaGalleryManagementV1GetGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogProductAttributeMediaGalleryManagementV1GetGetParamsWithHTTPClient(client *http.Client) *CatalogProductAttributeMediaGalleryManagementV1GetGetParams {
	return &CatalogProductAttributeMediaGalleryManagementV1GetGetParams{
		HTTPClient: client,
	}
}

/* CatalogProductAttributeMediaGalleryManagementV1GetGetParams contains all the parameters to send to the API endpoint
   for the catalog product attribute media gallery management v1 get get operation.

   Typically these are written to a http.Request.
*/
type CatalogProductAttributeMediaGalleryManagementV1GetGetParams struct {

	// EntryID.
	EntryID int64

	// Sku.
	Sku string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the catalog product attribute media gallery management v1 get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogProductAttributeMediaGalleryManagementV1GetGetParams) WithDefaults() *CatalogProductAttributeMediaGalleryManagementV1GetGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog product attribute media gallery management v1 get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogProductAttributeMediaGalleryManagementV1GetGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog product attribute media gallery management v1 get get params
func (o *CatalogProductAttributeMediaGalleryManagementV1GetGetParams) WithTimeout(timeout time.Duration) *CatalogProductAttributeMediaGalleryManagementV1GetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog product attribute media gallery management v1 get get params
func (o *CatalogProductAttributeMediaGalleryManagementV1GetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog product attribute media gallery management v1 get get params
func (o *CatalogProductAttributeMediaGalleryManagementV1GetGetParams) WithContext(ctx context.Context) *CatalogProductAttributeMediaGalleryManagementV1GetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog product attribute media gallery management v1 get get params
func (o *CatalogProductAttributeMediaGalleryManagementV1GetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog product attribute media gallery management v1 get get params
func (o *CatalogProductAttributeMediaGalleryManagementV1GetGetParams) WithHTTPClient(client *http.Client) *CatalogProductAttributeMediaGalleryManagementV1GetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog product attribute media gallery management v1 get get params
func (o *CatalogProductAttributeMediaGalleryManagementV1GetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntryID adds the entryID to the catalog product attribute media gallery management v1 get get params
func (o *CatalogProductAttributeMediaGalleryManagementV1GetGetParams) WithEntryID(entryID int64) *CatalogProductAttributeMediaGalleryManagementV1GetGetParams {
	o.SetEntryID(entryID)
	return o
}

// SetEntryID adds the entryId to the catalog product attribute media gallery management v1 get get params
func (o *CatalogProductAttributeMediaGalleryManagementV1GetGetParams) SetEntryID(entryID int64) {
	o.EntryID = entryID
}

// WithSku adds the sku to the catalog product attribute media gallery management v1 get get params
func (o *CatalogProductAttributeMediaGalleryManagementV1GetGetParams) WithSku(sku string) *CatalogProductAttributeMediaGalleryManagementV1GetGetParams {
	o.SetSku(sku)
	return o
}

// SetSku adds the sku to the catalog product attribute media gallery management v1 get get params
func (o *CatalogProductAttributeMediaGalleryManagementV1GetGetParams) SetSku(sku string) {
	o.Sku = sku
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogProductAttributeMediaGalleryManagementV1GetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entryId
	if err := r.SetPathParam("entryId", swag.FormatInt64(o.EntryID)); err != nil {
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
