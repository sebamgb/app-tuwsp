package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"app/src/models"
)

// PostLoginCreate do a request to CRUD to create a login
func PostLoginCreate(req *models.Login) (bool, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return false, err
	}
	// Create Request
	resp, err := http.Post("http://localhost:3000/create/login", "application/json",
		bytes.NewReader(requestBody))
	if err != nil {
		return false, err
	}
	// closing body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()
	return true, nil
}

// PostLoginAuth do a request post to CRUD to log in an auth
func PostLoginAuth(req *models.Login) (string, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	// Create Request
	resp, err := http.Post("http://localhost:3000/login", "application/json",
		bytes.NewReader(requestBody))
	if err != nil {
		return "", err
	}
	// closing body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()
	// response
	response := &models.LoginResponse{}
	// read body
	bytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// unmarshaL body
	if err = json.Unmarshal(bytesResp, response); err != nil {
		return "", err
	}
	return response.Token, nil
}

// PostSignupAuth do a request post to CRUD to sign up an auth
func PostSignupAuth(req *models.Signup) (bool, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return false, err
	}
	// Create Request
	resp, err := http.Post("http://localhost:3000/signup", "application/json",
		bytes.NewReader(requestBody))
	if err != nil {
		return false, err
	}
	// closing body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()
	return true, nil
}
