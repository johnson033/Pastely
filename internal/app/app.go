package app

import (
	event "Pastely/internal/app/events"
	"Pastely/internal/hotkeys"
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	Ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) Init(ctx context.Context) {
	a.Ctx = ctx

	event.Init(a.Ctx)

	// Setup the main key combos for the app.
	hotkeys.RegisterHotkey(hotkeys.HotKey{
		Name:        "Show Application",
		Combo:       "Alt+V",
		Description: "Show the application",
		Handler: func() error {
			// Implement the logic to show or hide the app
			println("Showing application")
			runtime.WindowShow(a.Ctx)
			runtime.WindowCenter(a.Ctx)
			return nil // Replace with actual logic to show the app
		},
	})

	hotkeys.RegisterHotkey(hotkeys.HotKey{
		Name:        "Hide Application",
		Combo:       "Esc",
		Description: "Hide the application",
		Handler: func() error {
			// Implement the logic to hide the app
			println("Hiding application")
			runtime.WindowHide(a.Ctx)
			return nil // Replace with actual logic to hide the app
		},
	})
}
