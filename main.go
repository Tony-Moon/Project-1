package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Tony-Moon/Project-1/admin"
	_ "github.com/Tony-Moon/Project-1/admin"
	"github.com/Tony-Moon/Project-1/chat"
	_ "github.com/Tony-Moon/Project-1/chat"
	"github.com/Tony-Moon/Project-1/home"
	_ "github.com/Tony-Moon/Project-1/home"
	"github.com/Tony-Moon/Project-1/t3"
	_ "github.com/Tony-Moon/Project-1/t3"
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

	http.HandleFunc("/home", home.Home)
	http.HandleFunc("/ulog", home.UserLogin)
	http.HandleFunc("/chat", chat.Chat)
	http.HandleFunc("/ttt", t3.T3)
	http.HandleFunc("/alog", admin.AdminLogin)
	http.HandleFunc("/admin", admin.Admin)

	http.ListenAndServe(":9000", nil)
}
