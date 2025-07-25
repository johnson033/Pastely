package database_tables

import (
	"database/sql"
)

func InitItemsTable(DB *sql.DB) {
	// Create the items table if it doesn't exist
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			content TEXT NOT NULL,
			type TEXT NOT NULL,
			times_used INTEGER DEFAULT 0,
			folder_id INTEGER,
			created_on DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (folder_id) REFERENCES folders(id)
		)
	`)
	if err != nil {
		panic(err)
	}
}
