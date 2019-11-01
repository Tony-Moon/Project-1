package chat

import (
	"fmt"
	"net/http"
	"text/template"
)

/*
UserLogin doc
*/
func UserLogin(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	chat, _ := template.ParseFiles("chat/template/ulog.html")

	name := req.Form["name"]
	pass := req.Form["pass"]
	req.ParseForm()
	chat.Execute(res, req)

	curr := currentUser{
		username: name,
		nickname: "Hal",
		userrank: false,
	}
	fmt.Println(curr)
	fmt.Println(pass)

}

type currentUser struct {
	username []string
	nickname string
	userrank bool
}
