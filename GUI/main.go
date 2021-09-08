package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/fatih/color"
)

var newPosEntries []*widget.Entry
var gripper *Gripper
var fingerBarList *fyne.Container
var circle *fyne.Container
var myWindow fyne.Window
var combo *widget.Select
var statusText binding.String

var fingersToRoute []bool

var current_mode int

const AUTO = 0
const MANUAL = 1

type fingerPos struct {
	index  int
	pos    int
	newPos int
	active bool
	// a      int
	// b      int
	A binding.Float
	B binding.Float
}

// var commandStack []command

func main() {
	current_mode = MANUAL
	gripper = NewGripper()

	myApp := app.New()
	myWindow = myApp.NewWindow("ALARIS Gripper Control")

	fingersToRoute = []bool{false, false, false, false, false, false, false, false}
	// fingers := []fingerPos{{0, 0, 0, false, 50, 0}, {1, 4, 4, false, 50, 0}, {2, 3, 3, false, 50, 0}}
	fingers := []fingerPos{}
	gripper.finger = fingers

	go sendUART()
	go serveGripper(gripper)
	go fetchUART()

	myWindow.SetContent(generateGUI())
	myWindow.ShowAndRun()
}

func generateGUI() *fyne.Container {
	color.Yellow("regenerating the GUI\n")

	circle = generateCircle()
	fingerBarList = generateFingerBarList()

	// centercontwithoutline := container.New(layout.NewHBoxLayout(), fingerList())

	var fingerMoveList fyne.CanvasObject
	if current_mode == AUTO {
		fingerMoveList = autoContainer()
	} else {
		fingerMoveList = fingerList()
	}

	buttons := bottom()

	topbuttons := topBar()

	// return container.New(layout.NewBorderLayout(topbuttons, buttons, fingerMoveList, circle), topbuttons, buttons, fingerMoveList, circle, fingerBarList)
	onTheLeft := container.New(layout.NewHBoxLayout(), fingerMoveList, circle)
	return container.New(layout.NewBorderLayout(topbuttons, buttons, onTheLeft, nil), topbuttons, buttons, onTheLeft, fingerBarList)
}

func send() {
	for _, v := range gripper.finger {
		A, _ := v.A.Get()
		B, _ := v.B.Get()
		// if v.pos != v.newPos || v.a != int(A) || v.b != int(B) {
		// gripper.tosend = fmt.Sprintf("%d", v.newPos)

		toSend := command{
			byte(v.index),
			byte(v.newPos),
			byte(int(A)),
			byte(int(B)),
		}
		EasyTransferEncode(toSend)
		v.pos = v.newPos
		myWindow.SetContent(generateGUI())
		break
		// }
	}

}

func reset() {
	for i, v := range gripper.finger {
		v.newPos = v.pos
		newPosEntries[i].SetText(fmt.Sprintf("%d", v.pos))
		myWindow.SetContent(generateGUI())
	}
}

func stop() {}
