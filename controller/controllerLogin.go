package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tubes/Art-Auction-Tubes/model"

	bcrypt "golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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
	err := db.QueryRow("SELECT id,blockedStatus,userType FROM user where username = ?", username).Scan(&user.ID, &user.BlockedStatus, &user.UserType)
	user.Username = username

	var response model.GeneralResponse
	switch {
	case err != nil:
		log.Print(err)
		response.Status = 400
		response.Message = "User not found."
		json.NewEncoder(w).Encode(response)
	default:
		match := CheckPasswordHash(password, hashedPassword)
		if match {
			generateToken(w, user.ID, user.Username, user.UserType)
			response.Status = 200
			response.Message = "Login Success"
		} else {
			response.Status = 400
			response.Message = "Password Failed"
		}

		json.NewEncoder(w).Encode(response)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")

	errForm := r.ParseForm()

	var response model.GeneralResponse
	if errForm != nil {
		response.Status = 400
		response.Message = "Bad Request"
		json.NewEncoder(w).Encode(response)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")
	email := r.Form.Get("email")
	hashedPassword, _ := hashPassword(password)
	etherium_key := r.Form.Get("etherium_key")
	etherium_password := r.Form.Get("etherium_password")
	coin := r.Form.Get("coin")

	var id int
	var hash string
	_, errQuery := db.Exec("INSERT INTO user(username, password, email,blockedStatus,userType) values (?,?,?,0,0)", username, hashedPassword, email)
	if errQuery != nil {
		response.Status = 400
		response.Message = "Bad Request"
		json.NewEncoder(w).Encode(response)
		return
	}

	rows, err := db.Query("SELECT id, password FROM user where username = '" + username + "'")
	if err != nil {
		log.Print(err)
	}

	rows.Next()

	if err := rows.Scan(&id, &hash); err == nil {
		match := CheckPasswordHash(password, hash)
		if match {
			_, errQuery2 := db.Exec("INSERT INTO user_wallet values (?,?,?,?)", etherium_key, etherium_password, coin, id)
			if errQuery2 != nil {
				response.Status = 400
				response.Message = "Bad Request"
				_, errQuery3 := db.Exec("DELETE FROM user WHERE id = ?", id)
				if errQuery3 != nil {
					log.Println(errQuery)
				}

				json.NewEncoder(w).Encode(response)
				return
			}
			response.Status = 200
			response.Message = "Register Success"
			generateToken(w, id, username, 0)
		} else {
			response.Status = 400
			response.Message = "Register Failed"
		}
	} else {
		log.Fatal(err.Error())
		response.Status = 400
		response.Message = "Register Failed"
	}

	json.NewEncoder(w).Encode(response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)
	var response model.GeneralResponse
	response.Status = 200
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
