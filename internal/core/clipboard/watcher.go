package clipboard

import (
	"context"
	"log"
	"time"

	"github.com/atotto/clipboard"
)

type Watcher struct {
	last string
	ctx  context.Context
}

// Listen creates a new clipboard watcher with context.
func NewWatcher() *Watcher {
	return &Watcher{}
}

// Start begins watching the clipboard and logs new content.
func (w *Watcher) Init(ctx context.Context) {
	w.ctx = ctx
	go func() {
		ticker := time.NewTicker(300 * time.Millisecond)
		defer ticker.Stop()

		log.Println("📋 Clipboard watcher started.")

		for {
			select {
			case <-w.ctx.Done():
				log.Println("🛑 Clipboard watcher stopped.")
				return
			case <-ticker.C:
				w.checkClipboard()
			}
		}
	}()
}

// checkClipboard checks for clipboard changes and prints them.
func (w *Watcher) checkClipboard() {
	content, err := clipboard.ReadAll()
	if err != nil {
		log.Println("⚠️ Failed to read clipboard:", err)
		return
	}

	if content != "" && content != w.last {
		w.last = content
		log.Println("📥 Clipboard changed:", content)
		process(content) // Process the new clipboard content
	}
}
