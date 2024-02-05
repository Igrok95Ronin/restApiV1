package main

import (
	"log"
	"mymodule/internal/handlers"
	"mymodule/internal/templates"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	var handler handlers.Ihandler
	handler = &templates.Handler{}
	handler.Register(mux)

	start(mux)
}

func start(mux *http.ServeMux) {

	server := &http.Server{
		Addr:         ":8081",
		Handler:      mux,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

}
