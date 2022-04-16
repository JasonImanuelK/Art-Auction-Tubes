package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/tubes/Art-Auction-Tubes/controller"
)

func main() {
	router := mux.NewRouter()

	// accessType 0 : User Biasa, 1 : Admin aja, 2 : 2 2nya bisa tapi butuh cookie.
	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Authenticate(controller.Logout, 2)).Methods("GET")
	router.HandleFunc("/picture/{id}", controller.Authenticate(controller.ReportPicture, 0)).Methods("PUT")
	router.HandleFunc("/users", controller.Authenticate(controller.GetUsers, 2)).Methods("GET")
	router.HandleFunc("/users/{key}", controller.Authenticate(controller.GetUsers, 2)).Methods("GET")
	router.HandleFunc("/users/{id}", controller.Authenticate(controller.ChangeBanStatus, 1)).Methods("PUT")
	router.HandleFunc("/bid", controller.InsertBid).Methods("POST")
	router.HandleFunc("/transaction", controller.GetLatestTransaction).Methods("GET")
	router.HandleFunc("/tax", controller.InsertTax).Methods("POST")
	router.HandleFunc("/tax", controller.GetTax).Methods("GET")
	//router.HandleFunc("/income",coontroller.GetIncome).Methods("GET")
	router.HandleFunc("/marketlist", controller.Authenticate(controller.GetMarketListByDate, 2)).Methods("POST")
	router.HandleFunc("/marketlist/id/{id}", controller.Authenticate(controller.GetMarketListById, 2)).Methods("GET")
	router.HandleFunc("/marketlist/{name}", controller.Authenticate(controller.GetMarketListByName, 2)).Methods("GET")
	router.HandleFunc("/buyout", controller.Authenticate(controller.Buyout, 2)).Methods("POST")
	router.HandleFunc("/marketlist/date", controller.Authenticate(controller.GetMarketListByDate, 2)).Methods("POST")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(router)

	http.Handle("/", router)

	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
	log.Fatal(http.ListenAndServe(":8080", router))

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("23:59").Do(func() { controller.CekDeadline() })
	s.StartBlocking()

}
