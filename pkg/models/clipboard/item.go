package models

import (
	"time"
)

type ItemType string
type itemTypeValues struct {
	Text  ItemType
	Code  ItemType
	URL   ItemType
	Email ItemType
	File  ItemType
	Image ItemType
	Video ItemType
}

var ItemTypes = itemTypeValues{
	Text:  "text",
	Code:  "code",
	URL:   "url",
	Email: "email",
	File:  "file",
	Image: "image",
	Video: "video",
}

type Item struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	Type      ItemType  `json:"type"`       // e.g., text, image, file
	Tags      []string  `json:"tags"`       // Tags associated with the item
	TimesUsed int       `json:"times_used"` // Number of times the item has been used
	FolderID  int64     `json:"folder_id"`  // ID of the folder this item belongs to
	CreatedOn time.Time `json:"created_on"` // Time when the item was created
	LastUsed  time.Time `json:"last_used"`  // Time when the item was last used
}
