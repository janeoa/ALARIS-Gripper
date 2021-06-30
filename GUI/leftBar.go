package main

import (
	"fmt"

	"fyne.io/fyne/v2"
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
