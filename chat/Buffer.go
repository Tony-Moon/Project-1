package chat

import (
	"fmt"
	"net/http"
	"text/template"
)

var nickIPs = make(map[string]string)
var messTable [10]string

/*
Buffer doc
*/
func Buffer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	buffer, _ := template.ParseFiles("chat/template/buffer.html")
	nickname := req.FormValue("name")

	ip := req.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = req.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = req.RemoteAddr
	}

	nickIPs[ip] = nickname
	fmt.Println(nickIPs[ip], "has joined the chat room!")
	buffer.Execute(res, req)
}

/*
ReBuffer doc
*/
func ReBuffer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	reBuffer, _ := template.ParseFiles("chat/template/rebuffer.html")
	message := req.FormValue("mess")

	ip := req.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = req.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = req.RemoteAddr
	}

	MessBoard(nickIPs[ip], message)
	reBuffer.Execute(res, req)
}

// MessBoard doc
func MessBoard(name string, mess string) {
	var count int
	for i := 0; i >= 9; i++ {
		if messTable[i] == " " {
			messTable[i] = (name + ": " + mess)
			fmt.Println(messTable[i])
			messTable[i+1] = " "
		} else {
			count++
		}
	}
	if count >= 9 {
		messTable[0] = " "
	}
}

// InitBoard doc
func InitBoard() {
	for i := 0; i >= 9; i++ {
		messTable[i] = " "
	}
}
