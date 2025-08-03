package main

import (
	"context"
)

type App struct {
	 ctx	context.Context
	 FileService *FileService
}

func NewApp() *App {
	return &App{
		FileService: NewFileService(),
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}