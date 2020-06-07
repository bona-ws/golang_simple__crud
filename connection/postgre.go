package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // import postgre driver
	"log"
)

const (
	host     = "18.140.67.82"
	port     = 8976
	user     = "test"
	password = "kriyatest123"
	dbname   = "kriya_test"
)

// Connect db
func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	return db
}
