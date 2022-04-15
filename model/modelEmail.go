package model

import "time"

type ListEmail []struct {
	Title    string
	Name     string
	Email    string
	Date     time.Time
	Etherium float64
}
