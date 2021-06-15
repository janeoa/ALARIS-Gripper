#include <Wire.h>
#include "Lights.h"

#define INBUFFSIZE 5

const int pins[] = {2,3,4,5,6,7,8,9};
Lights lights(pins);
char readbuffer[INBUFFSIZE];

void setup() {
  
  
  Serial.begin(9600);  // start serial for output

  Serial.print("init");
  
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
    byte a = -1,b = -1;
    char msg[50];
    bool checked = false;
    
    if(strstr(readbuffer, " ")>0){
      sscanf (readbuffer,"%d %d", &a,&b);
      if(a < 0 || a > 7 || b < 0 || b > 7){
        Serial.println("Enter okay number, please");  
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
        sprintf(msg, "Got: %d, its pin %d",a, pins[a]);
        lights.setLight(a);
        checked = true;
      }
    }
     
    Serial.println(msg);
    if(checked){
      byte toFinger[3];
      toFinger[0] =  a; // new roll  pos
      toFinger[1] = 50; // new rot   pos
      toFinger[2] = 65; // new close pos
      
      dexDump(toFinger,3);
      Wire.beginTransmission(9);
      Wire.write(toFinger, 3);
      Wire.endTransmission();
    }
  }

  checkFingerState(0);

}

void checkFingerState(byte finger_id){
  
  const byte bufflen = 4;
  byte readbuf[bufflen];
  
  Wire.requestFrom(finger_id, bufflen);

  
  byte binex = 0;
  while(Wire.available()){
    readbuf[binex++] = Wire.read(); // receive a byte as character
  }

  int a = (int)readbuf[0:1];
  int b = (int)readbuf[2];
  char msg[10];
  sprintf(msg,"%d\t%d",a,b);
  Serial.println(msg);

  dexDump(readbuf, bufflen);

  delay(300);
}

void sendFingerState(int fingerID){
  
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
