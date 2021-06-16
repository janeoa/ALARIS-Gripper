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
Roller roll(8, 9, 10, A0);

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
  unsigned char reply[6];
  unsigned int raw0 = analogRead(A0);//mid.getPos();
  unsigned int raw1 = analogRead(A1);//top.getPos();
  unsigned int raw2 = analogRead(A2);//top.getPos();
  memcpy(reply  , (char*)&raw0, 2);
  memcpy(reply+2, (char*)&raw1, 2);
  memcpy(reply+4, (char*)&raw2, 2);
  Wire.write(reply,6);
}

void loop() {
  mid.tick();
  top.tick();
  roll.tick();
}
