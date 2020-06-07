package users

import (
	"encoding/json"
	"fmt"
	"kriya_BE/helper"
	"log"
	"net/http"
)

// InsertUsersMultipart ...
func InsertUsersMultipart(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	uid := r.FormValue("auth_uid")
	id := r.FormValue("user_id")
	roleid := r.FormValue("role_id")
	newUserData.Email = r.FormValue("email")
	newUserData.Phone = r.FormValue("phone")
	newUserData.Password = r.FormValue("password")
	newUserData.Username = r.FormValue("username")
	newUserData.Firstname = r.FormValue("firstname")
	newUserData.Lastname = r.FormValue("lastname")
	newUserData.Usercode = r.FormValue("user_code")
	newUserData.Lastlogin = r.FormValue("last_login")
	newUserData.Isactive = r.FormValue("is_active")

	isAdmin := helper.IsAdmin(uid)

	convertToJSON, err := json.Marshal(newUserData)
	if err != nil {
		log.Print(err)
	}

	dataToJSON := fmt.Sprint(string(convertToJSON))

	if isAdmin {
		_, err = db.Exec("INSERT INTO users (id, data, role_id) values ($1,$2, $3)",
			id,
			dataToJSON,
			roleid,
		)
		if err != nil {
			log.Print(err)
		}
		response.Status = 200
		response.Message = "User added successfully"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		log.Print(response)
	} else {
		response.Message = "You are not Admin"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response.Message)
	}

}
