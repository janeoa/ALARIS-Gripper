package main

import "fyne.io/fyne/v2"

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
	return fyne.NewSize(w, maxHeightChild)
}

func (d *maxVbox) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, containerSize.Height-d.MinSize(objects).Height)

	for _, o := range objects {
		// size := o.MinSize()
		o.Resize(fyne.NewSize(containerSize.Width/float32(len(objects)), o.MinSize().Height))
		o.Move(pos)

		pos = pos.Add(fyne.NewPos(containerSize.Width/float32(len(objects)), 0))
	}
}
