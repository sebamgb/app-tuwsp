/*
 * Author: Sebastián Romero
 * Project: tuwsp
 */
package main

import (
	"app/src/routes"
	"app/src/server"

	"context"
	"log"

	"github.com/joho/godotenv"
)

// init preinitialization
func init() {
	if err := godotenv.
		Load(".env"); err != nil {
		log.Fatal(err)
		return
	}
}

// main calls to toraly web server
func main() {
	if config, err := server.
		NewConfig(); err != nil {
		log.Fatal(err)
	} else if s, err := server.
		NewServer(
			context.Background(),
			config,
		); err != nil {
		log.Fatal(err)
	} else if s.
		Up(
			routes.BindRoute,
		); err != nil {
		log.Fatal(err)
	}
}
