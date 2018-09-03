package main

import (
	"fmt"
	"github.com/edward-kwok/weather-api/internal"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hkweather/{date}", getHkWeather).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getHkWeather(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", internal.Extract(params["date"]))
}
