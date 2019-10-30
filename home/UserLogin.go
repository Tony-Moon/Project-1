package home

import (
	"fmt"
	"net"
	"net/http"
	"text/template"
)

/*
UserLogin doc
*/
func UserLogin(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	chat, _ := template.ParseFiles("web/template/ulog.html")
	chat.Execute(res, req)

	ifas, _ := net.Interfaces()
	for _, ifa := range ifas {
		mac := ifa.HardwareAddr.String()
		fmt.Println(mac)
	}

	//users := NewUser(mac)
}
