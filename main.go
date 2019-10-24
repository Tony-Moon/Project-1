package main

import (
	"net/http"

	"github.com/Tony-Moon/Project-1/web"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/hello", web.Hello)
	http.HandleFunc("/hellohtml", web.HelloHTML)
	http.HandleFunc("/formsubmit", web.FormSubmit)
	http.HandleFunc("/template", web.HelloTempHandler)
	http.ListenAndServe(":9000", nil)
}
