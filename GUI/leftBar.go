package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func generateFingerWithFunc(i int) fyne.CanvasObject {
	buttonAR := widget.NewButton(fmt.Sprintf("%d", i), nil)
	buttonAR.OnTapped = func() { autoRouteButtonPressed(i, buttonAR) }
	return buttonAR
}

func autoContainer() fyne.CanvasObject {
	var autoRouteButtons []fyne.CanvasObject

	for i := 0; i < 8; i++ {
		label := widget.NewLabel(fmt.Sprintf("%d", i))
		autoRouteButtons = append(autoRouteButtons, label)
	}
	for i := 0; i < 8; i++ {
		autoRouteButtons = append(autoRouteButtons, generateFingerWithFunc(i))
	}

	autoRouteButtonContainer := container.New(layout.NewGridLayout(8), autoRouteButtons...)

	return autoRouteButtonContainer
}

func fingerList() fyne.CanvasObject {
	var fingers []fyne.CanvasObject
	for i := 0; i < len(gripper.finger); i++ {
		fingers = append(fingers, fingerInfoItem(i, gripper.finger[i].active))
	}

	fingerss := container.New(layout.NewVBoxLayout(), fingers...)

	return fingerss //container.New(layout.NewVBoxLayout(), autoRouteButtonContainer, fingerss)
}

func fingerInfoItem(id int, active bool) fyne.CanvasObject {
	var button *widget.Button

	// for _, v := range gripper.finger {
	var primaryIcon, secondaryIcon fyne.Resource
	if active {
		primaryIcon = theme.ConfirmIcon()
		secondaryIcon = theme.InfoIcon()
	} else {
		primaryIcon = theme.InfoIcon()
		secondaryIcon = theme.ConfirmIcon()
	}
	button = widget.NewButtonWithIcon("", primaryIcon, func() {
		gripper.finger[id].active = !gripper.finger[id].active
		if gripper.finger[id].active {
			button.Icon = primaryIcon
		} else {
			button.Icon = secondaryIcon
		}
		myWindow.SetContent(generateGUI())
	})
	// }

	// activateButton.Text = "tapped"
	name := widget.NewLabel(fmt.Sprintf("#%d is at: %d, new = ", id, gripper.finger[id].pos))

	newpos := widget.NewEntry()
	newPosEntries = append(newPosEntries, newpos)
	newpos.SetPlaceHolder(fmt.Sprintf("%d", gripper.finger[id].newPos))
	newpos.OnChanged = func(s string) {
		if strings.ContainsAny(s, "01234567") && len(s) == 1 {
			log.Printf("new pos for %d is %s", id, s)
			npi, err := strconv.ParseInt(s, 10, 64)
			// npi, err := fmt.Scanf("%d", s)
			if err != nil {
				fmt.Println("ERROR ON SCANF")
			}
			if npi >= 0 && npi < 8 {
				gripper.finger[id].newPos = int(npi)
				// log.Printf("new pos for %d is %d", id, npi)
				myWindow.SetContent(generateGUI())
				// circle.Refresh()
				// log.Printf("refreshed")

			}
		}
		if len(s) > 1 {
			newpos.SetText(fmt.Sprintf("%d", gripper.finger[id].pos))
			gripper.finger[id].newPos = gripper.finger[id].pos
		}
	}

	// dummyText := widget.NewLabel(" ")

	line1 := container.New(layout.NewHBoxLayout(), button, name, newpos)
	// line2 := container.New(layout.NewCenterLayout(), dummyText)

	return container.New(layout.NewVBoxLayout(), widget.NewSeparator(), line1)
}

func autoRouteButtonPressed(id int, button *widget.Button) {
	if fingersToRoute[id] {
		// button.SetIcon(nil)
		button.SetText(fmt.Sprintf("%d", id))
		button.Resize(button.MinSize())
	} else {
		button.SetText("*")
		// button.SetIcon(theme.ConfirmIcon())
	}
	fingersToRoute[id] = !fingersToRoute[id]
	log.Printf("%v", fingersToRoute)
}
