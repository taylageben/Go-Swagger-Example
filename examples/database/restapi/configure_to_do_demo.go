package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"


	"span/go-swagger-example/examples/database/restapi/operations"
	"span/go-swagger-example/examples/database/restapi/operations/misc"
	"span/go-swagger-example/examples/database/restapi/operations/todos"
	"span/go-swagger-example/examples/database/restapi/dbHandler"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target ../examples/database --name  --spec ../swagger.yml

func configureFlags(api *operations.ToDoDemoAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ToDoDemoAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	var todoStore dbHandler.TodoStore

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "apikey" query is set
	api.ApikeyAuth = func(token string) (interface{}, error) {
		if token == "123" {
			return token, nil
		} else {
			return nil, errors.New(401, "incorrect api key auth")
		}
	}

	api.TodosDeleteTodoHandler = todos.DeleteTodoHandlerFunc(func(params todos.DeleteTodoParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation todos.DeleteTodo has not yet been implemented")
	})
	api.MiscGetMockedHandler = misc.GetMockedHandlerFunc(func(params misc.GetMockedParams) middleware.Responder {
		return middleware.NotImplemented("operation misc.GetMocked has not yet been implemented")
	})
	api.TodosGetTodoHandler = todos.GetTodoHandlerFunc(func(params todos.GetTodoParams) middleware.Responder {
		if params.TodoID == "" {
			msg := "Empty Body in Request"
			return todos.NewGetTodoForbidden().WithPayload(todos.GetTodoForbiddenBody{Message: &msg})
		}
		if success, payload := todoStore.GetTodo(params.TodoID); success {
			return todos.NewGetTodoOK().WithPayload(payload)
		}
		return todos.NewGetTodoInternalServerError().WithPayload("Something went wrong")
	})
	api.TodosGetTodosHandler = todos.GetTodosHandlerFunc(func(params todos.GetTodosParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation todos.GetTodos has not yet been implemented")
	})
	api.TodosPostTodosHandler = todos.PostTodosHandlerFunc(func(params todos.PostTodosParams, principal interface{}) middleware.Responder {
		//this check is important otherwise if the request contains an empty body, the next method will break
		if params.Body == nil {
			return todos.NewPostTodosForbidden().WithPayload("Empty Body in Request")
		}
		if success, payload := todoStore.AddTodo(params.Body); success {
			return todos.NewPostTodosCreated().WithPayload(payload)
		}
		return todos.NewPostTodosInternalServerError().WithPayload("Something went wrong")
	})
	api.TodosPutTodosHandler = todos.PutTodosHandlerFunc(func(params todos.PutTodosParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation todos.PutTodos has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
