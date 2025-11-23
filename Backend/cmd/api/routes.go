package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (a *applicationDependencies) routes() http.Handler {
	router := httprouter.New()

	// Disable httprouter's automatic OPTIONS handling.
    router.HandleOPTIONS = false

	// Custom error handlers
	router.NotFound = http.HandlerFunc(a.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(a.methodNotAllowedResponse)

	// HealthCheck handler
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", a.healthCheckHandler)

	// User routes (some public, some protected)
	router.HandlerFunc(http.MethodPost, "/v1/users", a.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", a.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", a.createAuthenticationTokenHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/users/:id", a.requireActivatedUser(a.updateUserHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/users/:id", a.requireActivatedUser(a.deleteUserHandler))
	
	return a.recoverPanic(a.enableCORS(a.rateLimit(a.authenticate(router))))
}