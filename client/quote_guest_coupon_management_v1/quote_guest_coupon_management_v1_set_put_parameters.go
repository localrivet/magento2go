// Code generated by go-swagger; DO NOT EDIT.

package quote_guest_coupon_management_v1

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

// NewQuoteGuestCouponManagementV1SetPutParams creates a new QuoteGuestCouponManagementV1SetPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQuoteGuestCouponManagementV1SetPutParams() *QuoteGuestCouponManagementV1SetPutParams {
	return &QuoteGuestCouponManagementV1SetPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQuoteGuestCouponManagementV1SetPutParamsWithTimeout creates a new QuoteGuestCouponManagementV1SetPutParams object
// with the ability to set a timeout on a request.
func NewQuoteGuestCouponManagementV1SetPutParamsWithTimeout(timeout time.Duration) *QuoteGuestCouponManagementV1SetPutParams {
	return &QuoteGuestCouponManagementV1SetPutParams{
		timeout: timeout,
	}
}

// NewQuoteGuestCouponManagementV1SetPutParamsWithContext creates a new QuoteGuestCouponManagementV1SetPutParams object
// with the ability to set a context for a request.
func NewQuoteGuestCouponManagementV1SetPutParamsWithContext(ctx context.Context) *QuoteGuestCouponManagementV1SetPutParams {
	return &QuoteGuestCouponManagementV1SetPutParams{
		Context: ctx,
	}
}

// NewQuoteGuestCouponManagementV1SetPutParamsWithHTTPClient creates a new QuoteGuestCouponManagementV1SetPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewQuoteGuestCouponManagementV1SetPutParamsWithHTTPClient(client *http.Client) *QuoteGuestCouponManagementV1SetPutParams {
	return &QuoteGuestCouponManagementV1SetPutParams{
		HTTPClient: client,
	}
}

/* QuoteGuestCouponManagementV1SetPutParams contains all the parameters to send to the API endpoint
   for the quote guest coupon management v1 set put operation.

   Typically these are written to a http.Request.
*/
type QuoteGuestCouponManagementV1SetPutParams struct {

	/* CartID.

	   The cart ID.
	*/
	CartID string

	/* CouponCode.

	   The coupon code data.
	*/
	CouponCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the quote guest coupon management v1 set put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuoteGuestCouponManagementV1SetPutParams) WithDefaults() *QuoteGuestCouponManagementV1SetPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the quote guest coupon management v1 set put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuoteGuestCouponManagementV1SetPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the quote guest coupon management v1 set put params
func (o *QuoteGuestCouponManagementV1SetPutParams) WithTimeout(timeout time.Duration) *QuoteGuestCouponManagementV1SetPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quote guest coupon management v1 set put params
func (o *QuoteGuestCouponManagementV1SetPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quote guest coupon management v1 set put params
func (o *QuoteGuestCouponManagementV1SetPutParams) WithContext(ctx context.Context) *QuoteGuestCouponManagementV1SetPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quote guest coupon management v1 set put params
func (o *QuoteGuestCouponManagementV1SetPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quote guest coupon management v1 set put params
func (o *QuoteGuestCouponManagementV1SetPutParams) WithHTTPClient(client *http.Client) *QuoteGuestCouponManagementV1SetPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quote guest coupon management v1 set put params
func (o *QuoteGuestCouponManagementV1SetPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the quote guest coupon management v1 set put params
func (o *QuoteGuestCouponManagementV1SetPutParams) WithCartID(cartID string) *QuoteGuestCouponManagementV1SetPutParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the quote guest coupon management v1 set put params
func (o *QuoteGuestCouponManagementV1SetPutParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithCouponCode adds the couponCode to the quote guest coupon management v1 set put params
func (o *QuoteGuestCouponManagementV1SetPutParams) WithCouponCode(couponCode string) *QuoteGuestCouponManagementV1SetPutParams {
	o.SetCouponCode(couponCode)
	return o
}

// SetCouponCode adds the couponCode to the quote guest coupon management v1 set put params
func (o *QuoteGuestCouponManagementV1SetPutParams) SetCouponCode(couponCode string) {
	o.CouponCode = couponCode
}

// WriteToRequest writes these params to a swagger request
func (o *QuoteGuestCouponManagementV1SetPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}

	// path param couponCode
	if err := r.SetPathParam("couponCode", o.CouponCode); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
