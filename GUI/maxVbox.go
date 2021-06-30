package main

import (
	"fyne.io/fyne/v2"
)

type maxVbox struct {
}

func (d *maxVbox) MinSize(objects []fyne.CanvasObject) fyne.Size {

	w := float32(0)

	maxHeightChild := objects[0].MinSize().Height
	for _, o := range objects {
		childSize := o.MinSize()

		w += childSize.Width
		if childSize.Height > maxHeightChild {
			maxHeightChild = childSize.Height
		}
	}
	// color.Cyan("#########################")
	// color.Red("the min W is %2.2f", w)
	return fyne.NewSize(w, maxHeightChild)
}

func (d *maxVbox) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, containerSize.Height-d.MinSize(objects).Height)
	w := float32(0)
	for _, o := range objects {
		w += o.MinSize().Width
	}

	for _, o := range objects {
		// color.Red("current widget is %2.2f", o.MinSize().Width)

		o.Resize(fyne.NewSize(containerSize.Width/float32(len(objects)), o.MinSize().Height))
		o.Move(pos)
		if containerSize.Width > w {
			pos = pos.Add(fyne.NewPos(containerSize.Width/float32(len(objects)), 0))
		} else {
			pos = pos.Add(fyne.NewPos(o.MinSize().Width, 0))
		}

	}
}
