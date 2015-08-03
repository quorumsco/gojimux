package gojimux

import (
	"net/http"
	"strconv"

	"github.com/quorumsco/jsonapi"
	"github.com/quorumsco/logs"
	"github.com/quorumsco/router"
)

func (app Gojimux) Cors(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin,content-type")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (app Gojimux) SetUID(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var (
			res    int
			userID uint
			err    error

			query = r.URL.Query()
		)
		res, err = strconv.Atoi(query.Get("user_id"))
		if err != nil {
			logs.Debug(err)
			jsonapi.Error(w, r, err.Error(), http.StatusBadRequest)
			return
		}
		userID = uint(res)
		router.Context(r).Env["UserID"] = userID
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (app *Gojimux) Apply(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		router.Context(r).Env["Application"] = app
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
