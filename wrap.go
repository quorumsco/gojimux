/*
Package application provides a convenient wrapper for common web
applications components.
*/
package gojimux

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

var router = "goji"

func Get(path interface{}, handle http.HandlerFunc, router interface{}) {
	if router == "mux" {
		router.(Mux).Get(path, handle)
	} else if router == "goji" {
		router.(*web.Mux).Get(path, handle)
	}
}

func Post(path interface{}, handle http.HandlerFunc, router interface{}) {
	if router == "mux" {
		router.(Mux).Post(path, handle)
	} else if router == "goji" {
		router.(*web.Mux).Post(path, handle)
	}
}

func Put(path interface{}, handle http.HandlerFunc, router interface{}) {
	if router == "mux" {
		router.(Mux).Put(path, handle)
	} else if router == "goji" {
		router.(*web.Mux).Put(path, handle)
	}
}

func Patch(path interface{}, handle http.HandlerFunc, router interface{}) {
	if router == "mux" {
		router.(Mux).Patch(path, handle)
	} else if router == "goji" {
		router.(*web.Mux).Patch(path, handle)
	}
}

func Delete(path interface{}, handle http.HandlerFunc, router interface{}) {
	if router == "mux" {
		router.(Mux).Delete(path, handle)
	} else if router == "goji" {
		router.(*web.Mux).Delete(path, handle)
	}
}

func Options(path interface{}, handle http.HandlerFunc, router interface{}) {
	if router == "mux" {
		router.(Mux).Options(path, handle)
	} else if router == "goji" {
		router.(*web.Mux).Options(path, handle)
	}
}

func Use(handler func(http.Handler) http.Handler, router interface{}) {
	if router == "mux" {
		router.(Mux).Use(handler)
	} else if router == "goji" {
		router.(*web.Mux).Use(handler)
	}
}

func ServeHTTP(w http.ResponseWriter, req *http.Request, router interface{}) {
	if router == "mux" {
		router.(Mux).ServeHTTP(w, req)
	} else if router == "goji" {
		router.(*web.Mux).ServeHTTP(w, req)
	}
}

func ListenAndServe(listen string, router interface{}) error {
	if router == "mux" {
		return ListenAndServe(listen, router.(Mux))
	} else if router == "goji" {
		return ListenAndServe(listen, router.(*web.Mux))
	}
	return nil
}
