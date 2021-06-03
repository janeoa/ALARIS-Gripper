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
    int a,b;
    char msg[50];
    
    if(strstr(readbuffer, " ")>0){
      sscanf (readbuffer,"%d %d", &a,&b);
      sprintf(msg, "Got: %d and %d, its pin %d and %d",a,b, pins[a],  pins[b]);
      lights.setLight(a, b);
    }else{
      sscanf (readbuffer,"%d", &a);
      sprintf(msg, "Got: %d, its pin %d",a, pins[a]);
      lights.setLight(a);
    }
     
    Serial.println(msg);
    
    Wire.beginTransmission(9);
    Wire.write(a);
    Wire.endTransmission();
  }

}
