package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"app/src/models"
)

// PostValidateLogin do a request to CRUD to validate a login
func PostValidateLogin(req *models.Login) (string, bool) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return "", false
	}
	// Create Request
	resp, err := http.Post("http://localhost:3000/validate", "application/json",
		bytes.NewReader(requestBody))
	if err != nil {
		return "", false
	}
	// closing body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	// translate response to bytes
	bytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", false
	}
	// translate bytes to struct
	response := &models.ValidateResponse{}
	err = json.Unmarshal(bytesResp, response)
	if err != nil {
		return "", false
	}
	return response.Message, response.Success
}

// PostLoginCreate do a request to CRUD to create a login
func PostLoginCreate(req *models.Login, token string) ([]byte, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// Create Request adding token to header
	reqResp, err := http.NewRequest("POST", "http://localhost:3000/api/v1/logins", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}
	reqResp.Header.Add("Authorization", token)
	reqResp.Header.Add("Content-Type", "application/json")
	// do request
	resp, err := http.DefaultClient.Do(reqResp)
	if err != nil {
		return nil, err
	}
	// closing body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	// translate response to bytes
	bytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytesResp, nil
}

// PostLoginAuth do a request post to CRUD to log in an auth
func PostLoginAuth(req *models.Login) (*models.LoginResponse, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// Create Request
	resp, err := http.Post("http://localhost:3000/login", "application/json",
		bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}
	// closing body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	// response
	response := &models.LoginResponse{}
	// translate response to bytes
	bytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// translate bytes to struct
	err = json.Unmarshal(bytesResp, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// PostSignupAuth do a request post to CRUD to sign up an auth
func PostSignupAuth(req *models.Signup) ([]byte, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// Create Request
	resp, err := http.Post("http://localhost:3000/signup", "application/json",
		bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}
	// closing body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	// translate response to bytes
	bytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytesResp, nil
}

// PostFormCreate do a request post to CRUD to create a form
func PostFormCreate(req *models.Form, token string) ([]byte, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// Create Request adding token to header
	reqResp, err := http.NewRequest("POST", "http://localhost:3000/api/v1/forms", bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}
	reqResp.Header.Add("Authorization", token)
	reqResp.Header.Add("Content-Type", "application/json")
	// do request
	resp, err := http.DefaultClient.Do(reqResp)
	if err != nil {
		return nil, err
	}
	// closing body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	// translate response to bytes
	bytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytesResp, nil
}
