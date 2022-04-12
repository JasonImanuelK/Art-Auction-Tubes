package main

import (
	"Tubes/Art-Auction-Tubes/controller"

	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// accessType 0 : User Biasa, 1 : Admin aja, 2 : 2 2nya bisa tapi butuh cookie.
	router.HandleFunc("/register", controller.Register).Methods("PUT")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Authenticate(controller.Logout, 2)).Methods("GET")
	router.HandleFunc("/picture/{id}", controller.Authenticate(controller.ReportPicture, 0)).Methods("PUT")
	router.HandleFunc("/users", controller.Authenticate(controller.GetUsers, 2)).Methods("GET")
	router.HandleFunc("/users/{key}", controller.Authenticate(controller.GetUsers, 2)).Methods("GET")
	router.HandleFunc("/users/{id}", controller.Authenticate(controller.ChangeBanStatus, 1)).Methods("PUT")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
