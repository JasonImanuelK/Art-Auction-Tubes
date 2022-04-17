package model

type Market struct {
	ID           int     `json:"ID"`
	StartingDate string  `json:"startingDate"`
	Deadline     string  `json:"deadline"`
	StartingBid  float64 `json:"startingBid"`
	BuyoutBid    float64 `json:"buyoutBid"`
	DatePosted   string  `json:"datePosted"`
	ImageId      int     `json:"imageId"`
	Status       bool    `json:"status"`
}

type MarketResponse struct {
	Status  int      `json:"Status"`
	Message string   `json:"Message"`
	Data    []Market `json:"market"`
}
