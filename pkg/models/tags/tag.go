package models

import "time"

type Tag struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`       // Name of the tag
	Color     string    `json:"color"`      // Color associated with the tag
	CreatedOn time.Time `json:"created_on"` // Time when the item was created
}
