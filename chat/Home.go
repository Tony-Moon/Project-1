package chat

import (
	"net/http"
	"text/template"
)

/*
Home doc
*/
func Home(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	home, _ := template.ParseFiles("chat/template/home.html")
	home.Execute(res, req)
}
