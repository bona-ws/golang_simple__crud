package models

// Users - Model for the uses table
type Users struct {
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
	IsActive string `form:"isactive" json:"isactive"`
}
