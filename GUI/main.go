package main

import (
	"fmt"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	fatih "github.com/fatih/color"
)

type fingerPos struct {
	index  int
	pos    int
	active bool
	A      float64
	B      float64
}

type TapableCircle struct {
	*canvas.Circle

	OnTapped func() `json:"-"`
	active   bool
}

func (obj *TapableCircle) Tapped(*fyne.PointEvent) {
	if obj.OnTapped != nil {
		obj.OnTapped()
	}
}

func (mc *TapableCircle) TappedSecondary(*fyne.PointEvent) {}

func NewTapCircle(color color.Color, tapped func()) *TapableCircle {
	return &TapableCircle{
		canvas.NewCircle(color),
		tapped,
		false,
	}
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
		subc := NewTapCircle(color.White, subCircleTab)
		subc.StrokeWidth = 2
		subc.StrokeColor = color.White
		subc.Circle.StrokeWidth = 2
		subc.Circle.StrokeColor = color.White
		subc.Circle.Move(fyne.NewPos(float32(newx-10), float32(newy-10)))
		subc.Circle.Resize(fyne.NewSize(20, 20))
		text := canvas.NewText(fmt.Sprintf("%d", i), color.RGBA{0, 0, 0, 40})
		text.Move(fyne.NewPos(float32(newx-5), float32(newy-10)))

		tooltipA := canvas.NewText("", text.Color)
		tooltipB := canvas.NewText("", text.Color)

		for _, v := range in {
			if v.pos == i {
				subc.active = true
				subc.Circle.Move(fyne.NewPos(float32(newx-20), float32(newy-20)))
				subc.Circle.Resize(fyne.NewSize(40, 40))

				text.Text = fmt.Sprintf("#%d", v.index)
				text.TextSize = 20
				text.Move(fyne.NewPos(float32(newx-12), float32(newy-15)))
				if v.active {
					subc.Circle.FillColor = color.RGBA{59, 50, 75, 255}
					text.Color = color.White
				} else {
					subc.Circle.FillColor = color.White
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

	return content
}

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

func main() {
	gripper := NewGripper()

	myApp := app.New()
	myWindow := myApp.NewWindow("ALARIS Gripper Control")
	myWindow.Resize(fyne.NewSize(900, 300))

	fingers := []fingerPos{{0, 0, false, 50, 0}, {1, 4, true, 50, 0}}
	newcont := container.NewWithoutLayout(generateCircle(fingers))
	newcont.Resize(fyne.NewSize(300, 300))

	var fingerWidged []fyne.CanvasObject
	for _, v := range fingers {
		if v.active {
			fingerWidged = append(fingerWidged, fingerBar(v))
		}
	}
	fingerBarContainer := container.New(layout.NewVBoxLayout(), fingerWidged...)

	absposBar := container.NewWithoutLayout(fingerBarContainer)
	fingerBarContainer.Move(fyne.NewPos(300, 0))
	fingerBarContainer.Resize(fyne.NewSize(400, 300))

	nnewcont := container.New(layout.NewHBoxLayout(), fingerList(), newcont, absposBar)

	sendButton := widget.NewButton("send", send)
	stopButton := widget.NewButton("stop", stop)
	resetButton := widget.NewButton("reset", reset)
	buttons := container.New(&maxVbox{}, resetButton, stopButton, sendButton)

	withlobar := container.New(layout.NewVBoxLayout(), nnewcont, buttons)

	go serveGripper(gripper)

	myWindow.SetContent(withlobar)
	myWindow.ShowAndRun()

	fatih.Cyan("it reaches")
}

func send() {}

func reset() {}

func stop() {}

func subCircleTab() {
	fmt.Printf("subcucle\n")
}
