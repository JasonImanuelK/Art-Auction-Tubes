package controller

import (
	"log"
	"os"
	"strconv"

	"github.com/tubes/Art-Auction-Tubes/model"
	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587

var CONFIG_SENDER_NAME = os.Getenv("CONFIG_SENDER_NAME")
var CONFIG_AUTH_EMAIL = os.Getenv("CONFIG_AUTH_EMAIL")
var CONFIG_AUTH_PASSWORD = os.Getenv("CONFIG_AUTH_PASSWORD")

func SendMail(email []model.ListEmail) {
	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)
	s, err := dialer.Dial()
	if err != nil {
		panic(err)
	}

	var totalEtherium float64

	mailer := gomail.NewMessage()
	for _, r := range email {
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetAddressHeader("To", r.Email, r.Username)
		mailer.SetHeader("Subject", "You won the bid !!!")
		mailer.SetBody("text/html", "Hello "+r.Username+" !"+" you won the bid at "+r.Date.String()+" with "+strconv.FormatFloat(r.Etherium, 'f', 2, 64)+" etherium.")

		totalEtherium = totalEtherium + r.Etherium
		hapusBid(r.MarketId)
		gantiStatusMarketList(r.MarketId)

		if err := gomail.Send(s, mailer); err != nil {
			log.Printf("Could not send email to %q: %v", r.Email, err)
		}
		mailer.Reset()
	}
	tambahIncome(totalEtherium)
	log.Println("Mail sent!")
}

func tambahIncome(totalEth float64) {
	db := connect()
	defer db.Close()

	var tax float64
	err := db.QueryRow("SELECT tax FROM accounting").Scan(&tax)
	if err != nil {
		log.Print(err)
		return
	}

	totalIncome := totalEth * tax

	_, errQuery := db.Exec("UPDATE income SET income = income+?", totalIncome)

	if errQuery != nil {
		log.Println(errQuery)
	} else {
		log.Println("Income berhasil ditambahkan.")
	}
}

func hapusBid(marketId int) {
	db := connect()
	defer db.Close()

	_, errQuery := db.Exec("DELETE FROM bid WHERE marketId = ?", marketId)

	if errQuery != nil {
		log.Println(errQuery)
	} else {
		log.Println("Bid terhapus.")
	}
}

func gantiStatusMarketList(marketId int) {
	db := connect()
	defer db.Close()

	_, errQuery := db.Exec("UPDATE marketlist SET stateStatus = 0 WHERE id = ?", marketId)

	if errQuery != nil {
		log.Println(errQuery)
	} else {
		log.Println("MarketList dengan id " + strconv.Itoa(marketId) + " sudah ditutup.")
	}
}
