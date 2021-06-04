#include <Wire.h>
#include "Lights.h"

#define INBUFFSIZE 5

const int pins[] = {2,3,4,5,6,7,8,9};
Lights lights(pins);
char readbuffer[INBUFFSIZE];

void setup() {
  
  Serial.begin(9600);  // start serial for output
  
  Wire.setClock(10000);
  Wire.begin();        // join i2c bus (address optional for master)
  for (int i=0; i<8; i++){
    pinMode(pins[i], OUTPUT);
  }
//  Serial.println("mid.pos top.pos");
  Serial.println(F("ready to get index [0-7]"));
}

void loop() {
  lights.tick();

  
  if (Serial.available()) {
    Serial.readBytes(readbuffer, INBUFFSIZE);
    int a = -1,b = -1;
    char msg[50];
    bool checked = false;
    
    if(strstr(readbuffer, " ")>0){
      sscanf (readbuffer,"%d %d", &a,&b);
      if(a < 0 || a > 7 || b < 0 || b > 7){
        Serial.println("You motherfucker");  
      }else{
        sprintf(msg, "Got: %d and %d, its pin %d and %d",a,b, pins[a],  pins[b]);
        lights.setLight(a, b);
        checked = true;
      }
    }else{
      sscanf (readbuffer,"%d", &a);
      if(a < 0 || a > 7){
        Serial.println("Enter okay number, please");  
      }else{
        sprintf(msg, "Got: %d and %d, its pin %d and %d",a, pins[a]);
        lights.setLight(a);
        checked = true;
      }
    }
     
    Serial.println(msg);
    delay(20);
    if(checked){
      Wire.beginTransmission(9);
      Wire.write(a);
      Wire.endTransmission();
    }
  }

}
