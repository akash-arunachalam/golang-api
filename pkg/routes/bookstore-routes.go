package routes

import (
	"simple-REST-master/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/signin/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/allusers/", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/notification/", controllers.Notification).Methods("GET")

}
