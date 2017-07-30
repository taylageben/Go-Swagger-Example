// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTodoParams creates a new GetTodoParams object
// with the default values initialized.
func NewGetTodoParams() GetTodoParams {
	var ()
	return GetTodoParams{}
}

// GetTodoParams contains all the bound params for the get todo operation
// typically these are obtained from a http.Request
//
// swagger:parameters get_todo
type GetTodoParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: path
	*/
	TodoID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetTodoParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rTodoID, rhkTodoID, _ := route.Params.GetOK("todoId")
	if err := o.bindTodoID(rTodoID, rhkTodoID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTodoParams) bindTodoID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.TodoID = raw

	return nil
}
