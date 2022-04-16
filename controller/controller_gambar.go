package controller

import (
	"encoding/json"
	"fmt"
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

	_, errQuery := db.Exec("UPDATE image SET report=report+1 WHERE ID = ?",
		id,
	)

	var response model.GeneralResponse
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
