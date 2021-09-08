package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type fingerBarContainer struct {
}

func (d *fingerBarContainer) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	w = 200
	h = float32(len(objects)/2) * (objects[0].MinSize().Height)
	return fyne.NewSize(w, h)
}

func generateFingerBarList() *fyne.Container {
	var fingerWidged []fyne.CanvasObject
	for _, v := range gripper.finger {
		if v.active {
			fingerWidged = append(fingerWidged, fingerBar(v))
		}
	}
	return container.New(layout.NewVBoxLayout(), fingerWidged...)
}

func (d *fingerBarContainer) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, 0)
	for i, o := range objects {
		size := fyne.NewSize(o.MinSize().Width, objects[0].MinSize().Height)
		if i%2 == 0 {
			pos.X = 0
			pos.Y = float32(i/2) * objects[0].MinSize().Height
		}
		// if i%3 == 1 {
		// 	size.Width = 45
		// }
		if i%2 == 1 {
			size.Width = containerSize.Width - objects[0].MinSize().Width
		}
		o.Resize(size)
		o.Move(pos)

		pos = pos.Add(fyne.NewPos(size.Width, 0))
	}
}

func fingerBar(in fingerPos) fyne.CanvasObject {
	label1 := widget.NewLabelWithData(binding.FloatToStringWithFormat(in.A, "%02.0f"))
	value1 := widget.NewSliderWithData(0, 100, in.A)
	value1.Value = 50
	// enter1 := widget.NewEntryWithData(binding.IntToString(A))
	// enter1.PlaceHolder = "50"
	// enter1.Validator = nil
	label2 := widget.NewLabelWithData(binding.FloatToStringWithFormat(in.B, "%02.0f"))
	value2 := widget.NewSliderWithData(0, 100, in.B)
	value2.Value = 0
	// enter2 := widget.NewEntryWithData(binding.IntToString(B))
	// enter2.PlaceHolder = "0"
	// enter2.Validator = nil

	// return container.New(&fingerBarContainer{}, label1, enter1, value1, label2, enter2, value2)
	return container.New(&fingerBarContainer{}, label1, value1, label2, value2)
}
