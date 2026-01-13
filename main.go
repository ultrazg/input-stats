package main

import (
	"context"
	"embed"

	App "input-stats/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/windows/icon.ico
var trayIcon []byte

func main() {
	// Create an instance of the app structure
	app := App.NewApp()
	trayStartFunc, _ := App.CreateTray(app, trayIcon)

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "Input Stats",
		Width:     460,
		Height:    768,
		MinWidth:  460,
		MinHeight: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: func(ctx context.Context) {
			app.Ctx = ctx
			trayStartFunc()
		},
		Bind: []interface{}{
			app,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
