package clipboard

import (
	event "Pastely/internal/app/events"
	events "Pastely/internal/app/events/types"
	database "Pastely/internal/db"
	database_tables "Pastely/internal/db/tables"
	models "Pastely/pkg/models/clipboard"
	"regexp"
	"strings"
	"time"
)

var EMAIL_REGEX = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
var URL_REGEX = `^(https?|ftp)://[^\s/$.?#].[^\s]*$`

func isURL(content string) bool {
	// Use a regex to check if the content is a valid URL
	matched, _ := regexp.MatchString(URL_REGEX, content)
	return matched
}

func isEmail(content string) bool {
	// Use a regex to check if the content is a valid email address
	matched, _ := regexp.MatchString(EMAIL_REGEX, content)
	return matched
}

func isCode(content string) bool {
	return false // TODO:  Implement a more sophisticated check for code content
}

func detectType(content string) models.ItemType {
	switch {
	case isURL(content):
		return models.ItemTypes.URL
	case isEmail(content):
		return models.ItemTypes.Email
	case isCode(content):
		return models.ItemTypes.Code
	default:
		return models.ItemTypes.Text
	}
}

func process(content string) {
	// Trim the content to avoid leading/trailing whitespace
	content = strings.TrimSpace(content)

	// Check if the item already exists in the database.
	item, _ := database_tables.FindItemByContent(database.DB, content)
	if item.ID != 0 {
		// Increment the usage count, and update the last used time
		item.TimesUsed++
		item.LastUsed = time.Now()
		database_tables.UpdateItem(database.DB, item)
		event.EmitEvent(events.ClipboardEvents.ItemUpdated, item)
		return
	}

	// If the item does not exist, create a new item.
	itemType := detectType(content)

	item = models.Item{
		Name:      "",
		Content:   content,
		Type:      itemType,
		Tags:      []string{}, // Tags can be added later
		TimesUsed: 1,          // Initial usage count
	}

	item, err := database_tables.CreateItem(database.DB, item)
	if err != nil {
		println("Error creating item: ", err.Error())
		return
	}
	event.EmitEvent(events.ClipboardEvents.ItemCreated, item)
}
