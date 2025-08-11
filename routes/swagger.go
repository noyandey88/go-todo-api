package routes

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterSwagger(mux *http.ServeMux) {
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
}
