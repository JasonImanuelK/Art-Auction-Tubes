package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tubes/Art-Auction-Tubes/model"
)

func GetLatestTransaction(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	query := "SELECT * FROM marketlist WHERE stateStatus = 1 ORDER BY datePosted DESC ;"

	rows, err := db.Query(query)
	var response model.MarketResponse

	if err != nil {
		response.Status = 500
		response.Message = "Internal Server Error;" + err.Error()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	var market model.Market
	var markets []model.Market

	for rows.Next() {

		var checkStartingDate sql.NullTime
		var checkDeadline sql.NullTime
		var checkDatePosted sql.NullTime
		if err := rows.Scan(&market.ID, checkStartingDate, checkDeadline, &market.StartingBid, &market.BuyoutBid, checkDatePosted, &market.Status, &market.ImageId); err != nil {
			fmt.Println(err.Error())
		} else {
			markets = append(markets, market)
		}
	}
	if len(markets) > 0 {
		response.Status = 200
		response.Message = "Success"
	} else {
		response.Status = 404
		response.Message = "Error - Data Not Found with Authentication Provided"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
