package web

import (
	"fmt"
	"net/http"
)

/*
FormSubmit doc
*/
func FormSubmit(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Welcome", request.FormValue("user"))
	fmt.Println("Your password is", request.FormValue("password"))
}
