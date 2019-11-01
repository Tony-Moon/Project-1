package main

import (
	"net/http"

	"github.com/Tony-Moon/Project-1/chat"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/home", chat.Home)
	http.HandleFunc("/ulog", chat.UserLogin)
	http.ListenAndServe(":9000", nil)
}
