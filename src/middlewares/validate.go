package middlewares

import (
	"fmt"
	"net/http"

	"app/src/requests"
	"app/src/server"
)

// ValidateMiddleware validate request from form
func ValidateMiddleware(s server.Server) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// parse form
			err := r.ParseForm()
			if err != nil {
				fmt.Printf("error in parse form: %v", err)
				return
			}
			// validate request
			validRequest := requests.ValidateRequest{
				Email:    r.Form.Get("email"),
				Password: r.Form.Get("password"),
			}
			// validate
			valid, err := requests.PostValidateLogin(&validRequest)
			if err != nil {
				fmt.Printf("err validating login: %v\n", err)
				return
			}
			// success false then game over
			if !valid.Success {
				http.Redirect(w, r, "/unauthorized", http.StatusFound)
				return
			}
			authId := w.Header().Get("X-Auth-Id")
			if authId == "" {
				authId = valid.Id
			}
			// id to header
			w.Header().Set("X-Auth-Id", authId)
			next.ServeHTTP(w, r)
		})
	}
}
