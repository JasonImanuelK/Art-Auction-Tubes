package model

import "time"

type User struct {
	ID            int    `json:"ID"`
	Username      string `json:"Username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	BlockedStatus bool   `json:"Status"`
	UserType      int    `json:"userType"`
}

type Image struct {
	ID          int       `form:"ID" json:"ID"`
	Title       string    `form:"title" json:"title"`
	Directory   string    `form:"directory" json:"directory"`
	Report      int       `form:"report" json:"report"`
	UserID      int       `form:"userid" json:"userid"`
	DatePosted  time.Time `form:"dataposted" json:"dataposted"`
	Description string    `form:"description" json:"description"`
	Status      int       `form:"status" json:"status"`
}

type Bid struct {
	ID         int       `json:"id"`
	DatePosted time.Time `json:"datePosted"`
	Etherium   float64   `json:"etherium"`
	UserID     int       `json:"userId"`
	MarketId   int       `json:"marketId"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type BidResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Bid    `json:"data"`
}

type BidsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Bid  `json:"data"`
}

type TaxResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Tax     float64 `json:"tax"`
}

type IncomeResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Income  float64 `json:"income"`
}
