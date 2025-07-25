package event

import (
	events "Pastely/internal/app/events/types"
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var appContext context.Context

func Init(ctx context.Context) {
	appContext = ctx
}

func EmitEvent(eventName events.EventType, data interface{}) {
	runtime.EventsEmit(appContext, string(eventName), data)
}
