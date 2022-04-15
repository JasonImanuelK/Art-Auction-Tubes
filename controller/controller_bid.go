package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/tubes/Art-Auction-Tubes/model"
)

func InsertBid(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	var response model.BidResponse
	if err != nil {
		return
	}
	ether, _ := strconv.Atoi(r.Form.Get("etherium"))
	userid, _ := strconv.Atoi(r.Form.Get("userId"))
	marketid, _ := strconv.Atoi(r.Form.Get("marketId"))

	_, errQuery := db.Exec("INSERT INTO bid(datePosted,etherium,userId,marketId) VALUES (current_timestamp(),?,?,?);", ether, userid, marketid)
	if errQuery != nil {
		log.Print(errQuery.Error())
		response.Status = 400
		response.Message = "Bad Request - Insert Failed"
		json.NewEncoder(w).Encode(response)
	} else {
		response.Status = 200
		response.Message = "Success Insert"
		json.NewEncoder(w).Encode(response)
	}

}
