// Code generated by go-swagger; DO NOT EDIT.

package checkout_guest_shipping_information_management_v1

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

// NewCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams creates a new CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams() *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams {
	return &CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParamsWithTimeout creates a new CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams object
// with the ability to set a timeout on a request.
func NewCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParamsWithTimeout(timeout time.Duration) *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams {
	return &CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams{
		timeout: timeout,
	}
}

// NewCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParamsWithContext creates a new CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams object
// with the ability to set a context for a request.
func NewCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParamsWithContext(ctx context.Context) *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams {
	return &CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams{
		Context: ctx,
	}
}

// NewCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParamsWithHTTPClient creates a new CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParamsWithHTTPClient(client *http.Client) *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams {
	return &CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams{
		HTTPClient: client,
	}
}

/* CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams contains all the parameters to send to the API endpoint
   for the checkout guest shipping information management v1 save address information post operation.

   Typically these are written to a http.Request.
*/
type CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams struct {

	// CartID.
	CartID string

	// CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody.
	CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the checkout guest shipping information management v1 save address information post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams) WithDefaults() *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the checkout guest shipping information management v1 save address information post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the checkout guest shipping information management v1 save address information post params
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams) WithTimeout(timeout time.Duration) *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checkout guest shipping information management v1 save address information post params
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checkout guest shipping information management v1 save address information post params
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams) WithContext(ctx context.Context) *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checkout guest shipping information management v1 save address information post params
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checkout guest shipping information management v1 save address information post params
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams) WithHTTPClient(client *http.Client) *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checkout guest shipping information management v1 save address information post params
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the checkout guest shipping information management v1 save address information post params
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams) WithCartID(cartID string) *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the checkout guest shipping information management v1 save address information post params
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody adds the checkoutGuestShippingInformationManagementV1SaveAddressInformationPostBody to the checkout guest shipping information management v1 save address information post params
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams) WithCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody(checkoutGuestShippingInformationManagementV1SaveAddressInformationPostBody CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody) *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams {
	o.SetCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody(checkoutGuestShippingInformationManagementV1SaveAddressInformationPostBody)
	return o
}

// SetCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody adds the checkoutGuestShippingInformationManagementV1SaveAddressInformationPostBody to the checkout guest shipping information management v1 save address information post params
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams) SetCheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody(checkoutGuestShippingInformationManagementV1SaveAddressInformationPostBody CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody) {
	o.CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody = checkoutGuestShippingInformationManagementV1SaveAddressInformationPostBody
}

// WriteToRequest writes these params to a swagger request
func (o *CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.CheckoutGuestShippingInformationManagementV1SaveAddressInformationPostBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
