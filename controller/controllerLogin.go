package controller

import (
	"Tubes/Art-Auction-Tubes/model"
	"encoding/json"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")

	errForm := r.ParseForm()
	if errForm != nil {
		var response model.GeneralResponse
		response.Status = 400
		response.Message = "Bad Request"
		json.NewEncoder(w).Encode(response)
		return
	}

	var user model.User
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	err := db.QueryRow("SELECT id,blockedStatus FROM user where username = ? and password = ?", username, password).Scan(&user.ID, &user.BlockedStatus)
	user.Username = username

	switch {
	case err != nil:
		log.Print(err)
		var response model.GeneralResponse
		response.Status = 400
		response.Message = "User not found."
		json.NewEncoder(w).Encode(response)
	default:
		generateToken(w, user.ID, user.Username, 0)
		var response model.GeneralResponse
		response.Status = 200
		response.Message = "Success"
		json.NewEncoder(w).Encode(response)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)
	var response model.GeneralResponse
	response.Status = 200
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
