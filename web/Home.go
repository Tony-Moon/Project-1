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
	home, _ := template.ParseFiles("home.html")
	home.Execute(response, request)

	//response.Header().Set("Content-Type", "text/html")
	//io.WriteString(response, "home.html")
	fmt.Println("Navigated to Home")
}
