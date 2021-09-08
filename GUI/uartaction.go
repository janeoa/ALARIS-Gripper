package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func topBar() *fyne.Container {

	statusText = binding.NewString()
	statusText.Set("Connected")
	connection_status := widget.NewLabelWithData(statusText)
	if !gripper.connected {
		listofdevices := testPorts()

		combo = widget.NewSelect(listofdevices, func(value string) {
			gripper.options.PortName = "/dev/" + value
			connection_status.Text = "Connecting..."
		})
	}

	refreshButton := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		gripper.finger = []fingerPos{}
		myWindow.SetContent(generateGUI())
	})
	return container.New(&maxVbox{}, container.New(layout.NewHBoxLayout(), connection_status, refreshButton), combo)
}

func bottom() *fyne.Container {
	sendButton := widget.NewButton("send", send)
	stopButton := widget.NewButton("stop", stop)
	resetButton := widget.NewButton("reset", reset)
	return container.New(&maxVbox{}, resetButton, stopButton, sendButton)
}

func fetchUART() {
	for {
		time.Sleep(time.Millisecond * 400)
		if gripper.options.PortName == "placeholder" {
			listofdevices := testPorts()
			statusText.Set(fmt.Sprintf("Choose UART device (%d found)", len(listofdevices)))

			combo.Options = listofdevices
			// log.Printf("%v", listofdevices)
		} else {
			if gripper.connected {
				stat, _ := statusText.Get()
				if stat == "Connected" {

				} else {
					gripper.finger = []fingerPos{}
					myWindow.SetContent(generateGUI())
					statusText.Set("Connected")
				}
			} else {
				statusText.Set("Disconnected")
			}
			// continue
		}
	}
}
