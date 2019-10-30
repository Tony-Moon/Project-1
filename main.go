package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Tony-Moon/Project-1/web"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {

	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	data, err := sql.Open("postgres", datasource)
	defer data.Close()

	if err != nil {
		log.Println(err)
		panic(err)
	}

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/home", web.Home)
	http.HandleFunc("/ulog", web.UserLogin(data))
	http.HandleFunc("/chat", web.Chat)
	http.HandleFunc("/ttt", web.T3)
	http.HandleFunc("/alog", web.AdminLogin)
	http.HandleFunc("/admin", web.Admin)

	http.ListenAndServe(":9000", nil)
}
