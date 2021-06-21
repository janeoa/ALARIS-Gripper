import serial


while True:
    try:
        ser = serial.Serial('/dev/cu.usbmodem14101')
        line = ser.readline()
        print(line)
        break
    except (FileNotFoundError, serial.serialutil.SerialException):
        print("No device connected")
    


# with serial.Serial('/dev/cu.usbmodem14101', 19200, timeout=1) as ser:
#     line = ser.readline()   # read a '\n' terminated line
#     print(line)

ser.close()             # close port