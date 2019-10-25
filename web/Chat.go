package web

import (
	"fmt"
	"io"
	"net/http"
)

/*
Chat doc
*/
func Chat(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	io.WriteString(response, "chat.html")
	fmt.Println("Navigated to Chat")
}
