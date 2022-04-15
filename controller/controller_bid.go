package controller

import (
	"encoding/json"
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

	if userid != 0 && marketid != 0 {
		_, errQuery := db.Exec("INSERT INTO `bid`(`id`,`datePosted`,`etherium`,`userId`,`marketId`) VALUES (NULL,current_timestamp(),?,?,?);", ether, userid, marketid)
		if errQuery != nil {
			response.Status = 500
			response.Message = "Internal Server Error"
		} else {
			response.Status = 200
			response.Message = "Success Insert"
		}
	} else {
		response.Status = 400
		response.Message = "Bad Request - Insert Failed"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
