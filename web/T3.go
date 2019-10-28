package web

import (
	"fmt"
	"net/http"
	"text/template"
)

// var t3 = template.Must(template.New(".").Parse("template/t3.html"))

/*
T3 is Tic Tac Toe
*/
func T3(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Navigated to T3")

	res.Header().Set("Content-Type", "text/html")
	temp, _ := template.ParseFiles("web/template/t3.html")
	temp.Execute(res, req)
}
