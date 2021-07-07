package main

import (
	"encoding/hex"
	"fmt"
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
	port      io.ReadWriteCloser
	finger    []fingerPos
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
	time.Sleep(time.Millisecond * 30)
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

		in.port = port
		// Make sure to close it later.
		defer port.Close()

		buf := []byte{}
		for {
			subbuff := make([]byte, 100)
			n, err := port.Read(subbuff)
			if err != nil {
				if err != io.EOF {
					log.Println("Error reading from serial port: ", err)
					in.connected = false
					port.Close()
					break
				}
			} else {
				in.connected = true
				subbuff = subbuff[:n]
				buf = append(buf, subbuff...)
				color.Cyan("%s", hex.Dump(buf))
				if strings.Contains(string(buf), "\r\n") {
					subs := strings.Split(string(buf), "\r\n")
					if len(subs) > 1 {
						for i := 0; i < len(subs); i++ {
							if len(subs[i]) > 1 {
								color.Red("len(subs[i]): %d", len(subs[i]))
								parseRX(subs[i])
							}
						}
						// buf = []byte(subs[len(subs)-1])
						buf = []byte{}
					}
				}
			}
		}
	}
}

func sendUART() {
	for {
		time.Sleep(time.Millisecond * 200)
		if gripper.tosend != "" {
			color.Cyan("Writing %s\n", gripper.tosend)
			gripper.port.Write([]byte(gripper.tosend + "\r\n"))
			gripper.tosend = ""
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
		if strings.Contains(f.Name(), "tty") && len(f.Name()) > 8 && f.Name() != "tty.Bluetooth-Incoming-Port" && f.Name() != "tty.jane-CSRGAIA" {
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

func parseRX(in string) {
	color.Yellow("parsing {%s}", in)
	if strings.Contains(in, "foundSlaves") {
		color.Green("found slaves {%s}", strings.Trim(in, "\r\n"))
		slaves := []int{0, 0, 0, 0, 0, 0, 0, 0}
		_, err := fmt.Scanf(strings.Trim(in, "\r\n"), "foundSlaves: %d%d%d%d%d%d%d%d", slaves[0])
		if err != nil {
			color.Red("Error: %v", err)
		}
	}
}
