package chat

import (
	"fmt"
	"net"
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

	ip := GetIP()
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
	ifaces, _ := net.Interfaces()
	var ip net.IP
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
		}
	}
	ips := ip.String()
	return ips
}
