package model

import "time"

type ListEmail struct {
	Title    string
	Username string
	Email    string
	Date     time.Time
	Etherium float64
	MarketId int
}
