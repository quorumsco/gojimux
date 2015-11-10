// Implement a wrap around goji to satisty mux
package gojimux

import (
	"net/http"

	"github.com/quorumsco/router"
	"github.com/zenazn/goji/web"
)

// Allows us to use the quorumsco/router context
func InitContext(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		router.SetContext(r)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
