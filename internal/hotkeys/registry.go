package hotkeys

import (
	"log"
	"strings"

	hook "github.com/robotn/gohook"
)

type HotKey struct {
	Name        string
	Combo       string
	Description string
	Handler     func() error
}

func RegisterHotkey(h HotKey) {
	keys := reorder(h.Combo)
	log.Printf("🎯  registering %q → %s", h.Combo, h.Name)

	hook.Register(hook.KeyDown, keys, func(e hook.Event) {
		if err := h.Handler(); err != nil {
			log.Printf("⚠️  handler for %s: %v", h.Name, err)
		} else {
			log.Printf("✅  handler for %s executed successfully", h.Name)
		}
	})
}

func reorder(combo string) []string {
	parts := strings.Split(strings.ToLower(combo), "+")

	// Trim whitespace, and sort the parts, shortest -> longest
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}

	// Sort the parts by length
	for i := 0; i < len(parts)-1; i++ {
		for j := i + 1; j < len(parts); j++ {
			if len(parts[i]) > len(parts[j]) {
				parts[i], parts[j] = parts[j], parts[i]
			}
		}
	}

	// Print the sorted parts for debugging
	log.Printf("🔄  reordered hotkey parts: %v", parts)

	return parts
}
