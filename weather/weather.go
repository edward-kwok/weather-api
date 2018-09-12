package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/edward-kwok/weather-api/internal"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hkweather/{date}", getHkWeather).Methods("GET")
	port := os.Getenv("PORT") // Heroku provides the port to bind to
	if port == "" {
		port = "8080"
	}
	log.Println("Started HK Weather API")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getHkWeather(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", internal.Extract(params["date"]))
}
