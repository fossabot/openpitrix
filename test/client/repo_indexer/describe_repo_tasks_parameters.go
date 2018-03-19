// Code generated by go-swagger; DO NOT EDIT.

package repo_indexer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeRepoTasksParams creates a new DescribeRepoTasksParams object
// with the default values initialized.
func NewDescribeRepoTasksParams() *DescribeRepoTasksParams {
	var ()
	return &DescribeRepoTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeRepoTasksParamsWithTimeout creates a new DescribeRepoTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeRepoTasksParamsWithTimeout(timeout time.Duration) *DescribeRepoTasksParams {
	var ()
	return &DescribeRepoTasksParams{

		timeout: timeout,
	}
}

// NewDescribeRepoTasksParamsWithContext creates a new DescribeRepoTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeRepoTasksParamsWithContext(ctx context.Context) *DescribeRepoTasksParams {
	var ()
	return &DescribeRepoTasksParams{

		Context: ctx,
	}
}

// NewDescribeRepoTasksParamsWithHTTPClient creates a new DescribeRepoTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeRepoTasksParamsWithHTTPClient(client *http.Client) *DescribeRepoTasksParams {
	var ()
	return &DescribeRepoTasksParams{
		HTTPClient: client,
	}
}

/*DescribeRepoTasksParams contains all the parameters to send to the API endpoint
for the describe repo tasks operation typically these are written to a http.Request
*/
type DescribeRepoTasksParams struct {

	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64
	/*Owner*/
	Owner []string
	/*RepoID*/
	RepoID []string
	/*RepoTaskID*/
	RepoTaskID []string
	/*Status*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe repo tasks params
func (o *DescribeRepoTasksParams) WithTimeout(timeout time.Duration) *DescribeRepoTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe repo tasks params
func (o *DescribeRepoTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe repo tasks params
func (o *DescribeRepoTasksParams) WithContext(ctx context.Context) *DescribeRepoTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe repo tasks params
func (o *DescribeRepoTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe repo tasks params
func (o *DescribeRepoTasksParams) WithHTTPClient(client *http.Client) *DescribeRepoTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe repo tasks params
func (o *DescribeRepoTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the describe repo tasks params
func (o *DescribeRepoTasksParams) WithLimit(limit *int64) *DescribeRepoTasksParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe repo tasks params
func (o *DescribeRepoTasksParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the describe repo tasks params
func (o *DescribeRepoTasksParams) WithOffset(offset *int64) *DescribeRepoTasksParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe repo tasks params
func (o *DescribeRepoTasksParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwner adds the owner to the describe repo tasks params
func (o *DescribeRepoTasksParams) WithOwner(owner []string) *DescribeRepoTasksParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the describe repo tasks params
func (o *DescribeRepoTasksParams) SetOwner(owner []string) {
	o.Owner = owner
}

// WithRepoID adds the repoID to the describe repo tasks params
func (o *DescribeRepoTasksParams) WithRepoID(repoID []string) *DescribeRepoTasksParams {
	o.SetRepoID(repoID)
	return o
}

// SetRepoID adds the repoId to the describe repo tasks params
func (o *DescribeRepoTasksParams) SetRepoID(repoID []string) {
	o.RepoID = repoID
}

// WithRepoTaskID adds the repoTaskID to the describe repo tasks params
func (o *DescribeRepoTasksParams) WithRepoTaskID(repoTaskID []string) *DescribeRepoTasksParams {
	o.SetRepoTaskID(repoTaskID)
	return o
}

// SetRepoTaskID adds the repoTaskId to the describe repo tasks params
func (o *DescribeRepoTasksParams) SetRepoTaskID(repoTaskID []string) {
	o.RepoTaskID = repoTaskID
}

// WithStatus adds the status to the describe repo tasks params
func (o *DescribeRepoTasksParams) WithStatus(status []string) *DescribeRepoTasksParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe repo tasks params
func (o *DescribeRepoTasksParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeRepoTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesOwner := o.Owner

	joinedOwner := swag.JoinByFormat(valuesOwner, "")
	// query array param owner
	if err := r.SetQueryParam("owner", joinedOwner...); err != nil {
		return err
	}

	valuesRepoID := o.RepoID

	joinedRepoID := swag.JoinByFormat(valuesRepoID, "")
	// query array param repo_id
	if err := r.SetQueryParam("repo_id", joinedRepoID...); err != nil {
		return err
	}

	valuesRepoTaskID := o.RepoTaskID

	joinedRepoTaskID := swag.JoinByFormat(valuesRepoTaskID, "")
	// query array param repo_task_id
	if err := r.SetQueryParam("repo_task_id", joinedRepoTaskID...); err != nil {
		return err
	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}