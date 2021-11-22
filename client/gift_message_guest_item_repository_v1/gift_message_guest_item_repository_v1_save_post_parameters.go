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

// NewGiftMessageGuestItemRepositoryV1SavePostParams creates a new GiftMessageGuestItemRepositoryV1SavePostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGiftMessageGuestItemRepositoryV1SavePostParams() *GiftMessageGuestItemRepositoryV1SavePostParams {
	return &GiftMessageGuestItemRepositoryV1SavePostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGiftMessageGuestItemRepositoryV1SavePostParamsWithTimeout creates a new GiftMessageGuestItemRepositoryV1SavePostParams object
// with the ability to set a timeout on a request.
func NewGiftMessageGuestItemRepositoryV1SavePostParamsWithTimeout(timeout time.Duration) *GiftMessageGuestItemRepositoryV1SavePostParams {
	return &GiftMessageGuestItemRepositoryV1SavePostParams{
		timeout: timeout,
	}
}

// NewGiftMessageGuestItemRepositoryV1SavePostParamsWithContext creates a new GiftMessageGuestItemRepositoryV1SavePostParams object
// with the ability to set a context for a request.
func NewGiftMessageGuestItemRepositoryV1SavePostParamsWithContext(ctx context.Context) *GiftMessageGuestItemRepositoryV1SavePostParams {
	return &GiftMessageGuestItemRepositoryV1SavePostParams{
		Context: ctx,
	}
}

// NewGiftMessageGuestItemRepositoryV1SavePostParamsWithHTTPClient creates a new GiftMessageGuestItemRepositoryV1SavePostParams object
// with the ability to set a custom HTTPClient for a request.
func NewGiftMessageGuestItemRepositoryV1SavePostParamsWithHTTPClient(client *http.Client) *GiftMessageGuestItemRepositoryV1SavePostParams {
	return &GiftMessageGuestItemRepositoryV1SavePostParams{
		HTTPClient: client,
	}
}

/* GiftMessageGuestItemRepositoryV1SavePostParams contains all the parameters to send to the API endpoint
   for the gift message guest item repository v1 save post operation.

   Typically these are written to a http.Request.
*/
type GiftMessageGuestItemRepositoryV1SavePostParams struct {

	/* CartID.

	   The cart ID.
	*/
	CartID string

	// GiftMessageGuestItemRepositoryV1SavePostBody.
	GiftMessageGuestItemRepositoryV1SavePostBody GiftMessageGuestItemRepositoryV1SavePostBody

	/* ItemID.

	   The item ID.
	*/
	ItemID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gift message guest item repository v1 save post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftMessageGuestItemRepositoryV1SavePostParams) WithDefaults() *GiftMessageGuestItemRepositoryV1SavePostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gift message guest item repository v1 save post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GiftMessageGuestItemRepositoryV1SavePostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the gift message guest item repository v1 save post params
func (o *GiftMessageGuestItemRepositoryV1SavePostParams) WithTimeout(timeout time.Duration) *GiftMessageGuestItemRepositoryV1SavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gift message guest item repository v1 save post params
func (o *GiftMessageGuestItemRepositoryV1SavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gift message guest item repository v1 save post params
func (o *GiftMessageGuestItemRepositoryV1SavePostParams) WithContext(ctx context.Context) *GiftMessageGuestItemRepositoryV1SavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gift message guest item repository v1 save post params
func (o *GiftMessageGuestItemRepositoryV1SavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gift message guest item repository v1 save post params
func (o *GiftMessageGuestItemRepositoryV1SavePostParams) WithHTTPClient(client *http.Client) *GiftMessageGuestItemRepositoryV1SavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gift message guest item repository v1 save post params
func (o *GiftMessageGuestItemRepositoryV1SavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartID adds the cartID to the gift message guest item repository v1 save post params
func (o *GiftMessageGuestItemRepositoryV1SavePostParams) WithCartID(cartID string) *GiftMessageGuestItemRepositoryV1SavePostParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the gift message guest item repository v1 save post params
func (o *GiftMessageGuestItemRepositoryV1SavePostParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithGiftMessageGuestItemRepositoryV1SavePostBody adds the giftMessageGuestItemRepositoryV1SavePostBody to the gift message guest item repository v1 save post params
func (o *GiftMessageGuestItemRepositoryV1SavePostParams) WithGiftMessageGuestItemRepositoryV1SavePostBody(giftMessageGuestItemRepositoryV1SavePostBody GiftMessageGuestItemRepositoryV1SavePostBody) *GiftMessageGuestItemRepositoryV1SavePostParams {
	o.SetGiftMessageGuestItemRepositoryV1SavePostBody(giftMessageGuestItemRepositoryV1SavePostBody)
	return o
}

// SetGiftMessageGuestItemRepositoryV1SavePostBody adds the giftMessageGuestItemRepositoryV1SavePostBody to the gift message guest item repository v1 save post params
func (o *GiftMessageGuestItemRepositoryV1SavePostParams) SetGiftMessageGuestItemRepositoryV1SavePostBody(giftMessageGuestItemRepositoryV1SavePostBody GiftMessageGuestItemRepositoryV1SavePostBody) {
	o.GiftMessageGuestItemRepositoryV1SavePostBody = giftMessageGuestItemRepositoryV1SavePostBody
}

// WithItemID adds the itemID to the gift message guest item repository v1 save post params
func (o *GiftMessageGuestItemRepositoryV1SavePostParams) WithItemID(itemID int64) *GiftMessageGuestItemRepositoryV1SavePostParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the gift message guest item repository v1 save post params
func (o *GiftMessageGuestItemRepositoryV1SavePostParams) SetItemID(itemID int64) {
	o.ItemID = itemID
}

// WriteToRequest writes these params to a swagger request
func (o *GiftMessageGuestItemRepositoryV1SavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.GiftMessageGuestItemRepositoryV1SavePostBody); err != nil {
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