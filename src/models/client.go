package models

type Render struct {
	Target    string `json:"title"`
	App       string `json:"app"`
	Renderize any    `json:"key_value"`
}

type SignupResponse struct {
	Success bool `json:"success"`
}

type SignupLabels struct {
	Error               string `json:"error"`
	Title               string `json:"title"`
	LabelName           string `json:"label_name"`
	PlaceholderName     string `json:"placeholder_name"`
	LabelNickname       string `json:"label_nickname"`
	PlaceholderNickname string `json:"placeholder_nickname"`
	LabelEmail          string `json:"label_email"`
	PlaceholderEmail    string `json:"placeholder_email"`
	LabelPhone          string `json:"label_phone"`
	PlaceholderPhone    string `json:"placeholder_phone"`
	LabelPassword       string `json:"label_password"`
	PlaceholderPassword string `json:"placeholder_password"`
	InputSubmit         string `json:"input_submit"`
}

type LoginLabels struct {
	Error               string `json:"error"`
	Title               string `json:"title"`
	LabelEmail          string `json:"label_email"`
	PlaceholderEmail    string `json:"placeholder_email"`
	LabelPassword       string `json:"label_password"`
	PlaceholderPassword string `json:"placeholder_password"`
	InputSubmit         string `json:"input_submit"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

type ValidateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
