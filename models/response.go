package models

// Response - Model for the uses table
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}
