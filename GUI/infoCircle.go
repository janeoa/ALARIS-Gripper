package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/go-gl/mathgl/mgl32"
)

type fixed300 struct {
	// canvas fyne.CanvasObject
}

func (d *fixed300) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(325, 300)
}

func (d *fixed300) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {

}

func generateCircle() *fyne.Container {
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

		for _, v := range gripper.finger {
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

				tooltipA = canvas.NewText(fmt.Sprintf("%02d", v.A), color.White)
				tooltipB = canvas.NewText(fmt.Sprintf("%02d", v.B), color.White)

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

	// var currx, curry, goalx, goaly float64
	for _, v := range gripper.finger {
		if v.pos != v.newPos {
			arrowColor := color.RGBA{255, 0, 0, 50}
			newPosIsValid := validateNewPos()
			log.Printf("the newPosIsValid is %v", newPosIsValid)
			if newPosIsValid {
				arrowColor = color.RGBA{0, 255, 0, 50}
			}
			// log.Printf("the newPosIsValid is %v", newPosIsValid)

			arrowbone := canvas.NewLine(arrowColor)
			arrowbone.Position1.X = float32(150 + math.Cos((360.0/8.0*float64(v.pos))/180.0*math.Pi)*125)
			arrowbone.Position1.Y = float32(150 + math.Sin((360.0/8.0*float64(v.pos))/180.0*math.Pi)*125)

			arrowbone.Position2.X = float32(150 + math.Cos((360.0/8.0*float64(v.newPos))/180.0*math.Pi)*125)
			arrowbone.Position2.Y = float32(150 + math.Sin((360.0/8.0*float64(v.newPos))/180.0*math.Pi)*125)

			// log.Printf("%v", arrowbone.Position2)

			p1 := mgl32.Vec2{arrowbone.Position1.X, arrowbone.Position1.Y}
			p2 := mgl32.Vec2{arrowbone.Position2.X, arrowbone.Position2.Y}

			vec1 := p2.Sub(p1)
			nvec1 := vec1.Normalize()

			arrowbone.Position2.X -= nvec1.X() * 20
			arrowbone.Position2.Y -= nvec1.Y() * 20

			arrowbone.Position1.X += nvec1.X() * 30
			arrowbone.Position1.Y += nvec1.Y() * 30

			degInPi := 160 * math.Pi / 180
			px := float32(nvec1.X()*float32(math.Cos(degInPi)) - nvec1.Y()*float32(math.Sin(degInPi)))
			py := float32(nvec1.X()*float32(math.Sin(degInPi)) + nvec1.Y()*float32(math.Cos(degInPi))) // x*sn + y*cs
			pp := mgl32.Vec2{px, py}

			arrowSid1 := canvas.NewLine(arrowColor)
			arrowSid1.Position1.X = arrowbone.Position2.X
			arrowSid1.Position1.Y = arrowbone.Position2.Y

			arrowSid1.Position2.X = arrowSid1.Position1.X + pp.X()*10
			arrowSid1.Position2.Y = arrowSid1.Position1.Y + pp.Y()*10

			degInPi = 200 * math.Pi / 180
			px = float32(nvec1.X()*float32(math.Cos(degInPi)) - nvec1.Y()*float32(math.Sin(degInPi)))
			py = float32(nvec1.X()*float32(math.Sin(degInPi)) + nvec1.Y()*float32(math.Cos(degInPi))) // x*sn + y*cs
			pp = mgl32.Vec2{px, py}

			arrowSid2 := canvas.NewLine(arrowColor)
			arrowSid2.Position1.X = arrowbone.Position2.X
			arrowSid2.Position1.Y = arrowbone.Position2.Y

			arrowSid2.Position2.X = arrowSid1.Position1.X + pp.X()*10
			arrowSid2.Position2.Y = arrowSid1.Position1.Y + pp.Y()*10

			content = container.NewWithoutLayout(arrowbone, content, arrowSid1, arrowSid2)
		}
	}

	return container.New(&fixed300{}, content)
}

func validateNewPos() bool {
	for _, v := range generateArrows() {
		if v > 1 {
			return false
		}
	}

	return true
}

func generateArrows() []int {
	mapp := []int{0, 0, 0, 0, 0, 0, 0, 0}
	for _, v := range gripper.finger {
		if v.pos == v.newPos {
			mapp[v.pos]++
			continue
		}
		index := v.pos
		mapp[index]++
		for {
			if is_next_on_right(v.pos, v.newPos) {
				index++
				if index > 7 {
					index -= 8
				}
			} else {
				index--
				if index < 0 {
					index += 8
				}
			}

			mapp[index]++
			if index == v.newPos {
				break
			}

			// fmt.Printf("%v\n", mapp)
		}

	}
	return mapp
}

func dirCrossesZero(pos int, newPos int) bool {
	return (is_next_on_right(pos, newPos) && pos > newPos) || (!is_next_on_right(pos, newPos) && newPos > pos)
}

func is_next_on_right(prev int, curr int) bool {
	is_right := false

	delta := curr - prev
	if delta > 0 && delta < 4 {
		is_right = true
	}
	if delta < -4 && delta > -8 {
		is_right = true
	}
	return is_right
}
