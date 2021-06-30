package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var gripper *Gripper

type fingerPos struct {
	index  int
	pos    int
	active bool
	A      float64
	B      float64
}

func main() {
	gripper = NewGripper()

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

	statusText := binding.NewString()
	listofdevices := testPorts()
	statusText.Set(fmt.Sprintf("Choose UART device (%d found)", len(listofdevices)))
	connection_status := widget.NewLabelWithData(statusText)
	combo := widget.NewSelect(listofdevices, func(value string) {
		gripper.options.PortName = "/dev/" + value
		connection_status.Text = "Connecting..."
	})
	go fetchUART(gripper, statusText, combo)
	go sendUART(gripper)
	topbuttons := container.New(&maxVbox{}, connection_status, combo)

	withlobar := container.New(layout.NewVBoxLayout(), topbuttons, nnewcont, buttons)

	go serveGripper(gripper)

	myWindow.SetContent(withlobar)
	myWindow.ShowAndRun()
}

func send() {}

func reset() {}

func stop() {}
