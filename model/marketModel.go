package model

import "time"

type Market struct {
	ID           int       `json:"ID"`
	StartingDate time.Time `json:"startingDate"`
	Deadline     time.Time `json:"deadline"`
	StartingBid  float64   `json:"startingBid"`
	BuyoutBid    float64   `json:"buyoutBid"`
	DatePosted   time.Time `json:"datePosted"`
	ImageId      int       `json:"imageId"`
	Status       bool      `json:"status"`
}

type MarketResponse struct {
	Status  int      `json:"Status"`
	Message string   `json:"Message"`
	Data    []Market `json:"market"`
}
