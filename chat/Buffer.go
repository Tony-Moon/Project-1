package chat

import (
	"fmt"
	"net/http"
	"text/template"
)

var nickname string

/*
Buffer doc
*/
func Buffer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	buffer, _ := template.ParseFiles("chat/template/buffer.html")

	nickname = req.FormValue("name")
	fmt.Println(nickname, "has joined the chat room!")

	buffer.Execute(res, req)

}

// Processor doc
func Processor(name string, message string) *User {
	return &User{
		name: name,
		mess: message,
	}
}

/*
ReBuffer doc
*/
func ReBuffer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	reBuffer, _ := template.ParseFiles("chat/template/rebuffer.html")

	message := req.FormValue("mess")
	current := Processor(nickname, message)
	fmt.Println(current.name + ": " + current.mess)

	reBuffer.Execute(res, req)
}

// User doc
type User struct {
	name string
	mess string
}
