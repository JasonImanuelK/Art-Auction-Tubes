package controller

import (
	"net/http"
)

func GetLatestTransaction(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM "

}
