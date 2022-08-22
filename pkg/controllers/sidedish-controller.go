package controllers

import (
	"encoding/json"
	"golang-api/pkg/models"
	"golang-api/pkg/utils"
	"net/http"
)

func CreateSidedish(w http.ResponseWriter, r *http.Request) {
	CreateSidedish := &models.SideDish{}
	utils.ParseBody(r, CreateSidedish)
	b := CreateSidedish.CreateSidedish()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetSidedishes(w http.ResponseWriter, r *http.Request) {
	Sidedishes := models.GetSidedishes()
	res, _ := json.Marshal(Sidedishes)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)

}
