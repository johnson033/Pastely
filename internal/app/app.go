package app

import (
	"Pastely/internal/hotkeys"
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) Init(ctx context.Context) {
	a.ctx = ctx

	// Setup the main key combos for the app.
	hotkeys.RegisterHotkey(hotkeys.HotKey{
		Name:        "Show Application",
		Combo:       "Ctrl+Shift+V",
		Description: "Show the application",
		Handler: func() error {
			// Implement the logic to show or hide the app
			runtime.WindowShow(a.ctx)
			runtime.WindowCenter(a.ctx)
			return nil // Replace with actual logic to show the app
		},
	})

	hotkeys.RegisterHotkey(hotkeys.HotKey{
		Name:        "Hide Application",
		Combo:       "Esc",
		Description: "Hide the application",
		Handler: func() error {
			// Implement the logic to hide the app
			runtime.WindowHide(a.ctx)
			return nil // Replace with actual logic to hide the app
		},
	})
}
