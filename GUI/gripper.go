package main

import (
	"io"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/jacobsa/go-serial/serial"
)

type Gripper struct {
	options   serial.OpenOptions
	connected bool
	tosend    string
}

func NewGripper() *Gripper {
	portname := "placeholder"

	options := serial.OpenOptions{
		PortName:        portname,
		BaudRate:        115200,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	gripper := &Gripper{
		options:   options,
		connected: false,
		tosend:    "",
	}

	return gripper
}

func serveGripper(in *Gripper) {
	color.CyanString("serve gripper...")
	for {
		if in.options.PortName == "placeholder" {
			time.Sleep(time.Second)
			continue
		}
		// Open the port.
		port, err := serial.Open(in.options)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		// Make sure to close it later.
		defer port.Close()

		for {
			buf := make([]byte, 32)
			n, err := port.Read(buf)
			if err != nil {
				if err != io.EOF {
					log.Println("Error reading from serial port: ", err)
					in.connected = false
					port.Close()
					break
				}
			} else {
				in.connected = true
				buf = buf[:n]
				color.Yellow("rx: [%s]\n", strings.Trim(string(buf), "\r\n"))
			}

			if in.tosend != "" {
				port.Write([]byte(in.tosend))
				in.tosend = ""
			}
		}
	}
}

func testPorts() []string {
	listofports := []string{}

	files, err := ioutil.ReadDir("/dev/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.Contains(f.Name(), "tty") && len(f.Name()) > 8 {
			options := serial.OpenOptions{
				PortName:        "/dev/" + f.Name(),
				BaudRate:        19200,
				DataBits:        8,
				StopBits:        1,
				MinimumReadSize: 4,
			}
			_, err := serial.Open(options)
			if err == nil {
				listofports = append(listofports, f.Name())
			}
		}
	}
	return listofports
}
