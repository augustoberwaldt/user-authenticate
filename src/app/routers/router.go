package routers


import (
	"../controllers"
	"github.com/gorilla/mux"
)

func SetRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/authenticate", controllers.Authenticate).Methods("POST")
	return router
}