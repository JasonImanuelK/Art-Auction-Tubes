package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tubes/Art-Auction-Tubes/model"
)

func GetIncome(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	query := "SELECT SUM(buyOutBid) FROM marketlist WHERE stateStatus = ? "
	err := r.ParseForm()
	if err != nil {
		return
	}
	state := r.Form.Get("stateStatus")
	rows, err := db.Query(query, state)

	var response model.MarketResponse
	if err != nil {
		response.Status = 500
		response.Message = "Internal Server Error;" + err.Error()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	var incomes []model.Market
	for rows.Next() {
		var state model.Market
		if err := rows.Scan(&state.Status); err != nil {
			fmt.Println(err.Error())
		} else {
			incomes = append(incomes, state)
		}
	}
	if len(incomes) > 0 {
		response.Status = 200
		response.Message = "Success"
		response.Data = incomes
	} else {
		response.Status = 404
		response.Message = "Error - Data Not Found with Authentication Provided"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
