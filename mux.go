//This package is a simple wrapper for the gorilla mux
package mux

import (
	"net/http"

	"github.com/gorilla/mux"
)

// This struct contains the main router and a map of subrouters
type Router struct {
	Main *mux.Router
	Subs map[string]*mux.Router
}

// This method returns a subrouter based on the provided prefix
func (router *Router) NewSubrouter(prefix string) *mux.Router {
	return router.Main.PathPrefix("/" + prefix).Subrouter()
}

// This function builds a new Router and return it
func NewRouter() *Router {
	return &Router{mux.NewRouter(), make(map[string]*mux.Router)}
}

// This middeware handles CORS
func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work (by calling the underneath ServeHTTP method
	r.Main.ServeHTTP(rw, req)
}
