package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func ping(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("pong")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", ping).Methods("GET")
	router.HandleFunc("/", ping).Methods("GET")
    connect()

	log.Fatal(http.ListenAndServe(":8000", router))
}
