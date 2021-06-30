package main

import (
	"fmt"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type fixed300 struct {
}

func (d *fixed300) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(325, 300)
}

func (d *fixed300) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {

}

func generateCircle(in []fingerPos) fyne.CanvasObject {
	text1 := canvas.NewText("Text Object", color.RGBA{120, 0, 0, 255})
	text1.Alignment = fyne.TextAlignTrailing
	text1.TextStyle = fyne.TextStyle{Italic: true}

	circ := canvas.NewCircle(color.Transparent)
	circ.StrokeWidth = 2
	circ.StrokeColor = color.White
	circ.Move(fyne.NewPos(25, 25))
	circ.Resize(fyne.NewSize(250, 250))

	var subcircles []*fyne.Container
	for i := 0; i < 8; i++ {
		newx := 150 + math.Cos((360.0/8.0*float64(i))/180.0*math.Pi)*125
		newy := 150 + math.Sin((360.0/8.0*float64(i))/180.0*math.Pi)*125
		subc := canvas.NewCircle(color.White)
		subc.StrokeWidth = 2
		subc.StrokeColor = color.White
		subc.StrokeWidth = 2
		subc.StrokeColor = color.White
		subc.Move(fyne.NewPos(float32(newx-10), float32(newy-10)))
		subc.Resize(fyne.NewSize(20, 20))
		text := canvas.NewText(fmt.Sprintf("%d", i), color.RGBA{0, 0, 0, 40})
		text.Move(fyne.NewPos(float32(newx-5), float32(newy-10)))

		tooltipA := canvas.NewText("", text.Color)
		tooltipB := canvas.NewText("", text.Color)

		for _, v := range in {
			if v.pos == i {
				subc.Move(fyne.NewPos(float32(newx-20), float32(newy-20)))
				subc.Resize(fyne.NewSize(40, 40))

				text.Text = fmt.Sprintf("#%d", v.index)
				text.TextSize = 20
				text.Move(fyne.NewPos(float32(newx-12), float32(newy-15)))
				if v.active {
					subc.FillColor = color.RGBA{59, 50, 75, 255}
					text.Color = color.White
				} else {
					subc.FillColor = color.White
					text.Color = color.RGBA{59, 50, 75, 255}
				}

				tooltipA = canvas.NewText(fmt.Sprintf("%03.0f", v.A), color.White)
				tooltipB = canvas.NewText(fmt.Sprintf("%03.0f", v.B), color.White)

				tooltipA.TextSize = 8
				tooltipB.TextSize = 8

				tooltipA.Move(fyne.NewPos(float32(newx+25), float32(newy-10)))
				tooltipB.Move(fyne.NewPos(float32(newx+25), float32(newy)))
			}
		}
		subcircles = append(subcircles, container.NewWithoutLayout(subc, text, tooltipA, tooltipB))
	}

	content := container.NewWithoutLayout(circ,
		subcircles[0], subcircles[1], subcircles[2], subcircles[3],
		subcircles[4], subcircles[5], subcircles[6], subcircles[7])

	return fyne.NewContainerWithLayout(&fixed300{}, content)
}
