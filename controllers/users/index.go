package users

import (
	"kriya_BE/connection"
	"kriya_BE/models"
)

var db = connection.Connect()
var user models.User
var newUserData models.NewUserData
var users models.Users
var temp models.Temp
var response models.Response
var arrUser []models.Users
