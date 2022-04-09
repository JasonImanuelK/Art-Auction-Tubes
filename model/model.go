package model

import "time"

type User struct {
	ID            int    `json:"ID"`
	Username      string `json:"Username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	BlockedStatus bool   `json:"Status"`
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
