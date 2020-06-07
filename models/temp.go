package models

// Temp - Model for the uses table
type Temp struct {
	Data  string      `form:"data" json:"data"`
	Roles string      `form:"roles" json:"roles"`
	Temp  interface{} `json:"temp"`
}
