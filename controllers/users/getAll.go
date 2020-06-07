package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// GetAllUsers ...
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT data FROM users")

	if err != nil {
		log.Print(err)
		json.NewEncoder(w).Encode(err)
	}
	for rows.Next() {
		err := rows.Scan(&temp.Data)
		if err != nil {
			log.Fatal(err.Error())

		} else {
			f, s := temp.Temp, temp.Data
			if err := json.Unmarshal([]byte(s), &f); err != nil {
				panic(err)
			}
			ud, _ := f.(map[string]interface{})
			users.Email = fmt.Sprint(ud["email"])
			users.Username = fmt.Sprint(ud["username"])
			users.IsActive = fmt.Sprint(ud["is_active"])

			arrUser = append(arrUser, users)
		}
	}
	response.Data = arrUser
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response.Data)
}
