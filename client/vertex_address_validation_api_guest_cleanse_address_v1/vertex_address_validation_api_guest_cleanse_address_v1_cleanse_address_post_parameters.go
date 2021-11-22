// Code generated by go-swagger; DO NOT EDIT.

package vertex_address_validation_api_guest_cleanse_address_v1

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

// NewVertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams creates a new VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams() *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams {
	return &VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParamsWithTimeout creates a new VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams object
// with the ability to set a timeout on a request.
func NewVertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParamsWithTimeout(timeout time.Duration) *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams {
	return &VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams{
		timeout: timeout,
	}
}

// NewVertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParamsWithContext creates a new VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams object
// with the ability to set a context for a request.
func NewVertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParamsWithContext(ctx context.Context) *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams {
	return &VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams{
		Context: ctx,
	}
}

// NewVertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParamsWithHTTPClient creates a new VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewVertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParamsWithHTTPClient(client *http.Client) *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams {
	return &VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams{
		HTTPClient: client,
	}
}

/* VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams contains all the parameters to send to the API endpoint
   for the vertex address validation Api guest cleanse address v1 cleanse address post operation.

   Typically these are written to a http.Request.
*/
type VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams struct {

	// VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody.
	VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vertex address validation Api guest cleanse address v1 cleanse address post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams) WithDefaults() *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vertex address validation Api guest cleanse address v1 cleanse address post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vertex address validation Api guest cleanse address v1 cleanse address post params
func (o *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams) WithTimeout(timeout time.Duration) *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vertex address validation Api guest cleanse address v1 cleanse address post params
func (o *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vertex address validation Api guest cleanse address v1 cleanse address post params
func (o *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams) WithContext(ctx context.Context) *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vertex address validation Api guest cleanse address v1 cleanse address post params
func (o *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vertex address validation Api guest cleanse address v1 cleanse address post params
func (o *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams) WithHTTPClient(client *http.Client) *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vertex address validation Api guest cleanse address v1 cleanse address post params
func (o *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody adds the vertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody to the vertex address validation Api guest cleanse address v1 cleanse address post params
func (o *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams) WithVertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody(vertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody) *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams {
	o.SetVertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody(vertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody)
	return o
}

// SetVertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody adds the vertexAddressValidationApiGuestCleanseAddressV1CleanseAddressPostBody to the vertex address validation Api guest cleanse address v1 cleanse address post params
func (o *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams) SetVertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody(vertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody) {
	o.VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody = vertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.VertexAddressValidationAPIGuestCleanseAddressV1CleanseAddressPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
