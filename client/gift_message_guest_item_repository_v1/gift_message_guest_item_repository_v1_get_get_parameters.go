// Code generated by go-swagger; DO NOT EDIT.

package gift_message_guest_item_repository_v1

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

// NewGiftMessageGuestItemRepositoryV1GetGetParams creates a new GiftMessageGuestItemRepositoryV1GetGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGiftMessageGuestItemRepositoryV1GetGetParams() *GiftMessageGuestItemRepositoryV1GetGetParams {
	return &GiftMessageGuestItemRepositoryV1GetGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGiftMessageGuestItemRepositoryV1GetGetParamsWithTimeout creates a new GiftMessageGuestItemRepositoryV1GetGetParams object
// with the ability to set a timeout on a request.
func NewGiftMessageGuestItemRepositoryV1GetGetParamsWithTimeout(timeout time.Duration) *GiftMessageGuestItemRepositoryV1GetGetParams {
	return &GiftMessageGuestItemRepositoryV1GetGetParams{
		timeout: timeout,
	}
}

// NewGiftMessageGuestItemRepositoryV1GetGetParamsWithContext creates a new GiftMessageGuestItemRepositoryV1GetGetParams object
// with the ability to set a context for a request.
func NewGiftMessageGuestItemRepositoryV1GetGetParamsWithContext(ctx context.Context) *GiftMessageGuestItemRepositoryV1GetGetParams {
	return &GiftMessageGuestItemRepositoryV1GetGetParams{
		Context: ctx,
	}
}

// NewGiftMessageGuestItemRepositoryV1GetGetParamsWithHTTPClient creates a new GiftMessageGuestItemRepositoryV1GetGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGiftMessageGuestItemRepositoryV1GetGetParamsWithHTTPClient(client *http.Client) *GiftMessageGuestItemRepositoryV1GetGetParams {
	return &GiftMessageGuestItemRepositoryV1GetGetParams{
		HTTPClient: client,
	}
}

/* GiftMessageGuestItemRepositoryV1GetGetParams contains all the parameters to send to the API endpoint
   for the gift message guest item repository v1 get get operation.

   Typically these are written to a http.Request.
*/
type GiftMessageGuestItemRepositoryV1GetGetParams struct {

	/* CartID.

	   The shopping cart ID.
	*/
	CartID string

	/* ItemID.

	   The item ID.
	*/
	ItemID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gift message guest item repository v1 get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftMessageGuestItemRepositoryV1GetGetParams) WithDefaults() *GiftMessageGuestItemRepositoryV1GetGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gift message guest item repository v1 get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftMessageGuestItemRepositoryV1GetGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the gift message guest item repository v1 get get params
func (o *GiftMessageGuestItemRepositoryV1GetGetParams) WithTimeout(timeout time.Duration) *GiftMessageGuestItemRepositoryV1GetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gift message guest item repository v1 get get params
func (o *GiftMessageGuestItemRepositoryV1GetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gift message guest item repository v1 get get params
func (o *GiftMessageGuestItemRepositoryV1GetGetParams) WithContext(ctx context.Context) *GiftMessageGuestItemRepositoryV1GetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gift message guest item repository v1 get get params
func (o *GiftMessageGuestItemRepositoryV1GetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gift message guest item repository v1 get get params
func (o *GiftMessageGuestItemRepositoryV1GetGetParams) WithHTTPClient(client *http.Client) *GiftMessageGuestItemRepositoryV1GetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gift message guest item repository v1 get get params
func (o *GiftMessageGuestItemRepositoryV1GetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the gift message guest item repository v1 get get params
func (o *GiftMessageGuestItemRepositoryV1GetGetParams) WithCartID(cartID string) *GiftMessageGuestItemRepositoryV1GetGetParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the gift message guest item repository v1 get get params
func (o *GiftMessageGuestItemRepositoryV1GetGetParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithItemID adds the itemID to the gift message guest item repository v1 get get params
func (o *GiftMessageGuestItemRepositoryV1GetGetParams) WithItemID(itemID int64) *GiftMessageGuestItemRepositoryV1GetGetParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the gift message guest item repository v1 get get params
func (o *GiftMessageGuestItemRepositoryV1GetGetParams) SetItemID(itemID int64) {
	o.ItemID = itemID
}

// WriteToRequest writes these params to a swagger request
func (o *GiftMessageGuestItemRepositoryV1GetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}

	// path param itemId
	if err := r.SetPathParam("itemId", swag.FormatInt64(o.ItemID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
