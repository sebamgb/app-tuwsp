package models

type Renderize struct {
	Url           string    `json:"url"`
	Method        string    `json:"method"`
	Error         string    `json:"error"`
	Title         string    `json:"title"`
	ErrorEmail    string    `json:"error-email"`
	ErrorPassword string    `json:"error-pass"`
	Label         *KeyValue `json:"label"`
	Placeholder   *KeyValue `json:"placeholder"`
	SuccessSend   string    `json:"success-send"`
}

type Render struct {
	Target    string     `json:"title"`
	App       string     `json:"app"`
	Error     string     `json:"error"`
	Renderize *Renderize `json:"render"`
}

type SignupResponse struct {
	Success bool `json:"success"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

type CreateLoginResponse struct {
	Id      string `json:"id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ValidateResponse struct {
	Id      string `json:"id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type KeyValueResponse struct {
	Id      string `json:"id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

type FormResponse struct {
	Id      string `json:"id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Author  string `json:"author"`
}
