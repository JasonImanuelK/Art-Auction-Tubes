package model

type user struct {
	ID            int    `json:"ID"`
	username      string `json:"Username"`
	email         string `json:"email"`
	password      string `json:"password"`
	blockedStatus bool   `json:"Status"`
}
