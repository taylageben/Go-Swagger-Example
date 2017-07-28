// Code generated by go-swagger; DO NOT EDIT.

package misc

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

// GetMockedHandlerFunc turns a function with the right signature into a get mocked handler
type GetMockedHandlerFunc func(GetMockedParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMockedHandlerFunc) Handle(params GetMockedParams) middleware.Responder {
	return fn(params)
}

// GetMockedHandler interface for that can handle valid get mocked params
type GetMockedHandler interface {
	Handle(GetMockedParams) middleware.Responder
}

// NewGetMocked creates a new http.Handler for the get mocked operation
func NewGetMocked(ctx *middleware.Context, handler GetMockedHandler) *GetMocked {
	return &GetMocked{Context: ctx, Handler: handler}
}

/*GetMocked swagger:route GET /mocked Misc getMocked

Mock Example

This is an example of a mocked endpoint.

This endpoint does not actually exist in the api - try visiting [http://todos.stoplight.io/mocked](http://todos.stoplight.io). You will see the default response (same as you get when you visit the root "/" url).

We have defined a 200 response below, with an example, and then explicitly turned on mocking for this endpoint via the dropdown in the right sidebar.

With this enabled, if you visit {your prism instance host}/mocked, you'll see the mocked example. You can find the mock server host for this API by clicking the "Design Dashboard" button at the top of this screen. You can also try sending a test request by clicking "send test request" button in the right sidebar.

*/
type GetMocked struct {
	Context *middleware.Context
	Handler GetMockedHandler
}

func (o *GetMocked) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetMockedParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetMockedOKBody get mocked o k body
// swagger:model GetMockedOKBody
type GetMockedOKBody struct {

	// mocked
	// Required: true
	Mocked *bool `json:"mocked"`

	// this
	// Required: true
	This *string `json:"this"`
}

// Validate validates this get mocked o k body
func (o *GetMockedOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMocked(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateThis(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMockedOKBody) validateMocked(formats strfmt.Registry) error {

	if err := validate.Required("getMockedOK"+"."+"mocked", "body", o.Mocked); err != nil {
		return err
	}

	return nil
}

func (o *GetMockedOKBody) validateThis(formats strfmt.Registry) error {

	if err := validate.Required("getMockedOK"+"."+"this", "body", o.This); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetMockedOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMockedOKBody) UnmarshalBinary(b []byte) error {
	var res GetMockedOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
