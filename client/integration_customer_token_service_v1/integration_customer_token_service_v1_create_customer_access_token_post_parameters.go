// Code generated by go-swagger; DO NOT EDIT.

package integration_customer_token_service_v1

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

// NewIntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams creates a new IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams() *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams {
	return &IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParamsWithTimeout creates a new IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams object
// with the ability to set a timeout on a request.
func NewIntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParamsWithTimeout(timeout time.Duration) *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams {
	return &IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams{
		timeout: timeout,
	}
}

// NewIntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParamsWithContext creates a new IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams object
// with the ability to set a context for a request.
func NewIntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParamsWithContext(ctx context.Context) *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams {
	return &IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams{
		Context: ctx,
	}
}

// NewIntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParamsWithHTTPClient creates a new IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewIntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParamsWithHTTPClient(client *http.Client) *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams {
	return &IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams{
		HTTPClient: client,
	}
}

/* IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams contains all the parameters to send to the API endpoint
   for the integration customer token service v1 create customer access token post operation.

   Typically these are written to a http.Request.
*/
type IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams struct {

	// IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody.
	IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the integration customer token service v1 create customer access token post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams) WithDefaults() *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the integration customer token service v1 create customer access token post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the integration customer token service v1 create customer access token post params
func (o *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams) WithTimeout(timeout time.Duration) *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the integration customer token service v1 create customer access token post params
func (o *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the integration customer token service v1 create customer access token post params
func (o *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams) WithContext(ctx context.Context) *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the integration customer token service v1 create customer access token post params
func (o *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the integration customer token service v1 create customer access token post params
func (o *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams) WithHTTPClient(client *http.Client) *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the integration customer token service v1 create customer access token post params
func (o *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody adds the integrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody to the integration customer token service v1 create customer access token post params
func (o *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams) WithIntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody(integrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody) *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams {
	o.SetIntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody(integrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody)
	return o
}

// SetIntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody adds the integrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody to the integration customer token service v1 create customer access token post params
func (o *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams) SetIntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody(integrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody) {
	o.IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody = integrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.IntegrationCustomerTokenServiceV1CreateCustomerAccessTokenPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
