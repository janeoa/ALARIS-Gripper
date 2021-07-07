package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var gripper *Gripper
var fingerBarList *fyne.Container
var circle *fyne.Container
var myWindow fyne.Window

type fingerPos struct {
	index  int
	pos    int
	newPos int
	active bool
	A      int
	B      int
}

func main() {
	gripper = NewGripper()

	myApp := app.New()
	myWindow = myApp.NewWindow("ALARIS Gripper Control")

	fingers := []fingerPos{{0, 0, 0, false, 50, 0}, {1, 4, 6, false, 50, 0}}
	gripper.finger = fingers

	go sendUART()
	go serveGripper(gripper)

	myWindow.SetContent(generateGUI())
	myWindow.ShowAndRun()
}

func generateGUI() *fyne.Container {
	circle = generateCircle()
	fingerBarList = generateFingerBarList()

	// centercontwithoutline := container.New(layout.NewHBoxLayout(), fingerList())
	fingerMoveList := fingerList()

	buttons := bottom()

	topbuttons := topBar()

	// return container.New(layout.NewBorderLayout(topbuttons, buttons, fingerMoveList, circle), topbuttons, buttons, fingerMoveList, circle, fingerBarList)
	onTheLeft := container.New(layout.NewHBoxLayout(), fingerMoveList, circle)
	return container.New(layout.NewBorderLayout(topbuttons, buttons, onTheLeft, nil), topbuttons, buttons, onTheLeft, fingerBarList)
}

func send() {}

func reset() {}

func stop() {}
