package models

// NewUserData - Model for the uses table
type NewUserData struct {
	Username  string `form:"username" json:"username"`
	Email     string `form:"email" json:"email"`
	Phone     string `form:"phone" json:"phone"`
	Password  string `form:"password" json:"password"`
	Firstname string `form:"firstname" json:"firstname"`
	Lastname  string `form:"lastname" json:"lastname"`
	Usercode  string `form:"usercode" json:"usercode"`
	Lastlogin string `form:"lastlogin" json:"lastlogin"`
	Isactive  string `form:"isactive" json:"isactive"`
}
