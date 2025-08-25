package main
import (
	"context"
	"database/sql"
	"katalog/internal/filesystem"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	 ctx			context.Context
	 DB				*sql.DB
	 ScannerService *filesystem.ScannerService
}

func NewApp(db *sql.DB) *App {
    app := &App{
        ScannerService: filesystem.NewScannerService(),
		DB: db,
    }
    return app
}


func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}


// SelectDirectory opens a dialog to select a directory and returns the path.
// This method is bindable to the frontend.
func (a *App) SelectDirectory() (string, error) {
	// Use the stored context when calling the dialog function.
	selection, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Please select your project directory",
	})

	if err != nil {
		// This will propagate the error to the frontend if the user cancels.
		return "", err
	}

	return selection, nil
}