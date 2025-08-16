package utils

import (
	"fmt"
	"net/http"
)

type MuxRouterInterface interface {
	Get(path string, handler http.Handler)
	Post(path string, handler http.Handler)
	Put(path string, handler http.Handler)
	Delete(path string, handler http.Handler)
}

type muxRouter struct {
	mux *http.ServeMux
}

func NewMuxRouter(mux *http.ServeMux) MuxRouterInterface {
	return &muxRouter{mux: mux}
}

func pathFormat(method string, path string) string {
	newPath := fmt.Sprintf("%s %s", method, path)
	return newPath
}

func (r *muxRouter) Get(path string, handler http.Handler) {
	r.mux.Handle(pathFormat("GET", path), handler)
}

func (r *muxRouter) Post(path string, handler http.Handler) {
	r.mux.Handle(pathFormat("POST", path), handler)
}

func (r *muxRouter) Put(path string, handler http.Handler) {
	r.mux.Handle(pathFormat("PUT", path), handler)
}

func (r *muxRouter) Delete(path string, handler http.Handler) {
	r.mux.Handle(pathFormat("DELETE", path), handler)
}
