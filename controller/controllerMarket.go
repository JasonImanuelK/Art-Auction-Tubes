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

func GetMarketListByName(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	vars := mux.Vars(r)
	marketId := vars["marketId"]

	rows, _ := db.Query("SELECT marketlist.ID, marketlist.StartingDate, marketlist.Deadline, marketlist.StartingBid, marketlist.BuyoutBid, marketlist.DatePosted, marketlist.ImageId, marketlist.Status FROM marketlist JOIN gambar WHERE gambar.title LIKE '%?%'", marketId)
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

func InsertMarket(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	StartingDate := r.Form.Get("startingDate")
	Deadline := r.Form.Get("deadline")
	StartingBid := r.Form.Get("startingBid")
	BuyoutBid := r.Form.Get("buyoutBid")
	DatePosted := r.Form.Get("datePosted")
	ImageId := r.Form.Get("imageId")
	Status := r.Form.Get("status")

	_, errQuery := db.Exec("INSERT INTO marketlist (startingDate,deadline,startingBid,buyoutBid,datePosted,imageId,status) values (?,?,?,?,?,?,?)", StartingDate, Deadline, StartingBid, BuyoutBid, DatePosted, ImageId, Status)

	var response model.GeneralResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "success"

	} else {
		response.Status = 400
		response.Message = "insert failed"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
