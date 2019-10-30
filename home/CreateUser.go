package home

import (
	"fmt"
	"net"
	"net/http"
	"text/template"
)

/*
CreateUser doc
*/
func CreateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	home, _ := template.ParseFiles("home/template/create.html")
	home.Execute(res, req)

	var macaddress string
	ifas, _ := net.Interfaces()
	for _, ifa := range ifas {
		macaddress = ifa.HardwareAddr.String()
	}

	username := req.FormValue("username")
	password := req.FormValue("password")
	nickname := req.FormValue("nickname")
	AddTo(macaddress, username, password, nickname, false)

	fmt.Println(macaddress)
	fmt.Println("Navigated to Create User")
}
