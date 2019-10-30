package web

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
	chat, _ := template.ParseFiles("web/template/ulog.html")
	chat.Execute(res, req)

	fmt.Println("Navigated to User Login")
}
