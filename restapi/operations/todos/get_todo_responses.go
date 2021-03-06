// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"span/go-swagger-example/models"
)

// GetTodoOKCode is the HTTP code returned for type GetTodoOK
const GetTodoOKCode int = 200

/*GetTodoOK get todo o k

swagger:response getTodoOK
*/
type GetTodoOK struct {

	/*
	  In: Body
	*/
	Payload *models.TodoFull `json:"body,omitempty"`
}

// NewGetTodoOK creates GetTodoOK with default headers values
func NewGetTodoOK() *GetTodoOK {
	return &GetTodoOK{}
}

// WithPayload adds the payload to the get todo o k response
func (o *GetTodoOK) WithPayload(payload *models.TodoFull) *GetTodoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todo o k response
func (o *GetTodoOK) SetPayload(payload *models.TodoFull) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTodoUnauthorizedCode is the HTTP code returned for type GetTodoUnauthorized
const GetTodoUnauthorizedCode int = 401

/*GetTodoUnauthorized get todo unauthorized

swagger:response getTodoUnauthorized
*/
type GetTodoUnauthorized struct {

	/*
	  In: Body
	*/
	Payload GetTodoUnauthorizedBody `json:"body,omitempty"`
}

// NewGetTodoUnauthorized creates GetTodoUnauthorized with default headers values
func NewGetTodoUnauthorized() *GetTodoUnauthorized {
	return &GetTodoUnauthorized{}
}

// WithPayload adds the payload to the get todo unauthorized response
func (o *GetTodoUnauthorized) WithPayload(payload GetTodoUnauthorizedBody) *GetTodoUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todo unauthorized response
func (o *GetTodoUnauthorized) SetPayload(payload GetTodoUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodoUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetTodoForbiddenCode is the HTTP code returned for type GetTodoForbidden
const GetTodoForbiddenCode int = 403

/*GetTodoForbidden get todo forbidden

swagger:response getTodoForbidden
*/
type GetTodoForbidden struct {

	/*
	  In: Body
	*/
	Payload GetTodoForbiddenBody `json:"body,omitempty"`
}

// NewGetTodoForbidden creates GetTodoForbidden with default headers values
func NewGetTodoForbidden() *GetTodoForbidden {
	return &GetTodoForbidden{}
}

// WithPayload adds the payload to the get todo forbidden response
func (o *GetTodoForbidden) WithPayload(payload GetTodoForbiddenBody) *GetTodoForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todo forbidden response
func (o *GetTodoForbidden) SetPayload(payload GetTodoForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodoForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetTodoNotFoundCode is the HTTP code returned for type GetTodoNotFound
const GetTodoNotFoundCode int = 404

/*GetTodoNotFound get todo not found

swagger:response getTodoNotFound
*/
type GetTodoNotFound struct {

	/*
	  In: Body
	*/
	Payload GetTodoNotFoundBody `json:"body,omitempty"`
}

// NewGetTodoNotFound creates GetTodoNotFound with default headers values
func NewGetTodoNotFound() *GetTodoNotFound {
	return &GetTodoNotFound{}
}

// WithPayload adds the payload to the get todo not found response
func (o *GetTodoNotFound) WithPayload(payload GetTodoNotFoundBody) *GetTodoNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todo not found response
func (o *GetTodoNotFound) SetPayload(payload GetTodoNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodoNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetTodoInternalServerErrorCode is the HTTP code returned for type GetTodoInternalServerError
const GetTodoInternalServerErrorCode int = 500

/*GetTodoInternalServerError get todo internal server error

swagger:response getTodoInternalServerError
*/
type GetTodoInternalServerError struct {

	/*
	  In: Body
	*/
	Payload GetTodoInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetTodoInternalServerError creates GetTodoInternalServerError with default headers values
func NewGetTodoInternalServerError() *GetTodoInternalServerError {
	return &GetTodoInternalServerError{}
}

// WithPayload adds the payload to the get todo internal server error response
func (o *GetTodoInternalServerError) WithPayload(payload GetTodoInternalServerErrorBody) *GetTodoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todo internal server error response
func (o *GetTodoInternalServerError) SetPayload(payload GetTodoInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
