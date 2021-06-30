package main

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func fetchUART(in *Gripper, status binding.String, combo *widget.Select) {
	for {
		time.Sleep(time.Millisecond * 400)
		if in.options.PortName == "placeholder" {
			listofdevices := testPorts()
			status.Set(fmt.Sprintf("Choose UART device (%d found)", len(listofdevices)))

			combo.Options = listofdevices
			log.Printf("%v", listofdevices)
		} else {
			if in.connected {
				status.Set("Connected")
			} else {
				status.Set("Disconnected")
			}
			// continue
		}
	}
}
