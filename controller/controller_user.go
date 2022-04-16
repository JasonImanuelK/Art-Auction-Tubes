package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/tubes/Art-Auction-Tubes/model"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	param := mux.Vars(r)
	key := param["key"]

	query := "SELECT id, username, email, password, blockedStatus FROM user WHERE userType = 0"

	if key != "" {
		query += " AND username LIKE '%" + key + "%'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	var user model.User
	var users []model.User
	var temp int

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &temp); err != nil {
			log.Fatal(err.Error())
		} else {
			if temp == 0 {
				user.BlockedStatus = false
			} else {
				user.BlockedStatus = true
			}
			users = append(users, user)
		}
	}

	var response model.UserResponse
	if len(users) > 0 {
		response.Status = 200
		response.Message = "Success Get Users"
		response.Data = users
	} else {
		response.Status = 400
		response.Message = "Failed Get Users"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ChangeBanStatus(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	param := mux.Vars(r)
	id := param["id"]

	query := "SELECT blockedStatus FROM user WHERE ID = " + id

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	var temp int

	for rows.Next() {
		if err := rows.Scan(&temp); err != nil {
			log.Fatal(err.Error())
		} else {
			if temp == 0 {
				temp = 1
			} else {
				temp = 0
			}
		}
	}

	query = "UPDATE user SET blockedStatus = " + strconv.Itoa(temp) + " WHERE ID = " + id

	_, errQuery := db.Exec(query)

	var response model.UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success Change Status User"
	} else {
		response.Status = 400
		response.Message = "Failed Change Status User"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
