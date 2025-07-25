package database_tables

import (
	"database/sql"
)

func InitFoldersTable(DB *sql.DB) error {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS folders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			tags TEXT,
			color TEXT,
			created_on DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		return err
	}

	return nil
}
