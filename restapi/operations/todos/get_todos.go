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

// GetTodosHandlerFunc turns a function with the right signature into a get todos handler
type GetTodosHandlerFunc func(GetTodosParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTodosHandlerFunc) Handle(params GetTodosParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetTodosHandler interface for that can handle valid get todos params
type GetTodosHandler interface {
	Handle(GetTodosParams, interface{}) middleware.Responder
}

// NewGetTodos creates a new http.Handler for the get todos operation
func NewGetTodos(ctx *middleware.Context, handler GetTodosHandler) *GetTodos {
	return &GetTodos{Context: ctx, Handler: handler}
}

/*GetTodos swagger:route GET /todos Todos getTodos

List Todos

*/
type GetTodos struct {
	Context *middleware.Context
	Handler GetTodosHandler
}

func (o *GetTodos) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTodosParams()

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

// GetTodosForbiddenBody get todos forbidden body
// swagger:model GetTodosForbiddenBody
type GetTodosForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this get todos forbidden body
func (o *GetTodosForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTodosForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getTodosForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTodosForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTodosForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetTodosForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetTodosInternalServerErrorBody get todos internal server error body
// swagger:model GetTodosInternalServerErrorBody
type GetTodosInternalServerErrorBody interface{}

// GetTodosNotFoundBody get todos not found body
// swagger:model GetTodosNotFoundBody
type GetTodosNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this get todos not found body
func (o *GetTodosNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *GetTodosNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getTodosNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *GetTodosNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("getTodosNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTodosNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTodosNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetTodosNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetTodosUnauthorizedBody get todos unauthorized body
// swagger:model GetTodosUnauthorizedBody
type GetTodosUnauthorizedBody interface{}