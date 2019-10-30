package admin

import (
	"fmt"
	"net/http"
	"text/template"
)

/*
AdminLogin doc
*/
func AdminLogin(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	chat, _ := template.ParseFiles("web/template/alog.html")
	chat.Execute(res, req)

	fmt.Println("Navigated to Admin Login")
}
