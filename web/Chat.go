package web

import (
	"fmt"
	"net/http"
	"text/template"
)

/*
Chat doc
*/
func Chat(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	chat, _ := template.ParseFiles("web/template/chat.html")
	chat.Execute(res, req)

	fmt.Println("Navigated to Chat")
}
