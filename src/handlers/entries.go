package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"app/src/models"
	"app/src/requests"
	"app/src/server"

	"github.com/golang-sql/civil"
)

// Login log in user from login form
func Login(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.Login{}
		if r.Method == "POST" {
			request.Password = r.FormValue("password")
			request.Email = r.FormValue("email")
			response, err := requests.PostLoginAuth(&request)
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return
			}
			request.CreatedAt = civil.DateTimeOf(time.Now())
			auth, err := requests.GetAuthByEmail(request.Email)
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return
			}
			request.AuthId = auth.Id
			w.Header().Set("Authorization", response)
		}
		var render = models.Render{
			Target:    request.FormId,
			Renderize: &models.LoginLabels{},
		}
		err := s.Config().Template.ExecuteTemplate(w, "base", &render)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Signup(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := models.Signup{}
		if r.Method == "POST" {

		}
		var render = models.Render{
			Target:    r.RequestURI,
			Renderize: &SignupLabels{},
		}
		s.Config().Template.ExecuteTemplate(w, "signup", &render)
	}
}
