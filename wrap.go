package gojimux

import (
	"net/http"

	"github.com/quorumsco/logs"
	"github.com/zenazn/goji/web"
)

type Gojimux struct {
	Components map[string]interface{}
}

func New() *Gojimux {
	var app = new(Gojimux)
	app.Components = make(map[string]interface{})
	return app
}

func (app Gojimux) Get(path interface{}, handle http.HandlerFunc) {
	app.Components["Mux"].(*web.Mux).Get(path, handle)
}

func (app Gojimux) Post(path interface{}, handle http.HandlerFunc) {
	app.Components["Mux"].(*web.Mux).Post(path, handle)
}

func (app Gojimux) Put(path interface{}, handle http.HandlerFunc) {
	app.Components["Mux"].(*web.Mux).Put(path, handle)
}

func (app Gojimux) Patch(path interface{}, handle http.HandlerFunc) {
	app.Components["Mux"].(*web.Mux).Patch(path, handle)
}

func (app Gojimux) Delete(path interface{}, handle http.HandlerFunc) {
	app.Components["Mux"].(*web.Mux).Delete(path, handle)
}

func (app Gojimux) Options(path interface{}, handle http.HandlerFunc) {
	app.Components["Mux"].(*web.Mux).Options(path, handle)
}

func (app Gojimux) Use(handler func(http.Handler) http.Handler) {
	app.Components["Mux"].(*web.Mux).Use(handler)
}

func (app Gojimux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	app.Components["Mux"].(*web.Mux).ServeHTTP(w, req)
}

func (app Gojimux) Serve(listen string) error {
	logs.Info("listening on http://%s", listen)
	return http.ListenAndServe(listen, app.Components["Mux"].(*web.Mux))
}
