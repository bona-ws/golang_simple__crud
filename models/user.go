package models

// User - Model for the uses table
type User struct {
	ID       string `form:"ID" json:"ID"`
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
	Role     string `form:"role" json:"role"`
}
