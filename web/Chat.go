package web

import (
	"fmt"
	"net/http"
	"text/template"
)

/*
Chat doc
*/
func Chat(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	chat, _ := template.ParseFiles("web/template/chat.html")
	chat.Execute(response, request)

	fmt.Println("Navigated to Chat")
}
