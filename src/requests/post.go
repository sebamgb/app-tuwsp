package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"app/src/models"
)

type ValidateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// PostValidateLogin do a request to CRUD to validate a login
func PostValidateLogin(req *ValidateRequest) (*models.ValidateResponse, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return &models.ValidateResponse{
			Message: "ocurrió un error en la validación, por favor intenta más tarde.",
			Success: false,
		}, fmt.Errorf("error converting struct to bytes: %v", err)
	}
	// Create Request
	resp, err := http.Post("http://localhost:3000/validate", "application/json",
		bytes.NewReader(requestBody))
	if err != nil {
		return &models.ValidateResponse{
			Message: "ocurrió un error en la validación, por favor intenta más tarde.",
			Success: false,
		}, fmt.Errorf("error doing request: %v", err)
	}
	// closing body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(fmt.Errorf("error closing body: %v", err))
		}
	}()
	// translate response to bytes
	bytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &models.ValidateResponse{
			Message: "ocurrió un error en la validación, por favor intenta más tarde.",
			Success: false,
		}, fmt.Errorf("error reading response: %v", err)
	}
	// translate bytes to struct
	response := &models.ValidateResponse{}
	err = json.Unmarshal(bytesResp, response)
	if err != nil {
		return &models.ValidateResponse{
			Message: "ocurrió un error en la validación, por favor intenta más tarde.",
			Success: false,
		}, fmt.Errorf("error unmarshaling response: %s, %v", string(bytesResp), err)
	}
	return response, nil
}

// PostLoginCreate do a request to CRUD to create a login
func PostLoginCreate(req *models.Login, token string) (*models.CreateLoginResponse, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error converting struct to bytes: %v", err)
	}
	// Create Request adding token to header
	reqResp, err := http.NewRequest("POST", "http://localhost:3000/api-tuwsp/v1/logins", bytes.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	reqResp.Header.Add("Authorization", token)
	reqResp.Header.Add("Content-Type", "application/json")
	// do request
	resp, err := http.DefaultClient.Do(reqResp)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	// closing body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(fmt.Errorf("error closing body: %v", err))
		}
	}()
	// translate response to bytes
	bytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}
	// translate bytes to struct
	response := &models.CreateLoginResponse{}
	err = json.Unmarshal(bytesResp, response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal: %s, %v", string(bytesResp), err)
	}
	return response, nil
}

// PostLoginAuth do a request post to CRUD to log in an auth
func PostLoginAuth(req *models.Login) (*models.LoginResponse, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error converting struct to bytes: %v", err)
	}
	// Create Request
	resp, err := http.Post("http://localhost:3000/login", "application/json",
		bytes.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	// closing body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(fmt.Errorf("error closing body: %v", err))
		}
	}()
	// response
	response := &models.LoginResponse{}
	// translate response to bytes
	bytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}
	// translate bytes to struct
	err = json.Unmarshal(bytesResp, response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %s, %v", string(bytesResp), err)
	}
	return response, nil
}

// PostSignupAuth do a request post to CRUD to sign up an auth
func PostSignupAuth(req *models.Signup) ([]byte, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error converting struct to bytes: %v", err)
	}
	// Create Request
	resp, err := http.Post("http://localhost:3000/signup", "application/json",
		bytes.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	// closing body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(fmt.Errorf("error closing body: %v", err))
		}
	}()
	// translate response to bytes
	bytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %s, %v", string(bytesResp), err)
	}
	return bytesResp, nil
}

// PostFormCreate do a request post to CRUD to create a form
func PostFormCreate(req *models.Form) (*models.FormResponse, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error converting struct to bytes: %v", err)
	}
	// Create Request adding token to header
	reqResp, err := http.NewRequest("POST", "http://localhost:3000/forms", bytes.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	reqResp.Header.Add("Content-Type", "application/json")
	// do request
	resp, err := http.DefaultClient.Do(reqResp)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	// closing body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(fmt.Errorf("error closing body: %v", err))
		}
	}()
	// translate response to bytes
	bytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}
	// translate bytes to struct
	response := &models.FormResponse{}
	err = json.Unmarshal(bytesResp, response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal: %s, %v", string(bytesResp), err)
	}
	return response, nil
}

// PostKeyValueLabelsCreate do a request post to CRUD to create a key value labels
func PostKeyValueLabelsCreate(req *models.KeyValue) (*models.KeyValueResponse, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error converting struct to bytes: %v", err)
	}
	// Create Request adding token to header
	reqResp, err := http.NewRequest("POST", "http://localhost:3000/key-value-labels", bytes.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	reqResp.Header.Add("Content-Type", "application/json")
	// do request
	resp, err := http.DefaultClient.Do(reqResp)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	// closing body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(fmt.Errorf("error closing body: %v", err))
		}
	}()
	// translate response to bytes
	bytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}
	// translate bytes to struct
	response := &models.KeyValueResponse{}
	err = json.Unmarshal(bytesResp, response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %s, %v", string(bytesResp), err)
	}
	return response, nil
}

// PostKeyValuePlaceholdersCreate do a request post to CRUD to create a key value placeholders
func PostKeyValuePlaceholdersCreate(req *models.KeyValue) (*models.KeyValueResponse, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error converting struct to bytes: %v", err)
	}
	// Create Request adding token to header
	reqResp, err := http.NewRequest("POST", "http://localhost:3000/key-value-placeholders", bytes.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	reqResp.Header.Add("Content-Type", "application/json")
	// do request
	resp, err := http.DefaultClient.Do(reqResp)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %v", err)
	}
	// closing body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(fmt.Errorf("error closing body: %v", err))
		}
	}()
	// translate response to bytes
	bytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}
	// translate bytes to struct
	response := &models.KeyValueResponse{}
	err = json.Unmarshal(bytesResp, response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %s, %v", string(bytesResp), err)
	}
	return response, nil
}
