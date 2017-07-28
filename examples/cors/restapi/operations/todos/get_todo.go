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

// GetTodoHandlerFunc turns a function with the right signature into a get todo handler
type GetTodoHandlerFunc func(GetTodoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTodoHandlerFunc) Handle(params GetTodoParams) middleware.Responder {
	return fn(params)
}

// GetTodoHandler interface for that can handle valid get todo params
type GetTodoHandler interface {
	Handle(GetTodoParams) middleware.Responder
}

// NewGetTodo creates a new http.Handler for the get todo operation
func NewGetTodo(ctx *middleware.Context, handler GetTodoHandler) *GetTodo {
	return &GetTodo{Context: ctx, Handler: handler}
}

/*GetTodo swagger:route GET /todos/{todoId} Todos getTodo

Get Todo

*/
type GetTodo struct {
	Context *middleware.Context
	Handler GetTodoHandler
}

func (o *GetTodo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTodoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetTodoForbiddenBody get todo forbidden body
// swagger:model GetTodoForbiddenBody
type GetTodoForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this get todo forbidden body
func (o *GetTodoForbiddenBody) Validate(formats strfmt.Registry) error {
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

func (o *GetTodoForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getTodoForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTodoForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTodoForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetTodoForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetTodoInternalServerErrorBody get todo internal server error body
// swagger:model GetTodoInternalServerErrorBody
type GetTodoInternalServerErrorBody interface{}

// GetTodoNotFoundBody get todo not found body
// swagger:model GetTodoNotFoundBody
type GetTodoNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this get todo not found body
func (o *GetTodoNotFoundBody) Validate(formats strfmt.Registry) error {
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

func (o *GetTodoNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getTodoNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *GetTodoNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("getTodoNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTodoNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTodoNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetTodoNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetTodoUnauthorizedBody get todo unauthorized body
// swagger:model GetTodoUnauthorizedBody
type GetTodoUnauthorizedBody interface{}
