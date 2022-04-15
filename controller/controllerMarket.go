package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/tubes/Art-Auction-Tubes/model"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

func ResetRedis(w http.ResponseWriter, r *http.Request) {
	var ctx = context.Background()
	db := connect()
	defer db.Close()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	er1r := r.ParseForm()
	if er1r != nil {
		return
	}

	email := r.Form.Get("email")

	err := client.Set(ctx, email, 0, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	var response model.GeneralResponse

	if err == nil {
		response.Status = 200
		response.Message = "success"

	} else {
		response.Status = 400
		response.Message = "internal error"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetMarketListByDate(w http.ResponseWriter, r *http.Request) {
	var ctx = context.Background()
	db := connect()
	defer db.Close()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	er1r := r.ParseForm()
	if er1r != nil {
		return
	}

	email := r.Form.Get("email")

	log.Print(email)

	current, err1 := client.Get(ctx, email).Result()

	current2, _ := strconv.Atoi(current)

	if err1 != nil {
		err := client.Set(ctx, email, 1, 0).Err()
		if err != nil {
			fmt.Println(err)
		}

	}

	rows, _ := db.Query("SELECT *  FROM marketlist WHERE ID >= ? LIMIT 10", current2)

	log.Print(current)

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

	err := client.Set(ctx, email, (current2 + 10), 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MarketResponse)
}

func GetMarketListById(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	vars := mux.Vars(r)
	marketId, _ := strconv.Atoi(vars["marketId"])

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
	name := vars["name"]
	rows, _ := db.Query(("SELECT marketlist.ID, marketlist.StartingDate, marketlist.Deadline, marketlist.StartingBid, marketlist.BuyoutBid, marketlist.DatePosted, marketlist.ImageId, marketlist.Status FROM marketlist JOIN gambar WHERE gambar.title LIKE '%" + name + "%'"))
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
	StartingBid, _ := strconv.Atoi(r.Form.Get("StartingBid"))
	BuyoutBid, _ := strconv.Atoi(r.Form.Get("BuyoutBid"))
	DatePosted := r.Form.Get("datePosted")
	ImageId, _ := strconv.Atoi(r.Form.Get("ImageId"))

	_, errQuery := db.Exec("INSERT INTO marketlist (startingDate,deadline,startingBid,buyoutBid,datePosted,imageId,status) values (?,?,?,?,?,?,'active')", StartingDate, Deadline, StartingBid, BuyoutBid, DatePosted, ImageId)

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

func Buyout(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	marketId, _ := strconv.Atoi(r.Form.Get("marketId"))
	userId, _ := strconv.Atoi(r.Form.Get("userId"))

	var response model.GeneralResponse

	rows, _ := db.Query("SELECT buyoutBid  FROM marketlist WHERE ID = ?", marketId)
	var eth int

	for rows.Next() {
		if err := rows.Scan(&eth); err != nil {
			log.Println(err)
			response.Status = 500
			response.Message = "internal error"
		}
	}

	_, errQuery := db.Exec("UPDATE marketlist SET status = 'ended' WHERE id = ?", marketId)

	_, errQuery1 := db.Exec("INSERT INTO bid (datePosted,etherium,userId,marketId) values (?,?,?,?)", time.Now().Format("01-02-2006"), eth, userId, marketId)

	// kurangi uang user di sini

	if errQuery == nil && errQuery1 == nil {
		response.Status = 200
		response.Message = "success"

	} else {
		response.Status = 400
		response.Message = "internal error"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
