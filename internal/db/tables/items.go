package database_tables

import (
	models "Pastely/pkg/models/clipboard"
	"database/sql"
	"time"
)

func InitItemsTable(DB *sql.DB) {
	// Create the items table if it doesn't exist
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			content TEXT NOT NULL UNIQUE,
			type TEXT NOT NULL,
			times_used INTEGER DEFAULT 0,
			folder_id INTEGER DEFAULT 0,
			created_on DATETIME DEFAULT CURRENT_TIMESTAMP,
			last_used DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (folder_id) REFERENCES folders(id)
		)
	`)
	if err != nil {
		panic(err)
	}
}

func CreateItem(DB *sql.DB, item models.Item) (models.Item, error) {
	if DB == nil {
		return models.Item{}, sql.ErrConnDone
	}

	result, err := DB.Exec(`
		INSERT INTO items (name, content, type)
		VALUES (?, ?, ?)
	`, item.Name, item.Content, item.Type)
	if err != nil {
		return models.Item{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Item{}, err
	}

	item.ID = id
	item.CreatedOn = time.Now()
	item.LastUsed = time.Now() // Set last used to now
	item.TimesUsed = 1         // Initial usage count

	return item, nil
}

func UpdateItem(DB *sql.DB, item models.Item) error {
	if DB == nil {
		return sql.ErrConnDone
	}

	_, err := DB.Exec(`
		UPDATE items
		SET name = ?, type = ?, times_used = ?, folder_id = ?, last_used = CURRENT_TIMESTAMP
		WHERE id = ?
	`, item.Name, item.Type, item.TimesUsed, item.FolderID, item.ID)
	return err
}

func GetItem(DB *sql.DB, id int64) (models.Item, error) {
	if DB == nil {
		return models.Item{}, sql.ErrConnDone
	}

	var item models.Item
	err := DB.QueryRow(`
		SELECT id, name, content, type, times_used, folder_id, created_on, last_used
		FROM items
		WHERE id = ?
	`, id).Scan(&item.ID, &item.Name, &item.Content, &item.Type, &item.TimesUsed, &item.FolderID, &item.CreatedOn, &item.LastUsed)
	if err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func FindItemByContent(DB *sql.DB, content string) (models.Item, error) {
	if DB == nil {
		return models.Item{}, sql.ErrConnDone
	}

	println("Searching for item with content:", content)

	var (
		item     models.Item
		lastUsed sql.NullTime
	)

	err := DB.QueryRow(`
		SELECT id, name, content, type, times_used, folder_id, created_on, last_used
		FROM items
		WHERE content = ?
	`, content).Scan(
		&item.ID,
		&item.Name,
		&item.Content,
		&item.Type,
		&item.TimesUsed,
		&item.FolderID,
		&item.CreatedOn,
		&lastUsed,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			println("Item not found with content:", content)
			return models.Item{}, nil // Not found is not an error
		}
		print("Error retrieving item by content:", err.Error())
		return models.Item{}, err
	}

	if lastUsed.Valid {
		item.LastUsed = lastUsed.Time
	}

	return item, nil
}

func ListItems(DB *sql.DB, offset, limit int) ([]models.Item, error) {
	if DB == nil {
		return nil, sql.ErrConnDone
	}

	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 100
	}

	rows, err := DB.Query(`
		SELECT id, name, content, type, times_used, folder_id, created_on, last_used
		FROM items
		LIMIT ? OFFSET ?
	`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		var lastUsed sql.NullTime

		err := rows.Scan(&item.ID, &item.Name, &item.Content, &item.Type, &item.TimesUsed, &item.FolderID, &item.CreatedOn, &lastUsed)
		if err != nil {
			return nil, err
		}

		if lastUsed.Valid {
			item.LastUsed = lastUsed.Time
		}

		items = append(items, item)
	}

	return items, nil
}
