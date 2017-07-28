// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// PostTodosHandlerFunc turns a function with the right signature into a post todos handler
type PostTodosHandlerFunc func(PostTodosParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostTodosHandlerFunc) Handle(params PostTodosParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostTodosHandler interface for that can handle valid post todos params
type PostTodosHandler interface {
	Handle(PostTodosParams, interface{}) middleware.Responder
}

// NewPostTodos creates a new http.Handler for the post todos operation
func NewPostTodos(ctx *middleware.Context, handler PostTodosHandler) *PostTodos {
	return &PostTodos{Context: ctx, Handler: handler}
}

/*PostTodos swagger:route POST /todos Todos postTodos

Create Todo

*/
type PostTodos struct {
	Context *middleware.Context
	Handler PostTodosHandler
}

func (o *PostTodos) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostTodosParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostTodosForbiddenBody post todos forbidden body
// swagger:model PostTodosForbiddenBody
type PostTodosForbiddenBody interface{}

// PostTodosInternalServerErrorBody post todos internal server error body
// swagger:model PostTodosInternalServerErrorBody
type PostTodosInternalServerErrorBody interface{}

// PostTodosNotFoundBody post todos not found body
// swagger:model PostTodosNotFoundBody
type PostTodosNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this post todos not found body
func (o *PostTodosNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTodosNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("postTodosNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *PostTodosNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("postTodosNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostTodosNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTodosNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PostTodosNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostTodosUnauthorizedBody post todos unauthorized body
// swagger:model PostTodosUnauthorizedBody
type PostTodosUnauthorizedBody interface{}
