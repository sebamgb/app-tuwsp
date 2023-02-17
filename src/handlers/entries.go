package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"app/src/models"
	"app/src/requests"
	"app/src/server"
)

// GetLogin shows login page with form
func GetLogin(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		form, err := requests.GetFormById(s.Config().Login.FormId)
		if err != nil {
			fmt.Printf("error in get form: %v", err)
			return
		}
		labelsRender, err := requests.GetKeyValueLabelsById(form.KeyValue)
		if err != nil {
			fmt.Printf("error in get key value: %v", err)
			return
		}
		placeholdersRender, err := requests.GetKeyValuePlaceholdersByLabelId(labelsRender.Id)
		if err != nil {
			fmt.Printf("error in get key value: %v", err)
			return
		}
		renderize := &models.Renderize{
			Url:         s.Config().Login.Url,
			Method:      s.Config().Login.Method,
			Title:       labelsRender.Title,
			Label:       labelsRender,
			Placeholder: placeholdersRender,
		}
		var render = models.Render{
			// target as title
			Target: strings.Title(s.Config().Login.Title),
			// app in upper case
			App:       strings.ToUpper(form.App),
			Renderize: renderize,
		}
		err = s.Config().Template.ExecuteTemplate(w, "base", &render)
		if err != nil {
			log.Fatal(fmt.Sprintf("error in execute template: %v", err))
		}
	}
}

// PostLogin validate login request
func PostLogin(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse form
		err := r.ParseForm()
		if err != nil {
			fmt.Printf("error in parse form: %v", err)
			return
		}
		// set login credentials
		s.Config().Login.Email = r.Form.Get("email")
		s.Config().Login.Password = r.Form.Get("password")
		// id from header
		authId := w.Header().Get("X-Auth-Id")
		if authId != "" {
			s.Config().Login.AuthId = authId
		}
		// create login
		response, err := requests.PostLoginAuth(s.Config().Login)
		if err != nil {
			fmt.Printf("err creating login: %v\n", err)
			return
		}
		// success true then...
		if response.Success {
			// register login creating in database
			resp, err := requests.PostLoginCreate(s.Config().Login, response.Token)
			if err != nil {
				fmt.Printf("err registing login: %v\n", err)
				return
			}
			// login id to header
			w.Header().Set("X-Login-Id", resp.Id)
			// message to header
			w.Header().Set("X-Successfully-Message", resp.Message)
			// token to header
			w.Header().Set("Authorization", response.Token)
			// redirect to dashboard
			http.Redirect(w, r, "/dashboard", http.StatusFound)
			return
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

// Unauthorized shows unauthorized page
func Unauthorized(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var render = models.Render{
			Target: "Alto ahí!",
			Error:  "No estas autorizado",
		}
		err := s.Config().Template.ExecuteTemplate(w, "base", &render)
		if err != nil {
			log.Fatal(fmt.Sprintf("error in execute template: %v", err))
		}
	}
}
