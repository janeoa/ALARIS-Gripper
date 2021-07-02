package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func fingerList() fyne.CanvasObject {
	var fingers []fyne.CanvasObject
	for i := 0; i < len(gripper.finger); i++ {
		fingers = append(fingers, fingerInfoItem(i, gripper.finger[i].active))
	}
	return container.New(layout.NewVBoxLayout(), fingers...)
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
	name := widget.NewLabel(fmt.Sprintf("#%d at:", id))
	pos := widget.NewLabel(fmt.Sprintf("%d", id))
	return container.New(layout.NewHBoxLayout(), button, name, pos)
}
