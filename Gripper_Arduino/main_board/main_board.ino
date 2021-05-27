#include <Wire.h>

void setup() {
  Wire.setClock(10000);
  Wire.begin();        // join i2c bus (address optional for master)
  Serial.begin(9600);  // start serial for output
  Serial.println("mid.pos top.pos");
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
