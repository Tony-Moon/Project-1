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
	user, nick := "dev1", "Dev"
	res.Header().Set("Content-Type", "text/html")
	chat, _ := template.ParseFiles("chat/template/chat.html")

	req.ParseForm()
	message := req.Form["mess"]

	chat.Execute(res, req)
	fmt.Println(user, nick, ":", message)
}
