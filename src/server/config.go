package server

import (
	"app/src/models"

	"html/template"
	"os"
	"strings"
)

type Config struct {
	port     string
	Template *template.Template
	Login    *models.Login
	Signup   *models.Signup
}

// NewConfig builder of Config
func NewConfig() (*Config, error) {
	Port := os.Getenv("PORT")
	prefix := strings.Index(Port, ":")
	sb := strings.Builder{}
	sb.WriteString(":")
	sb.WriteString(Port)
	if prefix != 0 {
		Port = sb.String()
	}
	var globPath = "public/templates/*.html"
	return &Config{
		port: Port,
		Template: template.
			Must(template.ParseGlob(
				globPath,
			)),
		Login:  NewLogin(),
		Signup: NewSignup(),
	}, nil
}

func (c *Config) Port() string {
	return c.port
}

func NewLogin() *models.Login {
	return &models.Login{}
}

func NewSignup() *models.Signup {
	return &models.Signup{}
}
