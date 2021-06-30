package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func fingerBar(in fingerPos) fyne.CanvasObject {
	A, B := binding.NewFloat(), binding.NewFloat()
	A.Set(in.A)
	B.Set(in.B)
	label1 := canvas.NewText(fmt.Sprintf("#%d A", in.index), color.Black)
	value1 := widget.NewSliderWithData(0, 100, A)
	enter1 := widget.NewEntryWithData(binding.FloatToString(A))
	enter1.PlaceHolder = "50"
	label2 := canvas.NewText(fmt.Sprintf("#%d B", in.index), color.Black)
	value2 := widget.NewSliderWithData(0, 100, B)
	enter2 := widget.NewEntryWithData(binding.FloatToStringWithFormat(B, "%f"))
	enter2.PlaceHolder = "0"
	value2.SetValue(0)
	row1 := container.New(layout.NewHBoxLayout(), label1, enter1)
	row2 := container.New(layout.NewHBoxLayout(), label2, enter2)

	grid := container.New(layout.NewFormLayout(), row1, value1, row2, value2)
	cont := container.New(layout.NewVBoxLayout(), widget.NewSeparator(), grid, widget.NewSeparator())

	return cont
}
