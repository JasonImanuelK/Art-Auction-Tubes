package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/tubes/Art-Auction-Tubes/model"
)

func InsertTax(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	err := r.ParseForm()
	var response model.AccountingResponse
	if err != nil {
		return
	}

	tax, _ := strconv.Atoi(r.Form.Get("tax"))
	income, _ := strconv.Atoi(r.Form.Get("income"))
	_, errQuery := db.Exec("INSERT INTO accounting(tax,income) VALUES (?,?);", tax, income)
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

func GetTax(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	query := "SELECT tax FROM accounting WHERE accounting.tax= ? "
	err := r.ParseForm()
	if err != nil {
		return
	}
	tax := r.Form.Get("tax")
	rows, err := db.Query(query, tax)

	var response model.AccountingResponse
	if err != nil {
		response.Status = 500
		response.Message = "Internal Server Error;" + err.Error()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	var accounts []model.Accounting
	for rows.Next() {
		var tax model.Accounting
		if err := rows.Scan(&tax.Tax, &tax.Income); err != nil {
			fmt.Println(err.Error())
		} else {
			accounts = append(accounts, tax)
		}
	}
	if len(accounts) > 0 {
		response.Status = 200
		response.Message = "Success"
	} else {
		response.Status = 404
		response.Message = "Error - Data Not Found with Authentication Provided"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
