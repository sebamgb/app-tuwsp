package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"app/src/models"
)

// GetMeById do a request get to CRUD to get an auth by id
func GetMeById(q string, token string) (*models.Auth, error) {
	// parse of url to call
	u, err := url.Parse("http://localhost:3000/api-tuwsp/v1/me")
	if err != nil {
		return nil, fmt.Errorf("error parse url: %v", err)
	}
	// query param
	query := u.Query()
	// adding key value to query params
	query.Add("q", q) // q is id
	// encoding into u
	u.RawQuery = query.Encode()
	// request get with token
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error request get: %v", err)
	}
	// adding token to header
	req.Header.Add("Authorization", token)
	// client
	client := &http.Client{}
	// do request
	resp, err := client.Do(req)
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
	var auth = models.Auth{}
	if err = json.Unmarshal(bytesBody, &auth); err != nil {
		return nil, fmt.Errorf("error unmarshal: %s, %v", string(bytesBody), err)
	}
	return &auth, nil
}

// GetFormById do a request get to CRUD
func GetFormById(q string) (*models.Form, error) {
	// parse of url to call
	u, err := url.Parse("http://localhost:3000/form")
	if err != nil {
		return nil, fmt.Errorf("error parse url: %v", err)
	}
	// query param
	query := u.Query()
	// adding key value to query params
	query.Add("q", q) // q is title
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
		return nil, fmt.Errorf("error unmarshal: %s, %v", string(bytesBody), err)
	}
	return &form, nil
}

// GetKeyValueLabelsById do a request get to CRUD
func GetKeyValueLabelsById(q string) (*models.KeyValue, error) {
	// parse of url to call
	u, err := url.Parse("http://localhost:3000/key-value-labels")
	if err != nil {
		return nil, fmt.Errorf("error parse url: %v", err)
	}
	// query param
	query := u.Query()
	// adding key value to query params
	query.Add("q", q) // q is id
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
		return nil, fmt.Errorf("error unmarshal: %s, %v", string(bytesBody), err)
	}
	return &keyValue, nil
}

// GetKeyValuePlaceholdersByLabelId do a request get to CRUD
func GetKeyValuePlaceholdersByLabelId(q string) (*models.KeyValue, error) {
	// parse of url to call
	u, err := url.Parse("http://localhost:3000/key-value-placeholders")
	if err != nil {
		return nil, fmt.Errorf("error parse url: %v", err)
	}
	// query param
	query := u.Query()
	// adding key value to query params
	query.Add("q", q) // q is label id
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
		return nil, fmt.Errorf("error unmarshal: %s, %v", string(bytesBody), err)
	}
	return &keyValue, nil
}
