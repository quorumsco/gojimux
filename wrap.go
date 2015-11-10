// Implement a wrap around goji to satisty mux
package gojimux

import (
	"net/http"

	"github.com/quorumsco/logs"
	"github.com/quorumsco/router"
	"github.com/zenazn/goji/web"
)

// Contains a web.Mux interface and the gojimux methods
type Gojimux struct {
	Mux *web.Mux
}

// Return a new gojimux wrap
func New() *Gojimux {
	var app = new(Gojimux)
	app.Mux = web.New()
	app.Mux.Use(InitContext)
	return app
}

// Wrap aroud goji for Get, calls the handle
func (app Gojimux) Get(path interface{}, handle http.HandlerFunc) {
	app.Mux.Get(path, app.putContext(handle))
}

// Wrap aroud goji for Post, calls the handle
func (app Gojimux) Post(path interface{}, handle http.HandlerFunc) {
	app.Mux.Post(path, app.putContext(handle))
}

// Wrap aroud goji for Put, calls the handle
func (app Gojimux) Put(path interface{}, handle http.HandlerFunc) {
	app.Mux.Put(path, handle)
}

// Wrap aroud goji for Patch, calls the handle
func (app Gojimux) Patch(path interface{}, handle http.HandlerFunc) {
	app.Mux.Patch(path, app.putContext(handle))
}

// Wrap aroud goji for Delete, calls the handle
func (app Gojimux) Delete(path interface{}, handle http.HandlerFunc) {
	app.Mux.Delete(path, app.putContext(handle))
}

// Wrap aroud goji for Options, calls the handle
func (app Gojimux) Options(path interface{}, handle http.HandlerFunc) {
	app.Mux.Options(path, app.putContext(handle))
}

// Wrap aroud goji for Use, calls the handle
func (app Gojimux) Use(handler func(http.Handler) http.Handler) {
	app.Mux.Use(handler)
}

// Wrap aroud goji for ServeHTTP
func (app Gojimux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	app.Mux.ServeHTTP(w, req)
}

// Wrap aroud goji for Serve
func (app Gojimux) Serve(listen string) error {
	logs.Info("listening on http://%s", listen)
	return http.ListenAndServe(listen, app.Mux)
}

// Middleware that convert the goji Context to a context usable by Mux
func (app Gojimux) putContext(handle http.HandlerFunc) func(web.C, http.ResponseWriter, *http.Request) {
	fn := func(c web.C, w http.ResponseWriter, r *http.Request) {
		param := new(router.Param)
		for name, value := range c.URLParams {
			param.Name = name
			param.Value = value
			router.Context(r).Params = append(router.Context(r).Params, *param)
		}
		handle(w, r)
	}
	return fn
}
