package chat

import (
	"net/http"
	"text/template"
)

/*
UserLogin doc
*/
func UserLogin(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	chat, _ := template.ParseFiles("chat/template/ulog.html")

	req.ParseForm()
	chat.Execute(res, req)

}
