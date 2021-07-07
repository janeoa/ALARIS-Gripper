package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func topBar() *fyne.Container {

	statusText := binding.NewString()
	listofdevices := testPorts()
	statusText.Set(fmt.Sprintf("Choose UART device (%d found)", len(listofdevices)))
	connection_status := widget.NewLabelWithData(statusText)
	combo := widget.NewSelect(listofdevices, func(value string) {
		gripper.options.PortName = "/dev/" + value
		connection_status.Text = "Connecting..."
	})
	go fetchUART(statusText, combo)
	return container.New(&maxVbox{}, connection_status, combo)
}

func bottom() *fyne.Container {
	sendButton := widget.NewButton("send", send)
	stopButton := widget.NewButton("stop", stop)
	resetButton := widget.NewButton("reset", reset)
	return container.New(&maxVbox{}, resetButton, stopButton, sendButton)
}

func fetchUART(status binding.String, combo *widget.Select) {
	for {
		time.Sleep(time.Millisecond * 400)
		if gripper.options.PortName == "placeholder" {
			listofdevices := testPorts()
			status.Set(fmt.Sprintf("Choose UART device (%d found)", len(listofdevices)))

			combo.Options = listofdevices
			// log.Printf("%v", listofdevices)
		} else {
			if gripper.connected {
				status.Set("Connected")
			} else {
				status.Set("Disconnected")
			}
			// continue
		}
	}
}
