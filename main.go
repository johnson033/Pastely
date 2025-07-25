package main

import (
	"Pastely/internal/app"
	bindings_clipboard "Pastely/internal/app/bindings/clipboard"
	bindings_folder "Pastely/internal/app/bindings/folder"
	bindings_tag "Pastely/internal/app/bindings/tag"
	"Pastely/internal/core/clipboard"
	database "Pastely/internal/db"
	"Pastely/internal/hotkeys"

	// events_hotkeys "Pastely/internal/events/hotkeys"
	// services_app "Pastely/services/app"
	// services_clipboard "Pastely/services/clipboard"
	// services_clipboard_bindings "Pastely/services/clipboard/bindings"
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var assets embed.FS

type WailsService interface {
	Init(ctx context.Context)
}

func main() {
	// Initialize the database...
	err := database.Init("Pastely", "pastely")
	if err != nil {
		println("Error initializing database:", err.Error())
		return
	}

	app := app.NewApp()
	clipboardWatcher := clipboard.NewWatcher()
	services := []WailsService{
		app,
		clipboardWatcher,
		&bindings_clipboard.ClipboardItem{},
		&bindings_folder.Folder{},
		&bindings_tag.Tag{},
	}

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Pastely",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:         false,
		HideWindowOnClose: true,
		StartHidden:       true,
		BackgroundColour:  &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			for _, service := range services {
				service.Init(ctx)
			}

			hotkeys.Listen(ctx)
		},
		Bind: convertToInterfaces(services),
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func convertToInterfaces(s []WailsService) []interface{} {
	result := make([]interface{}, len(s))
	for i, svc := range s {
		result[i] = svc
	}
	return result
}
