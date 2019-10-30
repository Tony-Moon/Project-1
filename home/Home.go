package home

import (
	"fmt"
	"net/http"
	"text/template"
)

/*
Home doc
*/
func Home(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	home, _ := template.ParseFiles("home/template/home.html")
	home.Execute(res, req)

	fmt.Println("Navigated to Home")
}
