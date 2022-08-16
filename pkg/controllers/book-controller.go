package controllers

import (
	"encoding/json"
	"fmt"
	model "golang-api/pkg/models"
	"net/http"
	"strconv"

	"golang-api/pkg/utils"

	"github.com/gorilla/mux"
)

var NewBook model.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {

	newBook := model.GetBooks()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("err while parsing")
	}
	bookDetails, _ := model.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	createBook := &model.Book{}
	utils.ParseBody(r, createBook)

	b := createBook.CreateBook()

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
