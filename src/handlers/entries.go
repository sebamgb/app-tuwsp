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
		if r.Method == "POST" {
			request := models.Login{
				Email:    r.FormValue("email"),
				Password: r.FormValue("password"),
			}
			// validate request
			message, ok := requests.PostValidateLogin(&request)
			if !ok {
				w.Header().Set("X-Error-Message", message)
				// redirect login with method GET
				req := http.Request{
					Method: "GET",
				}
				Login(s)(w, &req)
				return
			}
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
			// success true then inyect token to header0
			if response.Success {
				w.Header().Set("Authorization", response.Token)
			}
		} else {
			form, err := requests.GetForm(r.RequestURI[1:])
			if err != nil {
				fmt.Printf("error in get form: %v", err)
				return
			}
			fmt.Printf("form: %v", form.Key_Value)
			renderize, err := requests.GetKeyValue(form.Key_Value)
			if err != nil {
				fmt.Printf("error in get key value: %v", err)
				return
			}
			var render = models.Render{
				// target as title
				Target:    form.Title,
				Renderize: renderize,
			}
			err = s.Config().Template.ExecuteTemplate(w, "base", render)
			if err != nil {
				log.Fatal(fmt.Sprintf("error in execute template: %v", err))
			}
		}
	}
}

func Signup(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request := models.Signup{}
		if r.Method == "POST" {

		}
		var render = models.Render{
			Target: r.RequestURI,
			// Renderize: &SignupLabels{},
		}
		s.Config().Template.ExecuteTemplate(w, "signup", &render)
	}
}
