package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-REST-master/pkg/models"
	"simple-REST-master/pkg/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var Signin models.Signin

func CreateUser(w http.ResponseWriter, r *http.Request) {
	Signin := &models.Signin{}
	utils.ParseBody(r, Signin)
	b := Signin.CreateUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllUsers()
	res, _ := json.Marshal(newUsers)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
