package routes

import (
	"golang-api/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBranchRoutes = func(router *mux.Router) {

	router.HandleFunc("/Createbranch/", controllers.CreateBranch).Methods("POST")
	router.HandleFunc("/Getallbranch/", controllers.GetAllBranch).Methods("GET")
	router.HandleFunc("/Updatebranch/{branchId}", controllers.UpdateBranch).Methods("PUT")
	router.HandleFunc("/Deletebranch/{branchId}", controllers.DeleteBranch).Methods("DELETE")

}
