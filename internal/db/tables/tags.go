package database_tables

import "database/sql"

func InitTagsTable(DB *sql.DB) error {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS tags (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			color TEXT NOT NULL,
			created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	return err
}
