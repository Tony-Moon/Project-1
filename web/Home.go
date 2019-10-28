package web

import (
	"fmt"
	"net/http"
	"text/template"
)

/*
Home doc
*/
func Home(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	home, _ := template.ParseFiles("web/template/home.html")
	home.Execute(response, request)

	fmt.Println("Navigated to Home")
}
