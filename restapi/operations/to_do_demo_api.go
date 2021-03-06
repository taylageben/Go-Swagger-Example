// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"span/go-swagger-example/restapi/operations/misc"
	"span/go-swagger-example/restapi/operations/todos"
)

// NewToDoDemoAPI creates a new ToDoDemo instance
func NewToDoDemoAPI(spec *loads.Document) *ToDoDemoAPI {
	return &ToDoDemoAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		TodosDeleteTodoHandler: todos.DeleteTodoHandlerFunc(func(params todos.DeleteTodoParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation TodosDeleteTodo has not yet been implemented")
		}),
		MiscGetMockedHandler: misc.GetMockedHandlerFunc(func(params misc.GetMockedParams) middleware.Responder {
			return middleware.NotImplemented("operation MiscGetMocked has not yet been implemented")
		}),
		TodosGetTodoHandler: todos.GetTodoHandlerFunc(func(params todos.GetTodoParams) middleware.Responder {
			return middleware.NotImplemented("operation TodosGetTodo has not yet been implemented")
		}),
		TodosGetTodosHandler: todos.GetTodosHandlerFunc(func(params todos.GetTodosParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation TodosGetTodos has not yet been implemented")
		}),
		TodosPostTodosHandler: todos.PostTodosHandlerFunc(func(params todos.PostTodosParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation TodosPostTodos has not yet been implemented")
		}),
		TodosPutTodosHandler: todos.PutTodosHandlerFunc(func(params todos.PutTodosParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation TodosPutTodos has not yet been implemented")
		}),

		// Applies when the "apikey" query is set
		ApikeyAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (apikey) apikey from query param [apikey] has not yet been implemented")
		},
	}
}

/*ToDoDemoAPI ## Welcome

This is a place to put general notes and extra information, for internal use.

To get started designing/documenting this API, select a version on the left. ## Welcome

This is a place to put general notes and extra information, for internal use.

To get started designing/documenting this API, select a version on the left. */
type ToDoDemoAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// ApikeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key apikey provided in the query
	ApikeyAuth func(string) (interface{}, error)

	// TodosDeleteTodoHandler sets the operation handler for the delete todo operation
	TodosDeleteTodoHandler todos.DeleteTodoHandler
	// MiscGetMockedHandler sets the operation handler for the get mocked operation
	MiscGetMockedHandler misc.GetMockedHandler
	// TodosGetTodoHandler sets the operation handler for the get todo operation
	TodosGetTodoHandler todos.GetTodoHandler
	// TodosGetTodosHandler sets the operation handler for the get todos operation
	TodosGetTodosHandler todos.GetTodosHandler
	// TodosPostTodosHandler sets the operation handler for the post todos operation
	TodosPostTodosHandler todos.PostTodosHandler
	// TodosPutTodosHandler sets the operation handler for the put todos operation
	TodosPutTodosHandler todos.PutTodosHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *ToDoDemoAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ToDoDemoAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ToDoDemoAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ToDoDemoAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ToDoDemoAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ToDoDemoAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ToDoDemoAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ToDoDemoAPI
func (o *ToDoDemoAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ApikeyAuth == nil {
		unregistered = append(unregistered, "ApikeyAuth")
	}

	if o.TodosDeleteTodoHandler == nil {
		unregistered = append(unregistered, "todos.DeleteTodoHandler")
	}

	if o.MiscGetMockedHandler == nil {
		unregistered = append(unregistered, "misc.GetMockedHandler")
	}

	if o.TodosGetTodoHandler == nil {
		unregistered = append(unregistered, "todos.GetTodoHandler")
	}

	if o.TodosGetTodosHandler == nil {
		unregistered = append(unregistered, "todos.GetTodosHandler")
	}

	if o.TodosPostTodosHandler == nil {
		unregistered = append(unregistered, "todos.PostTodosHandler")
	}

	if o.TodosPutTodosHandler == nil {
		unregistered = append(unregistered, "todos.PutTodosHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ToDoDemoAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ToDoDemoAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "apikey":

			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.ApikeyAuth)

		}
	}
	return result

}

// ConsumersFor gets the consumers for the specified media types
func (o *ToDoDemoAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *ToDoDemoAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ToDoDemoAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the to do demo API
func (o *ToDoDemoAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ToDoDemoAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/todos/{todoId}"] = todos.NewDeleteTodo(o.context, o.TodosDeleteTodoHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/mocked"] = misc.NewGetMocked(o.context, o.MiscGetMockedHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/todos/{todoId}"] = todos.NewGetTodo(o.context, o.TodosGetTodoHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/todos"] = todos.NewGetTodos(o.context, o.TodosGetTodosHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/todos"] = todos.NewPostTodos(o.context, o.TodosPostTodosHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/todos/{todoId}"] = todos.NewPutTodos(o.context, o.TodosPutTodosHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ToDoDemoAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *ToDoDemoAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
