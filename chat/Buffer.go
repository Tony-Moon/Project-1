package chat

import (
	"fmt"
	"net"
	"net/http"
	"strings"
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
	ip := GetIP()
	fmt.Println(NickIPs[ip] + ": " + message)

	reBuffer.Execute(res, req)
}

// GetIP doc
func GetIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}
