package chat

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

/*
AddTo doc
*/
func AddTo(macaddress string, username string, userpass string, nickname string, userrank bool) {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()

	if err != nil {
		log.Println(err)
		panic(err)
	}
	db.Exec("INSERT INTO servicers VALUES ($1, $2, $3, $4, $5);",
		macaddress, username, userpass, nickname, userrank)
}
