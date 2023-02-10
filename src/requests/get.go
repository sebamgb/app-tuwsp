package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"app/src/models"
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

// GetForm do a request get to CRUD
func GetForm(q string) (*models.Form, error) {
	// parse of url to call
	u, err := url.Parse("http://localhost:3000/form")
	if err != nil {
		return nil, fmt.Errorf("error parse url: %v", err)
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
		return nil, fmt.Errorf("error request get: %v", err)
	}
	// closing the body
	defer resp.Body.Close()
	// reading body
	bytesBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %v", err)
	}
	// unmarshal to struct
	var form = models.Form{}
	if err = json.Unmarshal(bytesBody, &form); err != nil {
		return nil, fmt.Errorf("error unmarshal: %v", err)
	}
	return &form, nil
}

// GetKeyValue do a request get to CRUD
func GetKeyValue(q string) (*models.KeyValue, error) {
	// parse of url to call
	u, err := url.Parse("http://localhost:3000/key-value")
	if err != nil {
		return nil, fmt.Errorf("error parse url: %v", err)
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
		return nil, fmt.Errorf("error request get: %v", err)
	}
	// closing the body
	defer resp.Body.Close()
	// reading body
	bytesBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %v", err)
	}
	// unmarshal to struct
	var keyValue = models.KeyValue{}
	if err = json.Unmarshal(bytesBody, &keyValue); err != nil {
		return nil, fmt.Errorf("error unmarshal: %v", err)
	}
	return &keyValue, nil
}
