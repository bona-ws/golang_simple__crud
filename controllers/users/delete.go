package users

import (
	"encoding/json"
	"kriya_BE/helper"
	"log"
	"net/http"
)

// DeleteUsersMultipart ...
func DeleteUsersMultipart(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	id := r.FormValue("user_id")
	uid := r.FormValue("auth_uid")

	isAdmin := helper.IsAdmin(uid)

	// Delete user data if role is Admin
	if isAdmin {
		_, err = db.Exec("DELETE FROM users WHERE id = $1",
			id,
		)
		if err != nil {
			log.Print(err)
		}
		response.Status = 200
		response.Message = "Delete Success"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		log.Print(response)
	} else {
		response.Message = "You are not Admin"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response.Message)
	}
}
