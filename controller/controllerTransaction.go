package controller

import (
	//"Tubes/Art-Auction-Tubes/model"
	//"encoding/json"
	//"log"
	"net/http"
)

func GetLatestTransaction(w http.ResponseWriter, r *http.Request) {
	//db := connect()
	//defer db.Close()
	//query := "SELECT * FROM marketlist WHERE stateStatus = soldOut ;"

	//rows, err := db.Query(query)
	//var response model.MarketResponse

	//if err != nil {
	//	response.Status = 500
	//	response.Message = "Internal Server Error;" + err.Error()
	//	w.Header().Set("Content-Type", "application/json")
	//	json.NewEncoder(w).Encode(response)
	//	return
	//}

	//var market Market
	//var markets []Market

	//for rows.Next() {
	//	if err := rows.Scan(&market.id, &market.startingDate, &market.deadline, &market.startingBid, &market.buyOutBid, &market.datePosted, &market.stateStatus, &market.imageId); err != nil {
	//		log.Println(err)
	//		return
	//	} else {
	//		markets = append(markets, market)
	//	}
	//}

}
