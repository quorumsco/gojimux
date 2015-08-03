/*
Package application provides a convenient wrapper for common web
applications components.
*/
package gojimux

import (
	"net/http"

	"github/quorumsco/application"
)

func (r *Router) Get(path interface{}, handle http.HandlerFunc) {
	r.Handle("GET", path, handle)
}

func (r *Router) Post(path interface{}, handle http.HandlerFunc) {
	r.Handle("POST", path, handle)
}

func (r *Router) Put(path interface{}, handle http.HandlerFunc) {
	r.Handle("PUT", path, handle)
}

func (r *Router) Patch(path interface{}, handle http.HandlerFunc) {
	r.Handle("PATCH", path, handle)
}

func (r *Router) Delete(path interface{}, handle http.HandlerFunc) {
	r.Handle("DELETE", path, handle)
}

func (r *Router) Options(path interface{}, handle http.HandlerFunc) {
	r.Handle("OPTIONS", path, handle)
}

func (r *Router) Use(handler func(http.Handler) http.Handler) {
	app.Components["Mux"].(Mux).Use(handler)
}
