package routes

import (
	users "kriya_BE/controllers/users"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// IndexRouter ...
func IndexRouter() {
	router := mux.NewRouter()
	http.Handle("/", router)

	router.HandleFunc("/getuser", users.GetUser).Methods("GET")
	router.HandleFunc("/getusers", users.GetAllUsers).Methods("GET")
	router.HandleFunc("/adduser", users.InsertUsersMultipart).Methods("POST")
	router.HandleFunc("/updateuser", users.UpdateUsersMultipart).Methods("PUT")
	router.HandleFunc("/deleteuser", users.DeleteUsersMultipart).Methods("Delete")

	fmt.Println("Connected to port 9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
