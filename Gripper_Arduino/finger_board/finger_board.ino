#include <Wire.h>
#include "Motor.h"
#include "Roller.h"
// Include the required Wire library for I2C<br>#include <Wire.h>
/**
 * IC2 - M1 -  2/4 - 3  -  mid
 * IC2 - M2 -  6/7 - 5
 * IC3 - M3 -  8/9 -10
 */

/**
 * ADC0 - Light Res
 * ADC1 - R1 (down)
 * ADC2 - R3 (closer to i2c)
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

//  char msg[10];
//  sprintf(msg, "%d\t%d\t%d", (byte)buff[0],(byte)buff[1],(byte)buff[2]);
  roll.setGoal((byte)buff[0]);
  mid.setGoal((byte)buff[1]);
  top.setGoal((byte)buff[2]);
//  Serial.println(msg);
}

void requestEvent() {
  byte reply[4];
  reply[0] = mid.getPos();
  reply[2] = top.getPos();
//  sprintf(msg, "%4d\t%4d\r\n", mid.getPos(), top.getPos());
  Wire.write(reply,4);
//  Wire.endTransmission();
  Serial.println("sent");
}

void loop() {
  mid.tick();
  top.tick();
  roll.tick();
}
