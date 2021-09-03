#include <Wire.h>
#include "Lights.h"

#define INBUFFSIZE 5

const int pins[] = {2,3,4,5,6,7,8,9};
const char*emtpybuff5 = "\0\0\0\0\0";

Lights lights(pins);
char readbuffer[INBUFFSIZE];
byte foundSlaves[8] = {0, 0, 0, 0, 0, 0, 0, 0};

void setup() {
  Serial.begin(115200);  // start serial for output
  Wire.setClock(10000);
  Wire.begin();        // join i2c bus (address optional for master)
  for (int i=0; i<8; i++){
    pinMode(pins[i], OUTPUT);
  }
  Serial.println('\0');

//  Serial.println("searching for slaves");
  Wire.begin();
  for (byte i = 0; i < 8; i++) {
    Wire.beginTransmission(i);
    if (Wire.endTransmission() == 0) {
      foundSlaves[i] = 4;
    }
  }
  replyslC(3,50);
}

void replyslC(byte n, byte dl){
  char msg[50];
  sprintf(msg, "foundSlaves: %d%d%d%d%d%d%d%d\r\n", foundSlaves[0],foundSlaves[1],foundSlaves[2],foundSlaves[3],foundSlaves[4],foundSlaves[5],foundSlaves[6],foundSlaves[7]);
  if(dl>1){
    for (int i=0; i<n;i++){
      Serial.println(msg);
      delay(dl);
    }
  }else{
    Serial.println(msg);
  }
}

void replyslC(){
    replyslC(1,0);
}


long lastCall = millis();

void loop() {
  
  
  lights.tick();
  
  if (Serial.available()) {
    Serial.readBytes(readbuffer, INBUFFSIZE);
    int a = -1,b = -1;
    char msg[50];
    bool checked = false;
//    sprintf(msg, "Readbufis>>%s<<\n",readbuffer);
//    Serial.println(msg);
    
    if(strstr(readbuffer, "slC")>0){
      replyslC();
    }else if(strstr(readbuffer, " ")>0){
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
        foundSlaves[0] = a;
        checked = true;
      }
    }
    Serial.println(msg);
    msg[0] = '\0';
    memcpy(readbuffer, emtpybuff5, INBUFFSIZE);
    
    if(checked){
      byte toFinger[3];
      toFinger[0] =  a; // new roll  pos
      toFinger[1] = 50; // new rot   pos
      toFinger[2] = 65; // new close pos
      
//      dexDump(toFinger,3);
      Wire.beginTransmission(0);
      Wire.write(toFinger, 3);
      Wire.endTransmission(true);
    }
    replyslC();
  }

  if(millis()-lastCall>250){
    checkFingerState(0);
    lastCall = millis();
  }

}

void checkFingerState(byte finger_id){
  
  const byte bufflen = 10;
  byte readbuf[bufflen];
  Wire.requestFrom(finger_id, bufflen);
  bool got_data = false;
  
  byte binex = 0;
  while(Wire.available()){
    got_data = true;
    readbuf[binex++] = Wire.read(); // receive a byte as character
  }

  if(got_data){
    int aa,bb,cc,dd;
    memcpy((int*)&aa  , readbuf+2, 2);
    memcpy((int*)&bb  , readbuf+4, 2);
    memcpy((int*)&cc  , readbuf+6, 2);
    memcpy((int*)&dd  , readbuf+8, 2);
    char msg[20];
//    sprintf(msg,"%d\t%d\t%03d\t%d",aa,bb,cc,dd);
//    Serial.println(msg);
//    dexDump(readbuf, bufflen);
  }
  
//  delay(300);
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

unsigned char ToByte(bool b[8])
{
    unsigned char c = 0;
    for (int i=0; i < 8; ++i)
        if (b[i])
            c |= 1 << i;
    return c;
}
