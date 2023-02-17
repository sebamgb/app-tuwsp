package routes

import (
	"app/src/handlers"
	"app/src/middlewares"
	"app/src/server"

	"net/http"

	"github.com/gorilla/mux"
)

// BindRoute implement Router mux of gorilla for call handlers with its path
func BindRoute(s server.Server, r *mux.Router) {
	r.Path("/").Handler(handlers.HomeHandler(s))
	// router static files
	static := r.NewRoute().PathPrefix("/assets/")
	// router for all client
	public := r.NewRoute().Subrouter()
	// router for validate login by client
	private := r.NewRoute().Subrouter()
	// routes defined
	private.Use(middlewares.ValidateMiddleware(s))
	private.Handle("/login", handlers.PostLogin(s)).Methods(http.MethodPost)
	public.Handle("/login", handlers.GetLogin(s)).Methods(http.MethodGet)
	public.Handle("/signup", handlers.Signup(s)).Methods(http.MethodGet)
	// public.Handle("/signup-processor", handlers.PostSignup(s)).Methods(http.MethodPost)
	public.Handle("/unauthorized", handlers.Unauthorized(s))
	// handle dir for static files
	static.Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("public/assets")))).
		Methods(http.MethodGet)
}
