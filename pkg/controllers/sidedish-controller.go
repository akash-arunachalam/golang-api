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

func UpdateSidedish(w http.ResponseWriter, r *http.Request) {
	var updateSidedish = &models.SideDish{}
	utils.ParseBody(r, updateSidedish)
	vars := mux.Vars(r)
	sidedishId := vars["sidedishId"]
	ID, err := strconv.ParseInt(sidedishId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	sidedishDetails, db := models.GetSidedishById(ID)
	if updateSidedish.Item != "" {
		sidedishDetails.Item = updateSidedish.Item
	}
	if updateSidedish.Quantity != "" {
		sidedishDetails.Quantity = updateSidedish.Quantity
	}
	if updateSidedish.Unit != "" {
		sidedishDetails.Unit = updateSidedish.Unit
	}

	db.Save(&sidedishDetails)
	res, _ := json.Marshal(sidedishDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteSidedish(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	branchId := vars["sidedishId"]
	ID, err := strconv.ParseInt(branchId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	sidedish := models.DeleteSidedish(ID)
	res, _ := json.Marshal(sidedish)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
