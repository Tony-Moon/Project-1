package main

import (
	"net/http"

	"github.com/Tony-Moon/Project-1/web"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/home", web.Home)
	http.HandleFunc("/chat", web.Chat)
	http.HandleFunc("/ttt", web.T3)
	http.HandleFunc("/alog", web.AdminLogin)
	http.ListenAndServe(":9000", nil)
}
