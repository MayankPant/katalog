package main

import (
	"context"
)

type App struct {
	 ctx	context.Context
	 FileService *FileService
	 DialogService *DialogService
}

func NewApp() *App {
    app := &App{
        FileService: NewFileService(),
    }
    app.DialogService = NewDialogService(app)
    return app
}


func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}