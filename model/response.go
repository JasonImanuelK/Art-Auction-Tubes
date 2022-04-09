package model

type GeneralResponse struct {
	Status  int    `json:"Status"`
	Message string `json:"Message"`
}

type Gambaresponse struct {
	Status  int     `form:"status" json:"status"`
	Message string  `form:"message" json:"message"`
	Data    []Image `form:"data" json:"data"`
}

type UserResponse struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}
