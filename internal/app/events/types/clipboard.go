package events

type clipboardEvents struct {
	Copy        EventType
	Paste       EventType
	ItemCreated EventType
	ItemUpdated EventType
	ItemDeleted EventType
}

var ClipboardEvents = clipboardEvents{
	Copy:        "clipboard.copy",
	Paste:       "clipboard.paste",
	ItemCreated: "clipboard.item.created",
	ItemUpdated: "clipboard.item.updated",
	ItemDeleted: "clipboard.item.deleted",
}
