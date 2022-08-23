package controllers

import (
	"encoding/json"
	"fmt"
	"golang-api/pkg/models"
	"golang-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBranch models.Branch

func CreateBranch(w http.ResponseWriter, r *http.Request) {
	CreateBranch := &models.Branch{}
	utils.ParseBody(r, CreateBranch)
	b := CreateBranch.CreateBranch()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllBranch(w http.ResponseWriter, r *http.Request) {
	newBranches := models.GetAllBranch()
	res, _ := json.Marshal(newBranches)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)

}

func UpdateBranch(w http.ResponseWriter, r *http.Request) {
	
	var updateBranch = &models.Branch{}
	utils.ParseBody(r, updateBranch)
	vars := mux.Vars(r)
	branchId := vars["branchId"]
	ID, err := strconv.ParseInt(branchId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	branchDetails, db := models.GetBranchById(ID)
	if updateBranch.Branchname != "" {
		branchDetails.Branchname = updateBranch.Branchname
	}

	db.Save(&branchDetails)
	res, _ := json.Marshal(branchDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBranch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	branchId := vars["branchId"]
	ID, err := strconv.ParseInt(branchId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	branch := models.DeleteBranch(ID)
	res, _ := json.Marshal(branch)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
