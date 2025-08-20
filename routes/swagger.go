// package routes
//
// import (
// 	"net/http"
//
// 	httpSwagger "github.com/swaggo/http-swagger/v2"
// )
//
// func RegisterSwagger(mux *http.ServeMux) {
// 	mux.Handle("/swagger/", httpSwagger.WrapHandler)
// }

package routes

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func RegisterSwagger(mux *http.ServeMux) {
	// The httpSwagger.Handler function is a convenience function that
	// returns an http.Handler for the Swagger UI.
	mux.Handle("/swagger/", httpSwagger.Handler(
		// This option controls how the API documentation is displayed.
		// "none" - all endpoints are collapsed.
		// "list" - the endpoint list is expanded, but methods are collapsed.
		// "full" - all endpoints and methods are expanded.
		httpSwagger.DocExpansion("none"),
	))
}
