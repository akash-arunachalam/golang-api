package routes

import (
	"golang-api/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	
	router.HandleFunc("/signup/", controllers.SignUp).Methods("POST")
	router.HandleFunc("/signin/", controllers.SignIn).Methods("POST")

	router.HandleFunc("/allusers/", controllers.IsAuthorized(controllers.GetAllUsers)).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.IsAuthorized(controllers.DeleteUser)).Methods("DELETE")
	router.HandleFunc("/user/{userId}", controllers.IsAuthorized(controllers.UpdateUser)).Methods("PUT")

}
