package helper

import (
	"kriya_BE/connection"
	"kriya_BE/models"
)

var db = connection.Connect()
var user models.User
var temp models.Temp
