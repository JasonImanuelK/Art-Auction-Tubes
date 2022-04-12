package controller

import (
	"Tubes/Art-Auction-Tubes/model"
	"encoding/json"
	"log"
	"net/http"

	bcrypt "golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

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
	hashedPassword, _ := hashPassword(password)
	err := db.QueryRow("SELECT id,blockedStatus,userType FROM user where username = ? and password = ?", username, hashedPassword).Scan(&user.ID, &user.BlockedStatus, &user.UserType)
	user.Username = username

	switch {
	case err != nil:
		log.Print(err)
		var response model.GeneralResponse
		response.Status = 400
		response.Message = "User not found."
		json.NewEncoder(w).Encode(response)
	default:
		generateToken(w, user.ID, user.Username, user.UserType)
		var response model.GeneralResponse
		response.Status = 200
		response.Message = "Login Success"
		json.NewEncoder(w).Encode(response)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
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

	username := r.Form.Get("username")
	password := r.Form.Get("password")
	email := r.Form.Get("email")
	hashedPassword, _ := hashPassword(password)
	var id int

	_, errQuery := db.Exec("INSERT INTO user(username, password, email,blockedStatus,userType) values (?,?,?,0,0)", username, hashedPassword, email)

	w.Header().Set("Content-Type", "application/json")
	var response model.GeneralResponse
	json.NewEncoder(w).Encode(response)

	if errQuery != nil {
		response.Status = 400
		response.Message = "Bad Request"
		return
	}

	err := db.QueryRow("SELECT id FROM user where username = ? and password = ?", username, password).Scan(&id)

	if err == nil {
		response.Status = 200
		response.Message = "Register Success"
	} else {
		response.Status = 400
		response.Message = "Register Failed"
	}

	generateToken(w, id, username, 0)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)
	var response model.GeneralResponse
	response.Status = 200
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
