package main

import (
	"./container"
	"./routers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	container.Connect()
	router := mux.NewRouter()
	router = routers.SetRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
