package models

import "time"

type EntityType string
type entityTypeValues struct {
	Folder EntityType
	Item   EntityType
}

var EntityTypes = entityTypeValues{
	Folder: "folder",
	Item:   "item",
}

type TagAssignment struct {
	ID         int64      `json:"id"`
	TagID      int64      `json:"tag_id"`      // ID of the tag
	EntityID   int64      `json:"entity_id"`   // ID of the entity (folder or item)
	EntityType EntityType `json:"entity_type"` // Type of the entity (folder or item)
	CreatedOn  time.Time  `json:"created_on"`  // Time when the tag was assigned
}
