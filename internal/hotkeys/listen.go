package hotkeys

import (
	"context"
	"log"

	hook "github.com/robotn/gohook"
)

func Listen(ctx context.Context) {
	go func() {
		log.Println("🎧 Listening for global hotkeys...")

		// Start the hook event stream
		eventStream := hook.Start()

		// Wait for context cancel or continue processing
		select {
		case <-ctx.Done():
			log.Println("🛑 Stopping hotkey listener due to context cancellation.")
			hook.End()
		case <-hook.Process(eventStream):
			log.Println("⚠️  Hotkey processing ended unexpectedly.")
		}
	}()
}
