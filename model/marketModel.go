package model

type Market struct {
	ID           int    `json:"ID"`
	StartingDate string `json:"startingDate"`
	Deadline     string `json:"deadline"`
	StartingBid  int    `json:"startingBid"`
	BuyoutBid    int    `json:"buyoutBid"`
	DatePosted   string `json:"datePosted"`
	ImageId      int    `json:"imageId"`
	Status       int    `json:"status"`
}

type MarketResponse struct {
	Status  int      `json:"Status"`
	Message string   `json:"Message"`
	Data    []Market `json:"market"`
}
