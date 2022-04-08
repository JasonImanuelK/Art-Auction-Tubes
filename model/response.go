package model

type GeneralResponse struct {
	status  int    `json:"Status"`
	message string `json:"Message"`
}
