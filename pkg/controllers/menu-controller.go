package controllers

import (
	"encoding/json"
	"golang-api/pkg/models"
	"golang-api/pkg/utils"
	"net/http"
)

func CreateMenu(w http.ResponseWriter, r *http.Request) {
	CreateMenu := &models.Menu{}
	utils.ParseBody(r, CreateMenu)
	b := CreateMenu.CreateMenu()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMenu(w http.ResponseWriter, r *http.Request) {
	Menu := models.GetMenu()
	res, _ := json.Marshal(Menu)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)

}
