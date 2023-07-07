package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aryangp/goRest/controller"
	"github.com/Aryangp/goRest/database"

	"github.com/gorilla/mux"
)

func intializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.RegisterUser).Methods("POST")
	r.HandleFunc("/login", controller.LoginUser).Methods("POST")
	r.HandleFunc("/home", controller.VerifyJwt(controller.Home)).Methods("GET")
	// r.HandleFunc("/users", controller.CreateUsers).Methods("GET")
	// r.HandleFunc("/user/{id}", controller.CreateUsers).Methods("GET")
	// r.HandleFunc("/users", controller.CreateUsers).Methods("POST")
	// r.HandleFunc("/users/{id}", controller.CreateUsers).Methods("DELETE")
	// r.HandleFunc("/users/{id}", controller.CreateUsers).Methods("PUT")
	fmt.Println("serve is started")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	database.InitialMigration()
	intializeRouter()

}
