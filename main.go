package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/PetoriousBIG/docker-ex/handlers"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "my-api ", log.LstdFlags)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// create and pass handlers to serve mux
	setHandlers(sm, l)

	// create a new server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received
	sig := <-c
	l.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}

func setHandlers(sm *mux.Router, l *log.Logger) {

	getRouter := sm.Methods(http.MethodGet).Subrouter()

	// ping
	ping := handlers.NewPing(l)
	getRouter.HandleFunc("/ping", ping.Get)

	// get country data
	countryData := handlers.NewCountryData(l)
	getRouter.HandleFunc("/At-A-Glance/{id:[A-Z]{3}}", countryData.GetCountryData)

}
