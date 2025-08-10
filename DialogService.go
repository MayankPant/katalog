package main

import (
    "github.com/wailsapp/wails/v2/pkg/runtime"
)

type DialogService struct{
    app *App
}

func NewDialogService(app *App) *DialogService {
    return &DialogService{
        app: app,
    }
}

func (dialogService *DialogService) SelectDirectory() (string, error) {
    dir, err := runtime.OpenDirectoryDialog(dialogService.app.ctx, runtime.OpenDialogOptions{
    Title:                      "Select a target folder",
    DefaultDirectory:           "C://",
    ShowHiddenFiles:            true,
    CanCreateDirectories:       true,
    ResolvesAliases:            true,
    TreatPackagesAsDirectories: true,
})
    if err != nil {
        return "", err
    }
    return dir, nil
}
