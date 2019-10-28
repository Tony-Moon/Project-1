package web

import (
	"fmt"
	"net/http"
	"text/template"
)

/*
Admin doc
*/
func Admin(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	chat, _ := template.ParseFiles("web/template/admin.html")
	chat.Execute(res, req)

	fmt.Println("Navigated to Admin Page")
}