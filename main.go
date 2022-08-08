package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	// Create application with options
	err := wails.Run(&options.App{
		Title:            "swaypowerbtn",
		Width:            828,
		Height:           512,
		MinWidth:         828,
		MinHeight:        512,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 46, G: 52, B: 54, A: 1},
		Frameless:        true,
		AlwaysOnTop:      true,
		DisableResize:    true,
		Fullscreen:       false,
		LogLevel:         logger.TRACE,
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err)
	}
}
