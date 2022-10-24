package main

import (
	"simpleProjectOnWails/pkg/app"
)

func main() {
	a := app.NewApp()
	mWindow := app.NewMainWindow(a.App)
	mWindow.Run()
}
