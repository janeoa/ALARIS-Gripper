package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func fingerList() fyne.CanvasObject {
	var fingers []fyne.CanvasObject
	for i := 0; i < 8; i++ {
		fingers = append(fingers, fingerInfoItem(i))
	}
	return container.New(layout.NewVBoxLayout(), fingers...)
}

func fingerInfoItem(id int) fyne.CanvasObject {
	name := widget.NewLabel(fmt.Sprintf("Finger %d Position:", id))
	pos := widget.NewLabel(fmt.Sprintf("%d", id))
	return container.New(layout.NewHBoxLayout(), name, pos)
}

func makeCircle() fyne.CanvasObject {
	circle := canvas.NewCircle(color.Black)
	circle.StrokeColor = color.Gray{0x99}
	circle.StrokeWidth = 5
	return circle
}

func mainLayout() fyne.CanvasObject {
	return container.New(layout.NewHBoxLayout(), fingerList(), container.NewCenter(makeCircle()))
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")
	myWindow.Resize(fyne.NewSize(400, 300))

	myWindow.SetContent(mainLayout())
	myWindow.ShowAndRun()
}
