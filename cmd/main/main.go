package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang-api/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	routes.RegisterUserRoutes(r)
	http.Handle("/", r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	//log.Fatal(http.ListenAndServe("localhost:8080", r))
}
