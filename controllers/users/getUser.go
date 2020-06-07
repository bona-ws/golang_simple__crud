package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// GetUser by ID...
func GetUser(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Parameter 'id' is missing")
		return
	}
	ID := keys[0]

	rows, err := db.Query("SELECT users.ID, users.Data, roles.data FROM users JOIN roles ON roles.id=users.role_id WHERE users.id = $1", ID)

	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		err := rows.Scan(&user.ID, &temp.Data, &temp.Roles)
		if err != nil {
			log.Fatal(err.Error())

		} else {
			f, s := temp.Temp, temp.Data
			if err := json.Unmarshal([]byte(s), &f); err != nil {
				panic(err)
			}
			ud, _ := f.(map[string]interface{})
			user.Email = fmt.Sprint(ud["email"])
			user.Username = fmt.Sprint(ud["username"])

			x, z := temp.Temp, temp.Roles
			if err := json.Unmarshal([]byte(z), &x); err != nil {
				panic(err)
			}
			rd, _ := x.(map[string]interface{})
			user.Role = fmt.Sprint(rd["role_name"])
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrUser
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
