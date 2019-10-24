package web

import (
	"io"
	"net/http"
)

/*
HelloHTML doc
*/
func HelloHTML(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	//output := []byte("<html><body><h1>Hello There!</h1></body></html>")
	//response.Write(output)
	io.WriteString(response, `
	<DOCTYPE html>
	<html>
	<head>
		<title>My Page</title>
	</head>
	<body>
		<h2>Welcome to My Page</h2>
		<p>This is a test of a go server</p>
	</body>
	</html>
	`)
}
