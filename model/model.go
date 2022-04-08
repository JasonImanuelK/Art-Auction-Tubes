package model

type User struct {
	ID            int    `json:"ID"`
	Username      string `json:"Username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	BlockedStatus bool   `json:"Status"`
}
