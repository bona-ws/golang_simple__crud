package users

import (
	"encoding/json"
	"fmt"
	"kriya_BE/helper"
	"log"
	"net/http"
	"strings"
)

// UpdateUsersMultipart ...
func UpdateUsersMultipart(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	id := r.FormValue("user_id")
	uid := r.FormValue("auth_uid")

	if err := r.ParseMultipartForm(4096); err != nil {
		log.Print(err)
	}

	isAdmin := helper.IsAdmin(uid)

	if isAdmin {
		for key, value := range r.PostForm {
			valueToString := strings.Join(value, "[ ]")
			if key != "user_id" {
				query := fmt.Sprintf("UPDATE users SET data = data::jsonb - '%s' || '{\"%s\": \"%s\"}' WHERE id = $1",
					key, key, valueToString,
				)

				_, err := db.Exec(query, id)
				if err != nil {
					log.Print(err)
				}
			}
		}
		response.Status = 200
		response.Message = "Success Update Data"
		log.Print("Update data to database")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		response.Message = "You are not Admin"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response.Message)
	}

}
