package main

// Imports
import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"Pomodoro/WindowComponents"
)

func main() {
	application := app.New()
	window := application.NewWindow("Pomodoro Timer")
	window.Resize(fyne.NewSize(450, 300))

	// Window Design
	title := canvas.NewText("Let's Focus!", color.White)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	// Sets Container
	amountOfSetsLabel := canvas.NewText("How Many Sets?", color.White)
	setsInput := widget.NewEntry()
	setsInput.SetPlaceHolder("00")
	setsInput.Resize(fyne.NewSize(100, 100))
	setsContainer := container.NewHBox(amountOfSetsLabel, setsInput)

	// Work Timer container
	amountOfWorkTimelabel := canvas.NewText("How long would you like to work?", color.White)
	workTimeInput := widget.NewEntry()
	workTimeInput.SetPlaceHolder("00")
	workTimeInput.Resize(fyne.NewSize(100, 100))
	workTimeContainer := container.NewHBox(amountOfWorkTimelabel, workTimeInput)

	//Break Timer Container
	amountOfBreakTimeLabel := canvas.NewText("How long would you like to take breaks?", color.White)
	breakTimeInput := widget.NewEntry()
	breakTimeInput.SetPlaceHolder("00")
	breakTimeContainer := container.NewHBox(amountOfBreakTimeLabel, breakTimeInput)

	// Set component entry values
	

	hBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer())
	vBox := container.New(layout.NewVBoxLayout(), title, hBox, widget.NewSeparator(), setsContainer,
		workTimeContainer, breakTimeContainer )

	window.SetContent(vBox)
	window.ShowAndRun()
}
