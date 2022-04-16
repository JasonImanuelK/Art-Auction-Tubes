package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tubes/Art-Auction-Tubes/model"

	"github.com/gorilla/mux"
)

func ReportPicture(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	param := mux.Vars(r)
	id := param["id"]
	fmt.Println(id)

	query := "SELECT report FROM image WHERE id = " + id

	var report int

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&report); err != nil {
			log.Fatal(err.Error())
		}
	}

	fmt.Println("Masuk line 40", report)

	report++

	_, errQuery := db.Exec("UPDATE image SET report=? WHERE ID = ?",
		report,
		id,
	)
	fmt.Println("Masuk line 48")

	var response model.Gambaresponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success Add Report Picture"
	} else {
		response.Status = 400
		response.Message = "Failed Add Report Picture"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
