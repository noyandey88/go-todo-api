package routes

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func RegisterSwagger(mux *http.ServeMux) {
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
}
