package main

import (
	"fmt"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type fingerPos struct {
	index int
	pos   int
}

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
		subc.Move(fyne.NewPos(float32(newx-10), float32(newy-10)))
		subc.Resize(fyne.NewSize(20, 20))
		text := canvas.NewText(fmt.Sprintf("%d", i), color.RGBA{0, 0, 0, 40})
		text.Move(fyne.NewPos(float32(newx-5), float32(newy-10)))
		for _, v := range in {
			if v.pos == i {
				subc.FillColor = color.RGBA{59, 50, 75, 255}
				subc.Move(fyne.NewPos(float32(newx-20), float32(newy-20)))
				subc.Resize(fyne.NewSize(40, 40))
				text.Text = fmt.Sprintf("#%d", v.index)
				text.TextSize = 20
				text.Move(fyne.NewPos(float32(newx-12), float32(newy-15)))
				text.Color = color.White
			}
		}
		subcircles = append(subcircles, container.NewWithoutLayout(subc, text))
	}

	content := container.NewWithoutLayout(circ,
		subcircles[0], subcircles[1], subcircles[2], subcircles[3],
		subcircles[4], subcircles[5], subcircles[6], subcircles[7])

	return content
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")
	myWindow.Resize(fyne.NewSize(600, 300))

	// rect := canvas.NewRectangle(color.RGBA{59, 50, 75, 255})
	// rect.Resize(fyne.NewSize(300, 300))
	// rect.Move(fyne.NewPos(0, 0))

	fingers := []fingerPos{{0, 0}, {1, 4}}
	newcont := container.NewWithoutLayout(generateCircle(fingers))

	nnewcont := container.New(layout.NewHBoxLayout(), fingerList(), newcont)

	myWindow.SetContent(nnewcont)
	myWindow.ShowAndRun()
}
