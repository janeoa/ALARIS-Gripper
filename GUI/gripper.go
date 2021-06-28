package main

import "github.com/jacobsa/go-serial/serial"

type Gripper struct {
	options serial.OpenOptions
}

func NewGripper() *Gripper {
	options := serial.OpenOptions{
		PortName:        "/dev/tty.usbserial-A8008HlV",
		BaudRate:        19200,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	gripper := &Gripper{
		options: options,
	}

	return gripper
}

func serveGripper(in *Gripper) {

}
