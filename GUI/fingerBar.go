package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type fingerBarContainer struct {
}

func (d *fingerBarContainer) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	w = 200
	h = float32(len(objects)/3) * (objects[1].MinSize().Height)
	return fyne.NewSize(w, h)
}

func (d *fingerBarContainer) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, 0)
	for i, o := range objects {
		size := fyne.NewSize(o.MinSize().Width, objects[1].MinSize().Height)
		if i%3 == 0 {
			pos.X = 0
			pos.Y = float32(i/3) * objects[1].MinSize().Height
		}
		if i%3 == 1 {
			size.Width = 45
		}
		if i%3 == 2 {
			size.Width = containerSize.Width - objects[0].MinSize().Width - 45
		}
		o.Resize(size)
		o.Move(pos)

		pos = pos.Add(fyne.NewPos(size.Width, 0))
	}
}

func fingerBar(in fingerPos) fyne.CanvasObject {
	A, B := binding.NewInt(), binding.NewInt()
	A.Set(in.A)
	B.Set(in.B)
	label1 := canvas.NewText(fmt.Sprintf("#%d A", in.index), color.Black)
	value1 := widget.NewSlider(0, 100)
	value1.Value = 50
	enter1 := widget.NewEntryWithData(binding.IntToString(A))
	enter1.PlaceHolder = "50"
	enter1.Validator = nil
	label2 := canvas.NewText(fmt.Sprintf("#%d B", in.index), color.Black)
	value2 := widget.NewSlider(0, 100)
	value2.Value = 0
	enter2 := widget.NewEntryWithData(binding.IntToString(B))
	enter2.PlaceHolder = "0"
	enter2.Validator = nil

	return container.New(&fingerBarContainer{}, label1, enter1, value1, label2, enter2, value2)
}
