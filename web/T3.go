package web

import (
	"fmt"
	"io"
	"net/http"
)

/*
T3 is Tic Tac Toe
*/
func T3(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	io.WriteString(response, "t3.html")
	fmt.Println("Navigated to T3")
}
