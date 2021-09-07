#include <Wire.h>
#include <EasyTransfer.h>
#include <EasyTransferI2C.h>
#include "Lights.h"

EasyTransfer ETpcIN, ETpcOUT; 

#define INBUFFSIZE 5


struct PC_TO_MAIN{
  byte id;
  byte roll;
  byte rot;
  byte close;
};

PC_TO_MAIN pcdata;
PC_TO_MAIN mbdata;

const int pins[] = {2,3,4,5,6,7,8,9};
const char*emtpybuff5 = "\0\0\0\0\0";

Lights lights(pins);
char readbuffer[INBUFFSIZE];
byte foundSlaves[8] = {255,255,255,255,255,255,255,255};

void setup() {
  Serial.begin(115200);  // start serial for output
  Serial.println("init");
  ETpcIN.begin(details(pcdata), &Serial);
  ETpcOUT.begin(details(mbdata), &Serial);
//  Wire.setClock(10000);
  Wire.begin();        // join i2c bus (address optional for master)
  for (int i=0; i<8; i++){
    pinMode(pins[i], OUTPUT);
  }
//
////  Serial.println("searching for slaves");
  for (byte i = 0; i < 8; i++) {
    Wire.beginTransmission(i);
    if (Wire.endTransmission() == 0) {
      foundSlaves[i] = 4;
    }
  }
}

long lastCall = millis();

void loop() {
  if(ETpcIN.receiveData()){
    foundSlaves[pcdata.id] = pcdata.roll;
  }
  for(byte i=0; i<8; i++){
    if(foundSlaves[i]<255){
      mbdata.id = i;
      mbdata.roll = foundSlaves[i];   
      mbdata.rot = 255; 
      ETpcOUT.sendData();
    }
  }
//  delay(250);
}
//
//void checkFingerState(byte finger_id){
//  
//  const byte bufflen = 10;
//  byte readbuf[bufflen];
//  Wire.requestFrom(finger_id, bufflen);
//  bool got_data = false;
//  
//  byte binex = 0;
//  while(Wire.available()){
//    got_data = true;
//    readbuf[binex++] = Wire.read(); // receive a byte as character
//  }
//
//  if(got_data){
//    int aa,bb,cc,dd;
//    memcpy((int*)&aa  , readbuf+2, 2);
//    memcpy((int*)&bb  , readbuf+4, 2);
//    memcpy((int*)&cc  , readbuf+6, 2);
//    memcpy((int*)&dd  , readbuf+8, 2);
//    char msg[20];
//  }
//}
//
//void sendFingerState(int fingerID){
//  
//}

//void dexDump(char *in, int len){
//  Serial.print(">>");
//  for(int i=0; i<len; i++){
//    printHex(in[i]);
//  }
//  Serial.println("<<");
//}
//
//void printHex(uint8_t num) {
//  char hexCar[2];
//
//  sprintf(hexCar, "%02X", num);
//  Serial.print(hexCar);
//}
//
//unsigned char ToByte(bool b[8])
//{
//    unsigned char c = 0;
//    for (int i=0; i < 8; ++i)
//        if (b[i])
//            c |= 1 << i;
//    return c;
//}
