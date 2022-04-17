package controller

import (
	"encoding/json"
	"net/http"

	"github.com/tubes/Art-Auction-Tubes/model"
)

func GetIncome(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	query := "SELECT income FROM accounting"
	var income float64
	err := db.QueryRow(query).Scan(&income)

	var response model.IncomeResponse
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		response.Status = 500
		response.Message = "Internal Server Error;" + err.Error()
	} else {
		response.Status = 200
		response.Message = "Success"
		response.Income = income
	}

	json.NewEncoder(w).Encode(response)
}
