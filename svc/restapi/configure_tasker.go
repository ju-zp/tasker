// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ju-zp/tasker/svc/restapi/operations"
)

//go:generate swagger generate server --target ../../svc --name Tasker --spec ../../swagger.yml --exclude-main

func configureFlags(api *operations.TaskerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TaskerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CreateProjectHandler == nil {
		api.CreateProjectHandler = operations.CreateProjectHandlerFunc(func(params operations.CreateProjectParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateProject has not yet been implemented")
		})
	}
	if api.CreateTaskHandler == nil {
		api.CreateTaskHandler = operations.CreateTaskHandlerFunc(func(params operations.CreateTaskParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateTask has not yet been implemented")
		})
	}
	if api.CreateTaskTodoHandler == nil {
		api.CreateTaskTodoHandler = operations.CreateTaskTodoHandlerFunc(func(params operations.CreateTaskTodoParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateTaskTodo has not yet been implemented")
		})
	}
	if api.DeleteTaskHandler == nil {
		api.DeleteTaskHandler = operations.DeleteTaskHandlerFunc(func(params operations.DeleteTaskParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteTask has not yet been implemented")
		})
	}
	if api.DeleteTodoHandler == nil {
		api.DeleteTodoHandler = operations.DeleteTodoHandlerFunc(func(params operations.DeleteTodoParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteTodo has not yet been implemented")
		})
	}
	if api.GetPingHandler == nil {
		api.GetPingHandler = operations.GetPingHandlerFunc(func(params operations.GetPingParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetPing has not yet been implemented")
		})
	}
	if api.GetProjectsHandler == nil {
		api.GetProjectsHandler = operations.GetProjectsHandlerFunc(func(params operations.GetProjectsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetProjects has not yet been implemented")
		})
	}
	if api.GetTaskTodosHandler == nil {
		api.GetTaskTodosHandler = operations.GetTaskTodosHandlerFunc(func(params operations.GetTaskTodosParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetTaskTodos has not yet been implemented")
		})
	}
	if api.GetTasksHandler == nil {
		api.GetTasksHandler = operations.GetTasksHandlerFunc(func(params operations.GetTasksParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetTasks has not yet been implemented")
		})
	}
	if api.SetTodoStatusHandler == nil {
		api.SetTodoStatusHandler = operations.SetTodoStatusHandlerFunc(func(params operations.SetTodoStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.SetTodoStatus has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

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
func configureServer(s *http.Server, scheme, addr string) {
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
