package controllers

import (
	"encoding/json"
	"golang-api/pkg/models"
	"golang-api/pkg/utils"
	"net/http"
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
