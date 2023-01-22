package routes

import (
	"app/src/handlers"
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
	// routes defined
	public.Handle("/login", handlers.Login(s))
	public.Handle("/signup", handlers.Signup(s))
	// handle dir for this route
	static.Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("public/assets"))))
	// private := r.NewRoute().Subrouter()
}
