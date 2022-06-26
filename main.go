package main

// Imports
import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	//fmt.Println("Hello World!")

	application := app.New()
	window := application.NewWindow("Pomodoro Timer")
	window.Resize(fyne.NewSize(450, 300))

	title := canvas.NewText("Let's Focus!", color.White)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	hBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer())
	vBox := container.New(layout.NewVBoxLayout(), title, hBox)

	window.SetContent(vBox)
	window.ShowAndRun()
}
