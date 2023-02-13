package models

import "github.com/golang-sql/civil"

type Dashboard struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Menu  string `json:"key_value"`
	App   string `json:"app"`
	Owner string `json:"owner"`
}

type Login struct {
	Id        string         `json:"id"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt civil.DateTime `json:"created_at"`
	LogOut    civil.DateTime `json:"log_out"`
	AuthId    string         `json:"auth_id"`
	FormId    string         `json:"form_id"`
}

type Signup struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	NickName        string `json:"nick_name"`
	Email           string `json:"email"`
	Phone           int    `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	FormId          string `json:"form_id"`
}

type Auth struct {
	Id        string         `json:"id"`
	Email     string         `json:"email"`
	CreatedAt civil.DateTime `json:"created_at"`
	Password  string         `json:"password"`
	SignupId  string         `json:"signup_id"`
}

type Form struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	App       string `json:"app"`
	Key_Value string `json:"key_value"`
}

type KeyValue struct {
	Id                         string `json:"id"`
	Error                      string `json:"error"` // Error message
	Title                      string `json:"title"` // Title of the form
	LabelName                  string `json:"label_name"`
	PlaceholderName            string `json:"placeholder_name"`
	LabelNickname              string `json:"label_nickname"`
	PlaceholderNickname        string `json:"placeholder_nickname"`
	LabelEmail                 string `json:"label_email"`
	PlaceholderEmail           string `json:"placeholder_email"`
	LabelPhone                 string `json:"label_phone"`
	PlaceholderPhone           string `json:"placeholder_phone"`
	LabelPassword              string `json:"label_password"`
	PlaceholderPassword        string `json:"placeholder_password"`
	LabelConfirmPassword       string `json:"label_confirm_password"`
	PlaceholderConfirmPassword string `json:"placeholder_confirm_password"`
	InputSubmit                string `json:"input_submit"`
}
