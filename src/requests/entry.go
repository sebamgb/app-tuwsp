package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"app/src/models"
)

func Entrypoint(login *models.Login, signup *models.Signup) error {
	// Reading files json templates_key_value.json
	bytes_d, err := ioutil.ReadFile("src/models/json/templates_key_value.json")
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	// translate template json to keyvalues
	keyvalues := []*models.KeyValue{}
	if err := json.Unmarshal(bytes_d, &keyvalues); err != nil {
		return fmt.Errorf("error translating template json to keyvalues: %v", err)
	}
	// Reading files json templates_form.json
	bytes_d, err = ioutil.ReadFile("src/models/json/templates_form.json")
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	// translate template json to form
	forms := []*models.Form{}
	if err := json.Unmarshal(bytes_d, &forms); err != nil {
		return fmt.Errorf("error translating template json to forms: %v", err)
	}
	// Reading files json login_template.json
	bytes_d, err = ioutil.ReadFile("src/models/json/login_template.json")
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	// translate template json to login
	if err := json.Unmarshal(bytes_d, &login); err != nil {
		return fmt.Errorf("error translating template json to login: %v", err)
	}
	// Reading files json signup_template.json
	bytes_d, err = ioutil.ReadFile("src/models/json/signup_template.json")
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	// translate template json to signup
	if err := json.Unmarshal(bytes_d, &signup); err != nil {
		return fmt.Errorf("error translating template json to signup: %v", err)
	}
	var count int = 1
	// range all keyvalues and create
	for i := range forms {
		// create key value and get response
		respKeyValue, err := PostKeyValueLabelsCreate(keyvalues[i])
		if err != nil {
			return fmt.Errorf("error creating key value labels: %v", err)
		}
		if respKeyValue.Success {
			forms[i].KeyValue = respKeyValue.Id
			keyvalues[i].LabelId = respKeyValue.Id
		}
		respKeyValue, err = PostKeyValuePlaceholdersCreate(keyvalues[i])
		if err != nil {
			return fmt.Errorf("error creating key value placeholders: %v", err)
		}
		log.Printf("%v by %v", respKeyValue.Message, respKeyValue.Author)
		// create form and get response
		resp, err := PostFormCreate(forms[i])
		if err != nil {
			return fmt.Errorf("error creating form: %v", err)
		}
		log.Printf("%v by %v", resp.Message, resp.Author)
		if count == 1 {
			// compare login and signup
			compareLogin := strings.ToLower(fmt.Sprintf("%T", login))
			compareSignup := strings.ToLower(fmt.Sprintf("%T", signup))
			if resp.Success {
				// if compare contains login form target
				if strings.Contains(compareLogin, forms[i].Target) {
					login.FormId = resp.Id
				}
				// if compare contains signup form target
				if strings.Contains(compareSignup, forms[i].Target) {
					signup.FormId = resp.Id
				}
			}
		} else {
			continue
		}
		count++
	}
	return nil
}
