package events

type EventType string
type eventTypeValues struct {
	Clipboard clipboardEvents
}

var EventTypes = eventTypeValues{
	Clipboard: ClipboardEvents,
}
