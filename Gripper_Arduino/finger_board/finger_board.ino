#include <Wire.h>
#include "Motor.h"

// Include the required Wire library for I2C<br>#include <Wire.h>
int LED = 13;
int x = 0;

Motor mid(7,8,9, A0, 310, 810);
Motor top(7,8,9, A1,  80, 410);

void setup() {
  // Define the LED pin as Output
  pinMode (LED, OUTPUT);
  // Start the I2C Bus as Slave on address 9
  Wire.begin(9); 
  // Attach a function to trigger when something is received.
  Wire.onReceive(receiveEvent);
  Serial.begin(9600);

  mid.setGoal(50);
}


void receiveEvent(int bytes) {
  x = Wire.read();    // read one character from the I2C
}


void loop() {
  mid.tick();
  
  Serial.print(mid.getPos());
  Serial.print("\t");
  Serial.println(mid.getGoal());
// 
  
}
