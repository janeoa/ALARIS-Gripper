#include <Wire.h>
#include "Motor.h"
#include "Roller.h"
// Include the required Wire library for I2C<br>#include <Wire.h>
/**
 * IC2 - M1 -  2/4 - 3  -  mid    A2^
 * IC2 - M2 -  6/7 - 5  -  top    A1
 * IC3 - M3 -  8/9 -10  -  roll   A0
 */

/**
 * ADC0 - Light Res
 * ADC1 - R1 (down)           top
 * ADC2 - R3 (closer to i2c)  mid
 */
Motor   mid(2, 4,  3, A2, 500, 780);
Motor   top(6, 7,  5, A1, 250, 410);
Roller roll(8, 9, 10, A3);

void setup() {
  Serial.begin(9600);
  Wire.begin(0); 
  Wire.onReceive(receiveEvent);
  Wire.onRequest(requestEvent);
  
  mid.setGoal(50);
  top.setGoal(50);
  roll.setGoal(4);
}


void receiveEvent(int bytes) {
  char buff[3];
  for(int i=0; i<bytes; i++){
    buff[i] = Wire.read();    // read one character from the I2C
  }

  roll.setGoal((byte)buff[0]);
  mid.setGoal((byte)buff[1]);
  top.setGoal((byte)buff[2]);
}

void requestEvent() {
  unsigned char reply[10];
  const unsigned int rawP = mid.getPos();
  unsigned int raw0 = mid.getPos();
  unsigned int raw1 = top.getPos();
  unsigned int raw2 = roll.getGoal();
  unsigned int raw3 = roll.getAnal();
  memcpy(reply+2, (char*)&rawP, 2);
  memcpy(reply+2, (char*)&raw0, 2);
  memcpy(reply+4, (char*)&raw1, 2);
  memcpy(reply+6, (char*)&raw2, 2);
  memcpy(reply+8, (char*)&raw3, 2);
  Wire.write(reply,10);
}

void loop() {
  mid.tick();
  top.tick();
  roll.tick();
//  requestEvent();
//  roll.printVars();
}


void dexDump(char *in, int len){
  Serial.print(">>");
  for(int i=0; i<len; i++){
    printHex(in[i]);
  }
  Serial.println("<<");
}

void printHex(uint8_t num) {
  char hexCar[2];

  sprintf(hexCar, "%02X", num);
  Serial.print(hexCar);
}
