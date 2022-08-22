package controllers

import (
	"encoding/json"
	"golang-api/pkg/models"
	"golang-api/pkg/utils"
	"net/http"
)

func CreateFooditem(w http.ResponseWriter, r *http.Request) {
	CreateFooditem := &models.FoodItem{}
	utils.ParseBody(r, CreateFooditem)
	b := CreateFooditem.CreateFooditem()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFooditems(w http.ResponseWriter, r *http.Request) {
	Fooditems := models.GetFooditems()
	res, _ := json.Marshal(Fooditems)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)

}
