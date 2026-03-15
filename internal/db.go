package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func initDatabase() {
	TABLE_NAME := "files"
	db, err := sql.Open("sqlite3", "./photofiler.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	sqlStmt := `
	DROP TABLE ` + TABLE_NAME + `;
	CREATE TABLE IF NOT EXISTS ` + TABLE_NAME + ` (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		type INTEGER,
		metadata_id INTEGER,
		dir TEXT,
		filename TEXT,
		filepath TEXT,
		metadata TEXT
		FOREIGN KEY(metadata_id) REFERENCES ` + TABLE_NAME + `(id)
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("Error creating table: %q: %s\n", err, sqlStmt)
		return
	}

}
