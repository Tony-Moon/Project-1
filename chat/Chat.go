package web

import (
	"database/sql"
	"net/http"
	"text/template"
)

/*
Chat doc
*/
func Chat(res http.ResponseWriter, req *http.Request, db *sql.DB) {
	var message string
	res.Header().Set("Content-Type", "text/html")
	chat, _ := template.ParseFiles("web/template/chat.html")
	chat.Execute(res, req)

	message = "Navigated to Chat"
	MessageLog(db, message)
}
