package model

type UserDTO struct {
	ID            int    `json:"ID"`
	Username      string `json:"Username"`
	Email         string `json:"email"`
	BlockedStatus bool   `json:"Status"`
	UserType      int    `json:"userType"`
	CountReport   int    `json:"countReport"`
}

type UserDTOResponse struct {
	Status  int       `form:"status" json:"status"`
	Message string    `form:"message" json:"message"`
	Data    []UserDTO `form:"data" json:"data"`
}
