package controller

import (
	"encoding/json"
	"log"
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
		if err := rows.Scan(&market.ID, &market.StartingDate, &market.Deadline, &market.StartingBid, &market.BuyoutBid, &market.DatePosted, &market.Status, &market.ImageId); err != nil {
			log.Println(err)
			return
		} else {
			markets = append(markets, market)
		}
	}

}
