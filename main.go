package main

import (
	"embed"
	"fmt"
	"katalog/internal/db/migrations"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS
const dbFilePath string  = "./katalog.db"
func main() {
	// initialize database
	db, err := migrations.IntializeDatabase(dbFilePath)
	if err != nil {
		fmt.Printf("Could not initialize database schema: %w\n", err)
	}
	// Create an instance of the app structure
	app := NewApp(db)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "katalog",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
			app.ScannerService,
			app.DB,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

}

