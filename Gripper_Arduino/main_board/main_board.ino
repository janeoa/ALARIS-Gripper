#include <Wire.h>
#include "Lights.h"

const int pins[] = {2,3,4,5,6,7,8,9};
//Lights lights(pins);

LightsLights lights(pins);

void setup() {
  int array[8] = {1, 2, 3, 4, 5, 6, 7, 8};
  
  Serial.begin(9600);  // start serial for output
  
  Serial.print(">>");
  for(int i=0; i<8; i++){
    Serial.print(lights.getPin(i));  
    Serial.print("\t");  
  }
  Serial.println("<<");
  
  Wire.setClock(10000);
  Wire.begin();        // join i2c bus (address optional for master)
  
//  Serial.println("mid.pos top.pos");

  while(true){}
}

void loop() {
  Serial.println("...");
  // check if the I2C lines are LOW
  if (digitalRead(SDA) == LOW || digitalRead(SCL) == LOW)
  {
    Serial.println("Bus error");
  }else{
    Serial.println(Wire.requestFrom(9, 10));    // request 6 bytes from slave device #8
    Serial.println("requested");
  
    Serial.print(">>");
    while (Wire.available()) { // slave may send less than requested
  //    Serial.println("available");
      char c = Wire.read(); // receive a byte as character
  //    Serial.print(">>");
      Serial.print(c);         // print the character
    }
    Serial.println("<<");
  }
  delay(500);
}
