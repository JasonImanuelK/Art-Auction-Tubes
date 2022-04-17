package controller

import (
	"log"
	"time"

	"github.com/tubes/Art-Auction-Tubes/model"
)

func CekDeadline() {
	db := connect()
	defer db.Close()

	currentTime := time.Now()
	query := "SELECT  b.etherium, i.title, u.username,u.email, b.datePosted, b.marketId, b.id, i.imageId FROM (select id, datePosted,etherium, userId, marketId, max(etherium) over (partition by marketId) max_etherium from  bid ) b JOIN marketlist ml ON b.marketId = ml.id JOIN user u ON b.userId = u.id JOIN image i ON ml.imageId = i.id WHERE ml.deadline = '" + currentTime.Format("2006-02-01") + "' AND b.etherium = b.max_etherium;"
	log.Print(currentTime.Format("2006-02-01"))

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	var listEmails []model.ListEmail
	var listEmail model.ListEmail

	for rows.Next() {
		if err := rows.Scan(&listEmail.Etherium, &listEmail.Title, &listEmail.Username, &listEmail.Email, &listEmail.Date, &listEmail.MarketId, &listEmail.UserIdBuyer, &listEmail.ImageId); err != nil {
			log.Println(err)
			return
		} else {
			listEmails = append(listEmails, listEmail)
		}
	}

	SendMail(listEmails)
}
