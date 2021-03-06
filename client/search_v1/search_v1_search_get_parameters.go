// Code generated by go-swagger; DO NOT EDIT.

package search_v1

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

// NewSearchV1SearchGetParams creates a new SearchV1SearchGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchV1SearchGetParams() *SearchV1SearchGetParams {
	return &SearchV1SearchGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchV1SearchGetParamsWithTimeout creates a new SearchV1SearchGetParams object
// with the ability to set a timeout on a request.
func NewSearchV1SearchGetParamsWithTimeout(timeout time.Duration) *SearchV1SearchGetParams {
	return &SearchV1SearchGetParams{
		timeout: timeout,
	}
}

// NewSearchV1SearchGetParamsWithContext creates a new SearchV1SearchGetParams object
// with the ability to set a context for a request.
func NewSearchV1SearchGetParamsWithContext(ctx context.Context) *SearchV1SearchGetParams {
	return &SearchV1SearchGetParams{
		Context: ctx,
	}
}

// NewSearchV1SearchGetParamsWithHTTPClient creates a new SearchV1SearchGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchV1SearchGetParamsWithHTTPClient(client *http.Client) *SearchV1SearchGetParams {
	return &SearchV1SearchGetParams{
		HTTPClient: client,
	}
}

/* SearchV1SearchGetParams contains all the parameters to send to the API endpoint
   for the search v1 search get operation.

   Typically these are written to a http.Request.
*/
type SearchV1SearchGetParams struct {

	/* SearchCriteriaCurrentPage.

	   Current page.
	*/
	SearchCriteriaCurrentPage *int64

	/* SearchCriteriaFilterGroups0Filters0ConditionType.

	   Condition type
	*/
	SearchCriteriaFilterGroups0Filters0ConditionType *string

	/* SearchCriteriaFilterGroups0Filters0Field.

	   Field
	*/
	SearchCriteriaFilterGroups0Filters0Field *string

	/* SearchCriteriaFilterGroups0Filters0Value.

	   Value
	*/
	SearchCriteriaFilterGroups0Filters0Value *string

	/* SearchCriteriaPageSize.

	   Page size.
	*/
	SearchCriteriaPageSize *int64

	// SearchCriteriaRequestName.
	SearchCriteriaRequestName *string

	/* SearchCriteriaSortOrders0Direction.

	   Sorting direction.
	*/
	SearchCriteriaSortOrders0Direction *string

	/* SearchCriteriaSortOrders0Field.

	   Sorting field.
	*/
	SearchCriteriaSortOrders0Field *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search v1 search get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchV1SearchGetParams) WithDefaults() *SearchV1SearchGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search v1 search get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchV1SearchGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search v1 search get params
func (o *SearchV1SearchGetParams) WithTimeout(timeout time.Duration) *SearchV1SearchGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search v1 search get params
func (o *SearchV1SearchGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search v1 search get params
func (o *SearchV1SearchGetParams) WithContext(ctx context.Context) *SearchV1SearchGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search v1 search get params
func (o *SearchV1SearchGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search v1 search get params
func (o *SearchV1SearchGetParams) WithHTTPClient(client *http.Client) *SearchV1SearchGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search v1 search get params
func (o *SearchV1SearchGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSearchCriteriaCurrentPage adds the searchCriteriaCurrentPage to the search v1 search get params
func (o *SearchV1SearchGetParams) WithSearchCriteriaCurrentPage(searchCriteriaCurrentPage *int64) *SearchV1SearchGetParams {
	o.SetSearchCriteriaCurrentPage(searchCriteriaCurrentPage)
	return o
}

// SetSearchCriteriaCurrentPage adds the searchCriteriaCurrentPage to the search v1 search get params
func (o *SearchV1SearchGetParams) SetSearchCriteriaCurrentPage(searchCriteriaCurrentPage *int64) {
	o.SearchCriteriaCurrentPage = searchCriteriaCurrentPage
}

// WithSearchCriteriaFilterGroups0Filters0ConditionType adds the searchCriteriaFilterGroups0Filters0ConditionType to the search v1 search get params
func (o *SearchV1SearchGetParams) WithSearchCriteriaFilterGroups0Filters0ConditionType(searchCriteriaFilterGroups0Filters0ConditionType *string) *SearchV1SearchGetParams {
	o.SetSearchCriteriaFilterGroups0Filters0ConditionType(searchCriteriaFilterGroups0Filters0ConditionType)
	return o
}

// SetSearchCriteriaFilterGroups0Filters0ConditionType adds the searchCriteriaFilterGroups0Filters0ConditionType to the search v1 search get params
func (o *SearchV1SearchGetParams) SetSearchCriteriaFilterGroups0Filters0ConditionType(searchCriteriaFilterGroups0Filters0ConditionType *string) {
	o.SearchCriteriaFilterGroups0Filters0ConditionType = searchCriteriaFilterGroups0Filters0ConditionType
}

// WithSearchCriteriaFilterGroups0Filters0Field adds the searchCriteriaFilterGroups0Filters0Field to the search v1 search get params
func (o *SearchV1SearchGetParams) WithSearchCriteriaFilterGroups0Filters0Field(searchCriteriaFilterGroups0Filters0Field *string) *SearchV1SearchGetParams {
	o.SetSearchCriteriaFilterGroups0Filters0Field(searchCriteriaFilterGroups0Filters0Field)
	return o
}

// SetSearchCriteriaFilterGroups0Filters0Field adds the searchCriteriaFilterGroups0Filters0Field to the search v1 search get params
func (o *SearchV1SearchGetParams) SetSearchCriteriaFilterGroups0Filters0Field(searchCriteriaFilterGroups0Filters0Field *string) {
	o.SearchCriteriaFilterGroups0Filters0Field = searchCriteriaFilterGroups0Filters0Field
}

// WithSearchCriteriaFilterGroups0Filters0Value adds the searchCriteriaFilterGroups0Filters0Value to the search v1 search get params
func (o *SearchV1SearchGetParams) WithSearchCriteriaFilterGroups0Filters0Value(searchCriteriaFilterGroups0Filters0Value *string) *SearchV1SearchGetParams {
	o.SetSearchCriteriaFilterGroups0Filters0Value(searchCriteriaFilterGroups0Filters0Value)
	return o
}

// SetSearchCriteriaFilterGroups0Filters0Value adds the searchCriteriaFilterGroups0Filters0Value to the search v1 search get params
func (o *SearchV1SearchGetParams) SetSearchCriteriaFilterGroups0Filters0Value(searchCriteriaFilterGroups0Filters0Value *string) {
	o.SearchCriteriaFilterGroups0Filters0Value = searchCriteriaFilterGroups0Filters0Value
}

// WithSearchCriteriaPageSize adds the searchCriteriaPageSize to the search v1 search get params
func (o *SearchV1SearchGetParams) WithSearchCriteriaPageSize(searchCriteriaPageSize *int64) *SearchV1SearchGetParams {
	o.SetSearchCriteriaPageSize(searchCriteriaPageSize)
	return o
}

// SetSearchCriteriaPageSize adds the searchCriteriaPageSize to the search v1 search get params
func (o *SearchV1SearchGetParams) SetSearchCriteriaPageSize(searchCriteriaPageSize *int64) {
	o.SearchCriteriaPageSize = searchCriteriaPageSize
}

// WithSearchCriteriaRequestName adds the searchCriteriaRequestName to the search v1 search get params
func (o *SearchV1SearchGetParams) WithSearchCriteriaRequestName(searchCriteriaRequestName *string) *SearchV1SearchGetParams {
	o.SetSearchCriteriaRequestName(searchCriteriaRequestName)
	return o
}

// SetSearchCriteriaRequestName adds the searchCriteriaRequestName to the search v1 search get params
func (o *SearchV1SearchGetParams) SetSearchCriteriaRequestName(searchCriteriaRequestName *string) {
	o.SearchCriteriaRequestName = searchCriteriaRequestName
}

// WithSearchCriteriaSortOrders0Direction adds the searchCriteriaSortOrders0Direction to the search v1 search get params
func (o *SearchV1SearchGetParams) WithSearchCriteriaSortOrders0Direction(searchCriteriaSortOrders0Direction *string) *SearchV1SearchGetParams {
	o.SetSearchCriteriaSortOrders0Direction(searchCriteriaSortOrders0Direction)
	return o
}

// SetSearchCriteriaSortOrders0Direction adds the searchCriteriaSortOrders0Direction to the search v1 search get params
func (o *SearchV1SearchGetParams) SetSearchCriteriaSortOrders0Direction(searchCriteriaSortOrders0Direction *string) {
	o.SearchCriteriaSortOrders0Direction = searchCriteriaSortOrders0Direction
}

// WithSearchCriteriaSortOrders0Field adds the searchCriteriaSortOrders0Field to the search v1 search get params
func (o *SearchV1SearchGetParams) WithSearchCriteriaSortOrders0Field(searchCriteriaSortOrders0Field *string) *SearchV1SearchGetParams {
	o.SetSearchCriteriaSortOrders0Field(searchCriteriaSortOrders0Field)
	return o
}

// SetSearchCriteriaSortOrders0Field adds the searchCriteriaSortOrders0Field to the search v1 search get params
func (o *SearchV1SearchGetParams) SetSearchCriteriaSortOrders0Field(searchCriteriaSortOrders0Field *string) {
	o.SearchCriteriaSortOrders0Field = searchCriteriaSortOrders0Field
}

// WriteToRequest writes these params to a swagger request
func (o *SearchV1SearchGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SearchCriteriaCurrentPage != nil {

		// query param searchCriteria[currentPage]
		var qrSearchCriteriaCurrentPage int64

		if o.SearchCriteriaCurrentPage != nil {
			qrSearchCriteriaCurrentPage = *o.SearchCriteriaCurrentPage
		}
		qSearchCriteriaCurrentPage := swag.FormatInt64(qrSearchCriteriaCurrentPage)
		if qSearchCriteriaCurrentPage != "" {

			if err := r.SetQueryParam("searchCriteria[currentPage]", qSearchCriteriaCurrentPage); err != nil {
				return err
			}
		}
	}

	if o.SearchCriteriaFilterGroups0Filters0ConditionType != nil {

		// query param searchCriteria[filterGroups][0][filters][0][conditionType]
		var qrSearchCriteriaFilterGroups0Filters0ConditionType string

		if o.SearchCriteriaFilterGroups0Filters0ConditionType != nil {
			qrSearchCriteriaFilterGroups0Filters0ConditionType = *o.SearchCriteriaFilterGroups0Filters0ConditionType
		}
		qSearchCriteriaFilterGroups0Filters0ConditionType := qrSearchCriteriaFilterGroups0Filters0ConditionType
		if qSearchCriteriaFilterGroups0Filters0ConditionType != "" {

			if err := r.SetQueryParam("searchCriteria[filterGroups][0][filters][0][conditionType]", qSearchCriteriaFilterGroups0Filters0ConditionType); err != nil {
				return err
			}
		}
	}

	if o.SearchCriteriaFilterGroups0Filters0Field != nil {

		// query param searchCriteria[filterGroups][0][filters][0][field]
		var qrSearchCriteriaFilterGroups0Filters0Field string

		if o.SearchCriteriaFilterGroups0Filters0Field != nil {
			qrSearchCriteriaFilterGroups0Filters0Field = *o.SearchCriteriaFilterGroups0Filters0Field
		}
		qSearchCriteriaFilterGroups0Filters0Field := qrSearchCriteriaFilterGroups0Filters0Field
		if qSearchCriteriaFilterGroups0Filters0Field != "" {

			if err := r.SetQueryParam("searchCriteria[filterGroups][0][filters][0][field]", qSearchCriteriaFilterGroups0Filters0Field); err != nil {
				return err
			}
		}
	}

	if o.SearchCriteriaFilterGroups0Filters0Value != nil {

		// query param searchCriteria[filterGroups][0][filters][0][value]
		var qrSearchCriteriaFilterGroups0Filters0Value string

		if o.SearchCriteriaFilterGroups0Filters0Value != nil {
			qrSearchCriteriaFilterGroups0Filters0Value = *o.SearchCriteriaFilterGroups0Filters0Value
		}
		qSearchCriteriaFilterGroups0Filters0Value := qrSearchCriteriaFilterGroups0Filters0Value
		if qSearchCriteriaFilterGroups0Filters0Value != "" {

			if err := r.SetQueryParam("searchCriteria[filterGroups][0][filters][0][value]", qSearchCriteriaFilterGroups0Filters0Value); err != nil {
				return err
			}
		}
	}

	if o.SearchCriteriaPageSize != nil {

		// query param searchCriteria[pageSize]
		var qrSearchCriteriaPageSize int64

		if o.SearchCriteriaPageSize != nil {
			qrSearchCriteriaPageSize = *o.SearchCriteriaPageSize
		}
		qSearchCriteriaPageSize := swag.FormatInt64(qrSearchCriteriaPageSize)
		if qSearchCriteriaPageSize != "" {

			if err := r.SetQueryParam("searchCriteria[pageSize]", qSearchCriteriaPageSize); err != nil {
				return err
			}
		}
	}

	if o.SearchCriteriaRequestName != nil {

		// query param searchCriteria[requestName]
		var qrSearchCriteriaRequestName string

		if o.SearchCriteriaRequestName != nil {
			qrSearchCriteriaRequestName = *o.SearchCriteriaRequestName
		}
		qSearchCriteriaRequestName := qrSearchCriteriaRequestName
		if qSearchCriteriaRequestName != "" {

			if err := r.SetQueryParam("searchCriteria[requestName]", qSearchCriteriaRequestName); err != nil {
				return err
			}
		}
	}

	if o.SearchCriteriaSortOrders0Direction != nil {

		// query param searchCriteria[sortOrders][0][direction]
		var qrSearchCriteriaSortOrders0Direction string

		if o.SearchCriteriaSortOrders0Direction != nil {
			qrSearchCriteriaSortOrders0Direction = *o.SearchCriteriaSortOrders0Direction
		}
		qSearchCriteriaSortOrders0Direction := qrSearchCriteriaSortOrders0Direction
		if qSearchCriteriaSortOrders0Direction != "" {

			if err := r.SetQueryParam("searchCriteria[sortOrders][0][direction]", qSearchCriteriaSortOrders0Direction); err != nil {
				return err
			}
		}
	}

	if o.SearchCriteriaSortOrders0Field != nil {

		// query param searchCriteria[sortOrders][0][field]
		var qrSearchCriteriaSortOrders0Field string

		if o.SearchCriteriaSortOrders0Field != nil {
			qrSearchCriteriaSortOrders0Field = *o.SearchCriteriaSortOrders0Field
		}
		qSearchCriteriaSortOrders0Field := qrSearchCriteriaSortOrders0Field
		if qSearchCriteriaSortOrders0Field != "" {

			if err := r.SetQueryParam("searchCriteria[sortOrders][0][field]", qSearchCriteriaSortOrders0Field); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
