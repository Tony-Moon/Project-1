package chat

import (
	"fmt"
	"net/http"
	"text/template"
)

var nickname string

// NickIPs will store IP address and nicknames
var NickIPs = make(map[string]string)

/*
Buffer doc
*/
func Buffer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	buffer, _ := template.ParseFiles("chat/template/buffer.html")

	nickname = req.FormValue("name")

	ip := req.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = req.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = req.RemoteAddr
	}

	fmt.Println(ip)
	NickIPs[ip] = nickname
	fmt.Println(NickIPs[ip], "has joined the chat room!")

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

	fmt.Println(NickIPs[ip] + ": " + message)

	reBuffer.Execute(res, req)
}
