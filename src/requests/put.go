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

// PutLoginModel do a request put to CRUD to update the model login
func PutLoginModel(req *models.Login, token string) ([]byte, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error converting struct to bytes: %v", err)
	}
	// Create Request adding token to header
	reqResp, err := http.NewRequest("PUT", "http://localhost:3000/api-tuwsp/v1/login", bytes.NewReader(requestBody))
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
	return bytesResp, nil
}

// PutLogoutModel do a request put to CRUD to update the model login in logout
func PutLogoutModel(req *models.Login, token string) ([]byte, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error converting struct to bytes: %v", err)
	}
	// Create Request adding token to header
	reqResp, err := http.NewRequest("PUT", "http://localhost:3000/api-tuwsp/v1/logout", bytes.NewReader(requestBody))
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
	return bytesResp, nil
}

// PutSignupModel do a request put to CRUD to update the model signup
func PutSignupModel(req *models.Signup, token string) ([]byte, error) {
	// Converting struct to bytes
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error converting struct to bytes: %v", err)
	}
	// Create Request adding token to header
	reqResp, err := http.NewRequest("PUT", "http://localhost:3000/api-tuwsp/v1/signups", bytes.NewReader(requestBody))
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
	return bytesResp, nil
}
