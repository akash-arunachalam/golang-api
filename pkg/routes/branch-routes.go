package routes

import (
	"golang-api/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBranchRoutes = func(router *mux.Router) {

	router.HandleFunc("/Createbranch/", controllers.IsAuthorized(controllers.CreateBranch)).Methods("POST")
	router.HandleFunc("/Getallbranch/", controllers.IsAuthorized(controllers.GetAllBranch)).Methods("GET")
	router.HandleFunc("/Updatebranch/{branchId}", controllers.IsAuthorized(controllers.UpdateBranch)).Methods("PUT")
	router.HandleFunc("/Deletebranch/{branchId}", controllers.IsAuthorized(controllers.DeleteBranch)).Methods("DELETE")

}
