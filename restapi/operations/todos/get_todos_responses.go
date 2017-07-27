// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"span/go-swagger-example/models"
)

// GetTodosOKCode is the HTTP code returned for type GetTodosOK
const GetTodosOKCode int = 200

/*GetTodosOK get todos o k

swagger:response getTodosOK
*/
type GetTodosOK struct {

	/*
	  In: Body
	*/
	Payload []*models.TodoFull `json:"body,omitempty"`
}

// NewGetTodosOK creates GetTodosOK with default headers values
func NewGetTodosOK() *GetTodosOK {
	return &GetTodosOK{}
}

// WithPayload adds the payload to the get todos o k response
func (o *GetTodosOK) WithPayload(payload []*models.TodoFull) *GetTodosOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todos o k response
func (o *GetTodosOK) SetPayload(payload []*models.TodoFull) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodosOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.TodoFull, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetTodosUnauthorizedCode is the HTTP code returned for type GetTodosUnauthorized
const GetTodosUnauthorizedCode int = 401

/*GetTodosUnauthorized get todos unauthorized

swagger:response getTodosUnauthorized
*/
type GetTodosUnauthorized struct {

	/*
	  In: Body
	*/
	Payload GetTodosUnauthorizedBody `json:"body,omitempty"`
}

// NewGetTodosUnauthorized creates GetTodosUnauthorized with default headers values
func NewGetTodosUnauthorized() *GetTodosUnauthorized {
	return &GetTodosUnauthorized{}
}

// WithPayload adds the payload to the get todos unauthorized response
func (o *GetTodosUnauthorized) WithPayload(payload GetTodosUnauthorizedBody) *GetTodosUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todos unauthorized response
func (o *GetTodosUnauthorized) SetPayload(payload GetTodosUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodosUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetTodosForbiddenCode is the HTTP code returned for type GetTodosForbidden
const GetTodosForbiddenCode int = 403

/*GetTodosForbidden get todos forbidden

swagger:response getTodosForbidden
*/
type GetTodosForbidden struct {

	/*
	  In: Body
	*/
	Payload GetTodosForbiddenBody `json:"body,omitempty"`
}

// NewGetTodosForbidden creates GetTodosForbidden with default headers values
func NewGetTodosForbidden() *GetTodosForbidden {
	return &GetTodosForbidden{}
}

// WithPayload adds the payload to the get todos forbidden response
func (o *GetTodosForbidden) WithPayload(payload GetTodosForbiddenBody) *GetTodosForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todos forbidden response
func (o *GetTodosForbidden) SetPayload(payload GetTodosForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodosForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetTodosNotFoundCode is the HTTP code returned for type GetTodosNotFound
const GetTodosNotFoundCode int = 404

/*GetTodosNotFound get todos not found

swagger:response getTodosNotFound
*/
type GetTodosNotFound struct {

	/*
	  In: Body
	*/
	Payload GetTodosNotFoundBody `json:"body,omitempty"`
}

// NewGetTodosNotFound creates GetTodosNotFound with default headers values
func NewGetTodosNotFound() *GetTodosNotFound {
	return &GetTodosNotFound{}
}

// WithPayload adds the payload to the get todos not found response
func (o *GetTodosNotFound) WithPayload(payload GetTodosNotFoundBody) *GetTodosNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todos not found response
func (o *GetTodosNotFound) SetPayload(payload GetTodosNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodosNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetTodosInternalServerErrorCode is the HTTP code returned for type GetTodosInternalServerError
const GetTodosInternalServerErrorCode int = 500

/*GetTodosInternalServerError get todos internal server error

swagger:response getTodosInternalServerError
*/
type GetTodosInternalServerError struct {

	/*
	  In: Body
	*/
	Payload GetTodosInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetTodosInternalServerError creates GetTodosInternalServerError with default headers values
func NewGetTodosInternalServerError() *GetTodosInternalServerError {
	return &GetTodosInternalServerError{}
}

// WithPayload adds the payload to the get todos internal server error response
func (o *GetTodosInternalServerError) WithPayload(payload GetTodosInternalServerErrorBody) *GetTodosInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todos internal server error response
func (o *GetTodosInternalServerError) SetPayload(payload GetTodosInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodosInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
