package requests

import (
	"app/src/models"
	"encoding/json"
	"io/ioutil"

	"net/http"
	"net/url"
)

// GetAuthById do a request get to CRUD
func GetAuthById(q string) (*models.Auth, error) {
	// parse of url to call
	u, err := url.Parse("http://localhost:3000/auths")
	if err != nil {
		return nil, err
	}
	// query param
	query := u.Query()
	// adding key value to query params
	query.Add("q", q)
	// encoding into u
	u.RawQuery = query.Encode()
	// request get
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	// closing the body
	defer resp.Body.Close()
	// reading body
	bytesBody, err := ioutil.ReadAll(resp.Body)
	// unmarshal to struct
	var auth = models.Auth{}
	if err = json.Unmarshal(bytesBody, &auth); err != nil {
		return nil, err
	}
	return &auth, nil
}

// GetAuthByEmail do a request get to CRUD
func GetAuthByEmail(q string) (*models.Auth, error) {
	// parse of url to call
	u, err := url.Parse("http://localhost:3000/auths-by-email")
	if err != nil {
		return nil, err
	}
	// query param
	query := u.Query()
	// adding key value to query params
	query.Add("q", q)
	// encoding into u
	u.RawQuery = query.Encode()
	// request get
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	// closing the body
	defer resp.Body.Close()
	// reading body
	bytesBody, err := ioutil.ReadAll(resp.Body)
	// unmarshal to struct
	var auth = models.Auth{}
	if err = json.Unmarshal(bytesBody, &auth); err != nil {
		return nil, err
	}
	return &auth, nil
}
