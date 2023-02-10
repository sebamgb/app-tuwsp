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

// GetForm do a request get to CRUD
func GetForm(q string) (*models.Form, error) {
	// parse of url to call
	u, err := url.Parse("http://localhost:3000/form")
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
	var form = models.Form{}
	if err = json.Unmarshal(bytesBody, &form); err != nil {
		return nil, err
	}
	return &form, nil
}

// GetKeyValue do a request get to CRUD
func GetKeyValue(q string) (*models.KeyValue, error) {
	// parse of url to call
	u, err := url.Parse("http://localhost:3000/key-values")
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
	var keyValue = models.KeyValue{}
	if err = json.Unmarshal(bytesBody, &keyValue); err != nil {
		return nil, err
	}
	return &keyValue, nil
}
