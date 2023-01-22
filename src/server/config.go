package server

import (
	"html/template"
	"os"
	"strings"
)

type Config struct {
	port     string
	Template *template.Template
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
	}, nil
}
