package models

type SignupResponse struct {
	Success bool `json:"success"`
}

type SignupLabels struct {
	LabelName  string
	LabelEmail string
}

type LoginLabels struct {
	LabelEmail    string
	LabelPassword string
	InputSubmit   string
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}
