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

	router.HandleFunc("/login", controller.Login).Methods("PUT")
	router.HandleFunc("/logout", controller.Logout).Methods("GET")
	router.HandleFunc("/picture/{id}", controller.ReportPicture).Methods("PUT")
	router.HandleFunc("/users", controller.GetUsers).Methods("GET")
	router.HandleFunc("/users/{key}", controller.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controller.ChangeBanStatus).Methods("PUT")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
