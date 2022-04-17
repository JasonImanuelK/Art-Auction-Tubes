package controller

import (
	"encoding/json"
	"net/http"

	"github.com/tubes/Art-Auction-Tubes/model"
)

func UpdateTax(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	err := r.ParseForm()
	var response model.GeneralResponse
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		response.Status = 400
		response.Message = "Bad Request"
		json.NewEncoder(w).Encode(response)
		return
	}
	tax := r.Form.Get("tax")
	_, errQuery := db.Exec("UPDATE accounting SET TAX = ?", tax)

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		json.NewEncoder(w).Encode(response)
	} else {
		response.Status = 400
		response.Message = "Bad Request"
		json.NewEncoder(w).Encode(response)
	}

}

func GetTax(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	query := "SELECT tax FROM accounting"
	var tax float64
	err := db.QueryRow(query).Scan(&tax)

	var response model.TaxResponse
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		response.Status = 500
		response.Message = "Internal Server Error;" + err.Error()
	} else {
		response.Status = 200
		response.Message = "Success"
		response.Tax = tax
	}

	json.NewEncoder(w).Encode(response)
}
