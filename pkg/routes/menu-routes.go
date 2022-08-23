package routes

import (
	"golang-api/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterMenuRoutes = func(router *mux.Router) {

	router.HandleFunc("/CreateSidedish/", controllers.CreateSidedish).Methods("POST")
	router.HandleFunc("/GetSidedishes/", controllers.GetSidedishes).Methods("GET")
	router.HandleFunc("/UpdateSidedish/{sidedishId}", controllers.GetSidedishes).Methods("PUT")
	router.HandleFunc("/DeleteSidedish/{sidedishId}", controllers.GetSidedishes).Methods("DELETE")

	router.HandleFunc("/CreateFooditem/", controllers.CreateFooditem).Methods("POST")
	router.HandleFunc("/GetFooditems/", controllers.GetFooditems).Methods("GET")
}
