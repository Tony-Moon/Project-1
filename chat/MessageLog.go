package chat

import (
	"database/sql"
)

/*
MessageLog doc
*/
func MessageLog(db *sql.DB, message string) {
	var mID int
	rows, _ := db.Query("SELECT mID FROM messages ORDER BY mID DESC LIMIT 1;")
	rows.Scan(&mID)
	mID++

	db.Exec("INSERT INTO messages VALUES ($1, $2);",
		mID, message)
}
