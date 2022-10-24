package app

import (
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
)

type App struct {
	App *fyne.App
}

func NewApp() *App {
	a := fyneApp.NewWithID("TestChatByShagi")
	return &App{
		App: &a,
	}
}

func (a *App) SendNotification(title, content string) {
	ap := *a.App

	notification := fyne.NewNotification(title, content)
	ap.SendNotification(notification)

}
