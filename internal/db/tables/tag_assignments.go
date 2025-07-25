package database_tables

import "database/sql"

func InitTagAssignmentsTable(DB *sql.DB) error {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS tag_assignments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			tag_id INTEGER NOT NULL,
			item_id INTEGER NOT NULL,
			FOREIGN KEY (tag_id) REFERENCES tags(id),
			FOREIGN KEY (item_id) REFERENCES clipboard_items(id)
		);
	`)
	return err
}
