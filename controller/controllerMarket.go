package controller

import (
	"Tubes/Art-Auction-Tubes/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMarketListById(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	vars := mux.Vars(r)
	marketId := vars["marketId"]

	rows, _ := db.Query("SELECT *  FROM marketlist WHERE ID = ?", marketId)
	var MarketResponse model.MarketResponse
	var data model.Market

	for rows.Next() {
		if err := rows.Scan(&data.ID, &data.StartingDate, &data.Deadline, &data.StartingBid, &data.BuyoutBid, &data.DatePosted, &data.ImageId, &data.Status); err != nil {
			log.Println(err)
			MarketResponse.Status = 500
			MarketResponse.Message = "internal error"
			json.NewEncoder(w).Encode(MarketResponse)
			return
		} else {
			MarketResponse.Data = append(MarketResponse.Data, data)
			MarketResponse.Status = 200
			MarketResponse.Message = "Success"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MarketResponse)
}
