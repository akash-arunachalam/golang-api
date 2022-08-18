package routes

import (
	"simple-REST-master/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/signin/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/allusers/", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")

}
