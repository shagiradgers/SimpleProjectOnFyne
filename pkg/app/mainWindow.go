package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"simpleProjectOnWails/pkg/messages"
	"strings"
)

var chat []messages.Message

type MainWindow struct {
	App    *fyne.App
	Window fyne.Window
	Width  float32
	Height float32
}

func NewMainWindow(a *fyne.App) *MainWindow {
	app := *a
	m := &MainWindow{
		App:    a,
		Window: app.NewWindow("Chat"),
		Width:  300,
		Height: 200,
	}

	m.setMainWindowSize()
	m.setMainWindowContent()
	return m
}

func (m *MainWindow) setMainWindowSize() {
	m.Window.Resize(fyne.Size{
		Width:  m.Width,
		Height: m.Height,
	})
	m.Window.SetFixedSize(true)
	m.Window.CenterOnScreen()
}

func (m *MainWindow) setMainWindowContent() {
	// input text
	inputMessage := widget.NewEntry()
	inputMessage.SetPlaceHolder("Enter text...")

	// input nickname
	inputLogin := widget.NewEntry()
	inputLogin.SetPlaceHolder("Enter nickname...")

	// list with messages
	list := widget.NewList(
		func() int {
			return len(chat)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("nil")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(chat[i].Text)
		},
	)

	// input button
	button := widget.NewButton("Send",
		func() {
			// проверка на пустой ввод
			if len(strings.Replace(inputMessage.Text, " ", "", -1)) != 0 {
				chat = append(chat, messages.NewMessage(inputLogin.Text, inputMessage.Text))
				m.sendMessage()
			}
		},
	)

	// layout with button and input
	inputLayout := container.NewVBox(
		inputLogin,
		inputMessage,
		button,
	)

	// layout with input components and list with images
	content := container.NewGridWithColumns(2,
		list,
		inputLayout,
	)

	m.Window.Content()
	m.Window.SetContent(content)
}

func (m *MainWindow) Run() {
	m.Window.ShowAndRun()
}

func (m *MainWindow) sendMessage() {
	m.setMainWindowContent()
}
