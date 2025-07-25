package models

type Folder struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`       // Name of the history item
	Tags      []string `json:"tags"`       // Tags associated with the history item
	Color     string   `json:"color"`      // Color associated with the history item
	CreatedOn string   `json:"created_on"` // Time when the history item was created
}
