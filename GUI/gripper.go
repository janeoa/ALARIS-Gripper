package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
	"time"

	"fyne.io/fyne/v2/data/binding"
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

const printUARTlogs = false

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

		// buf := []byte{}
		rx_len := 0
		faults := 0
		buf := make([]byte, 2000)
		for {
			subbuff := make([]byte, 2000)
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
				subbuff = subbuff[:n]
				buf = append(buf, subbuff...)
				if printUARTlogs {
					fmt.Print(hex.Dump(buf[:40]))
					fmt.Println("Rx: ", strings.TrimRight(hex.EncodeToString(buf), "0"))
					fmt.Print("\n\n\n")
				}
				if rx_len == 0 && len(buf) >= 8 {
					for i, v := range buf {
						if v == 0x85 {
							faults = 0
							rx_len = int(buf[i+1])
							if printUARTlogs {
								color.Yellow("Data len is %d", rx_len)
							}
							if rx_len == 4 {
								if printUARTlogs {
									color.Green("Data len valid")
								}
								parseArduinoCommand(buf[i+2 : i+6])
								buf = buf[i+6:]
								rx_len = 0
								break
							} else {
								color.Red("Data is not 4 bytes long (%d)", rx_len)
								rx_len = 0
								break
							}
						}
					}
					faults++
					if faults > 10 {
						// for i := 0; i < 50; i++ {
						color.Red("UART buffur Fault accured\n")
						gripper.finger = []fingerPos{}
						// }
						break
					}
				}
			}
		}
	}
}

type command struct {
	id  byte
	pos byte
	rol byte
	ang byte
}

func EasyTransferEncode(in command) {
	size := reflect.TypeOf(in).Size()
	CS := byte(size)
	toOut := []byte{0x06, 0x85}
	toOut = append(toOut, byte(size))

	toOut = append(toOut, in.id)
	CS ^= in.id
	toOut = append(toOut, in.pos)
	CS ^= in.pos
	toOut = append(toOut, in.rol)
	CS ^= in.rol
	toOut = append(toOut, in.ang)
	CS ^= in.ang

	toOut = append(toOut, CS)
	// if printUARTlogs {
	color.Cyan("Writing %v, as %v bytes using EasyTransfer\n", in, toOut)
	// }
	gripper.port.Write(toOut)
}

func parseArduinoCommand(in []byte) {
	if printUARTlogs {
		color.Cyan("Finger ID: %d\nroll: \t%d\nrott: \t%d\n", in[0], in[1], in[2])
	}
	updated := false
	for i, v := range gripper.finger {
		if v.index == int(in[0]) {
			if v.pos != int(in[1]) {
				color.Cyan("ID: %d|\t\t%d->%d\n", v.index, v.pos, int(in[1]))
				gripper.finger[i].pos = int(in[1])
				myWindow.SetContent(generateGUI())
			}
			updated = true
		}
	}
	if !updated {
		gripper.finger = append(gripper.finger, fingerPos{
			index:  int(in[0]),
			pos:    int(in[1]),
			newPos: int(in[1]),
			active: false,
			A:      binding.NewFloat(),
			B:      binding.NewFloat(),
		})
		if int(in[2]) < 255 {
			gripper.finger[len(gripper.finger)-1].A.Set(float64((in[2])))
			gripper.finger[len(gripper.finger)-1].B.Set(float64((in[3])))
		} else {
			gripper.finger[len(gripper.finger)-1].A.Set(50)
			gripper.finger[len(gripper.finger)-1].B.Set(0)
		}
		color.Cyan("new finger ID: %d, pos: %d\n", int(in[0]), int(in[1]))
		myWindow.SetContent(generateGUI())
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
		if strings.Contains(f.Name(), "tty") && len(f.Name()) > 8 && f.Name() != "tty.Bluetooth-Incoming-Port" && f.Name() != "tty.jane-CSRGAIA" && f.Name() != "tty.GalaxyBudsLive4589-GEAR" {
			options := serial.OpenOptions{
				PortName:        "/dev/" + f.Name(),
				BaudRate:        115200,
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
