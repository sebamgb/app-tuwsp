package server

import (
	"app/src/models"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Broker struct {
	config *Config
	router *mux.Router
	servr  *http.Server
}

type binder func(Server, *mux.Router)

type entry func(*models.Login, *models.Signup) error

// Config method of Broker that return a Config
func (b *Broker) Config() *Config {
	return b.config
}

// Up method of Broker to up the server
func (b *Broker) Up(f binder, e entry) (err error) {
	b.router = mux.NewRouter()
	f(b, b.router)
	handler := cors.
		Default().Handler(b.router)
	b.servr = &http.Server{
		Handler: handler,
		Addr:    b.config.port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if err = e(b.config.Login, b.config.Signup); err != nil {
		return
	}
	log.Println("Starting to listen in port", b.config.Port())
	if err = b.servr.ListenAndServe(); err != nil {
		return
	}
	return
}
