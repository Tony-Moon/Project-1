package web

import (
	"expvar"
	"net/http"
	"text/template"
)

const hellotemplate = `
<!DOCTYPE html>
<html>
	<head>
		<title>Template page</title>
	</head>
	<body>
		<h1>Hello, {{.Name}}!</h1>
	</body>
</html>
`

var hellotmpl = template.Must(template.New(".").Parse(hellotemplate))
var myCount = expvar.NewInt("my.count")
var myStatus = expvar.NewString("my.status")

/*
HelloTempHandler doc
*/
func HelloTempHandler(response http.ResponseWriter, request *http.Request) {
	myCount.Add(1)
	myStatus.Set("Good")
	hellotmpl.Execute(response, map[string]interface{}{
		"Name": "Bob",
	})
}
