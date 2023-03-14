package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Status string `json:"status"`
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Status: "OK"})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", getHealth).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
