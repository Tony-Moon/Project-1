package main

import (
	"net/http"

	"github.com/Tony-Moon/Project-1/web"
)

func main() {
	// router := mux.NewRouter()
	// router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("web/template").HTTPBox()))

	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/home", web.Home)
	http.HandleFunc("/chat", web.Chat)
	http.HandleFunc("/ttt", web.T3)
	http.HandleFunc("/alog", web.AdminLogin)
	http.HandleFunc("/admin", web.Admin)

	http.ListenAndServe(":9000", nil)
}
