package web

import (
	"fmt"
	"net/http"
)

/*
Hello doc
*/
func Hello(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	output := []byte("Hello " + name)
	fmt.Println(name, "says hello!")
	response.Write(output)
}
