package main

import (
	"net/http"

	"github.com/Tony-Moon/Project-1/chat"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/home", chat.Home)
	http.HandleFunc("/ulog", chat.UserLogin)
	http.HandleFunc("/buffer", chat.Buffer)
	http.HandleFunc("/rebuffer", chat.ReBuffer)
	http.ListenAndServe("ec2-3-19-77-195.us-east-2.compute.amazonaws.com:80", nil)
}
