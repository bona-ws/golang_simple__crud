package helper

import (
	"encoding/json"
	"fmt"
	"log"
)

// IsAdmin Checking user role admin
func IsAdmin(uid string) bool {
	var role string
	rows, err := db.Query("SELECT users.ID, roles.data FROM users JOIN roles ON roles.id=users.role_id WHERE users.id = $1", uid)
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		err := rows.Scan(&user.ID, &temp.Roles)
		if err != nil {
			log.Fatal(err.Error())

		} else {
			x, z := temp.Temp, temp.Roles
			if err := json.Unmarshal([]byte(z), &x); err != nil {
				panic(err)
			}
			rd, _ := x.(map[string]interface{})
			role = fmt.Sprint(rd["role_name"])
		}
	}
	if role == "Admin" {
		return true
	}
	return false
}
