package home

import (
	"net/http"
	"text/template"
)

/*
UserLogin doc
*/
func UserLogin(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	chat, _ := template.ParseFiles("home/template/ulog.html")
	chat.Execute(res, req)
}
