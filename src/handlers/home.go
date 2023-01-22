package handlers

import (
	"app/src/server"
	"encoding/json"
	"fmt"

	"net/http"
)

type response struct {
	Title []string `url:"title"`
}

// HomeHandler handle the root path
func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := response{}
		err := r.ParseForm()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		bytes, _ := json.Marshal(r.PostForm)
		_ = json.Unmarshal(bytes, &res)
		fmt.Printf("r.PostForm: %v\n", r.PostForm)
		if err := s.
			Config().Template.ExecuteTemplate(
			w,
			"base",
			res,
		); err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
	}
}
