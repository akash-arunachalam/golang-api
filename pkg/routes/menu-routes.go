package routes

import (
	"golang-api/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterMenuRoutes = func(router *mux.Router) {

	router.HandleFunc("/CreateSidedish/", controllers.IsAuthorized(controllers.CreateSidedish)).Methods("POST")
	router.HandleFunc("/GetSidedishes/", controllers.IsAuthorized(controllers.GetSidedishes)).Methods("GET")
	router.HandleFunc("/UpdateSidedish/{sidedishId}", controllers.IsAuthorized(controllers.UpdateSidedish)).Methods("PUT")
	router.HandleFunc("/DeleteSidedish/{sidedishId}", controllers.IsAuthorized(controllers.DeleteSidedish)).Methods("DELETE")

	router.HandleFunc("/CreateFooditem/", controllers.IsAuthorized(controllers.CreateFooditem)).Methods("POST")
	router.HandleFunc("/GetFooditems/", controllers.IsAuthorized(controllers.GetFooditems)).Methods("GET")

	router.HandleFunc("/CreateMenu/", controllers.IsAuthorized(controllers.CreateMenu)).Methods("POST")
	router.HandleFunc("/GetMenu/", controllers.IsAuthorized(controllers.GetMenu)).Methods("GET")
}
