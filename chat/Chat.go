package chat

import (
	"fmt"
	"net/http"
	"text/template"
)

/*
Chat doc
*/
func Chat(res http.ResponseWriter, req *http.Request) {
	var message string
	res.Header().Set("Content-Type", "text/html")
	chat, _ := template.ParseFiles("chat/template/chat.html")
	chat.Execute(res, req)

	message = "Navigated to Chat"
	fmt.Println(message)
	//MessageLog(db, message)
}
