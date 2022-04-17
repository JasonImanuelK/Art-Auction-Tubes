package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tubes/Art-Auction-Tubes/model"
)

func GetLatestTransaction(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	query := "SELECT * FROM marketlist WHERE stateStatus = 1 ORDER BY datePosted DESC ;"
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.Query(query)
	var response model.MarketResponse

	if err != nil {
		response.Status = 500
		response.Message = "Internal Server Error;" + err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	var market model.Market
	var markets []model.Market

	for rows.Next() {

		if err := rows.Scan(&market.ID, &market.StartingDate, &market.Deadline, &market.StartingBid, &market.BuyoutBid, &market.DatePosted, &market.Status, &market.ImageId); err != nil {
			fmt.Println(err.Error())
		} else {
			markets = append(markets, market)
		}
	}
	if err == nil {
		response.Status = 200
		response.Message = "Success"
		response.Data = markets
	} else {
		response.Status = 404
		response.Message = "Error - Data Not Found with Authentication Provided"
	}
	json.NewEncoder(w).Encode(response)
}
