#include <Wire.h>
#include "Motor.h"

// Include the required Wire library for I2C<br>#include <Wire.h>

Motor mid(7,8,9, A0, 310, 810);
Motor top(7,8,9, A1,  80, 410);


void setup() {
  Serial.begin(9600);
  Wire.begin(9); 
  Wire.onReceive(receiveEvent);
  Wire.onRequest(requestEvent);
  
  mid.setGoal(50);
}


void receiveEvent(int bytes) {
//  x = Wire.read();    // read one character from the I2C
}

void requestEvent() {
//  char msg[11];
//  sprintf(msg, "%4d\t%4d\r\n", mid.getPos(), top.getPos());
  Wire.write("help");
//  Wire.endTransmission();
}

void loop() {
//  mid.tick();
}
