package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/PetoriousBIG/my-go-app/handlers"
	"github.com/PetoriousBIG/my-go-app/util"
	"github.com/gorilla/mux"
)

const FILE_PATH = "../countries_codes_and_coordinates.csv"

var PORT = os.Getenv("env_port")

func main() {

	l := log.New(os.Stdout, "my-api ", log.LstdFlags)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// create and pass handlers to serve mux
	setHandlers(sm, l)

	portArg := ":" + PORT

	// create a new server
	s := &http.Server{
		Addr:         portArg,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// start the server
	go func() {
		l.Println("Starting server on port", PORT)

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

	pingRouter := sm.Methods(http.MethodGet).Subrouter()
	apiRouter := sm.Methods(http.MethodGet).Subrouter()

	// ping
	ping := handlers.NewPing(l)
	pingRouter.HandleFunc("/ping", ping.Get)

	// get country data
	countryData := handlers.NewCountryData(l)
	apiRouter.HandleFunc("/At-A-Glance/{id:[A-Z]{3}}", countryData.GetCountryData)
	dict, err := util.ReadCountryCSV(FILE_PATH)
	if err != nil {
		l.Printf("Error reading csv: %s\n", err)
		os.Exit(1)
	}
	apiRouter.Use(countryData.GetMiddlewareValidateCountryFunc(dict))

	// not found
	sm.NotFoundHandler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		l.Println("[DEBUG] resource not found", r.URL.Path)

		rw.WriteHeader(http.StatusNotFound)
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"error": "resource does not exist",
		})
	})
}
