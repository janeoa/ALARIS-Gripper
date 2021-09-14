#include <Wire.h>
#include <EasyTransfer.h>
#include <EasyTransferI2C.h>
#include "Lights.h"

EasyTransfer ETfromPCtoMain, ETfromMainToPC; 
EasyTransferI2C ETfromMainToFinger; 
EasyTransferI2C ETfromFingerToMain;


struct PC_TO_MAIN{
  byte id;
  byte pos;
  byte A;
  byte B;
};

struct MAIN_TO_PC{
  byte id;
  byte pos;
  byte A;
  byte B;
};

struct SEND_DATA_STRUCTURE{
  byte dir;
  byte A;
  byte B;
};

struct RECIEVE_FINGER_STATE{
  byte id;
  byte A;
  byte B;
};

PC_TO_MAIN dataFromPC;
MAIN_TO_PC dataToPC;

SEND_DATA_STRUCTURE dataToFinger;
RECIEVE_FINGER_STATE dataFromFinger;
bool is_next_on_right(int prev, int curr);

const byte pins[] = {2,3,4,5,6,7,8,9};

Lights lights(pins);
byte foundSlaves[8] = {255,255,255,255,255,255,255,255};
byte prev[8] = {255,255,255,255,255,255,255,255};

void setup() {
  Serial.begin(115200);  // start serial for output
  Serial.println(F("init"));
  ETfromPCtoMain.begin(details(dataFromPC), &Serial);
  ETfromMainToPC.begin(details(dataToPC), &Serial);
//  Wire.setClock(10000);
  Wire.begin(69);        // join i2c bus (address optional for master)
  Wire.onReceive(receive);
  ETfromMainToFinger.begin(details(dataToFinger), &Wire);
  ETfromFingerToMain.begin(details(dataFromFinger), &Wire);
  for (byte i=0; i<8; i++){
    pinMode(pins[i], OUTPUT);
  }
  delay(100);
  for (byte i=0; i<8; i++){
    digitalWrite(pins[i], HIGH);
    delay(100);
    digitalWrite(pins[i], LOW);
  }
//
////  Serial.println("searching for slaves");
  for (byte i = 0; i < 8; i++) {
    Wire.beginTransmission(i);
    if (Wire.endTransmission() == 0) {
      foundSlaves[i] = i;
      prev[i] = i;
    }
  }
}

void receive(int numBytes) {}

byte slaveCallIndex = 0;
unsigned long last = millis();

void loop() {
//  if(millis()-last>200){
//    last = millis();
//    dataToFinger.dir = 69;
//    if(foundSlaves[slaveCallIndex] != 255){
//      ETfromMainToFinger.sendData(slaveCallIndex);
//    }
//    slaveCallIndex = (slaveCallIndex==7)?0:slaveCallIndex+1;
//  }
  if(ETfromPCtoMain.receiveData()){
    digitalWrite(pins[prev[dataFromPC.id]], LOW);
    digitalWrite(pins[dataFromPC.pos], HIGH);
    delay(50);
    prev[dataFromPC.id] = dataFromPC.pos;
    if(foundSlaves[dataFromPC.id] != dataFromPC.pos){
      dataToFinger.dir = is_next_on_right(foundSlaves[dataFromPC.id], dataFromPC.pos)?0:1;//dataFromPC.pos;
    }else{
//      digitalWrite(pins[dataFromPC.pos], LOW);
      dataToFinger.dir = 255;
    }
    
    dataToFinger.A = dataFromPC.A;
    dataToFinger.B = dataFromPC.B;
    ETfromMainToFinger.sendData(dataFromPC.id);

    foundSlaves[dataFromPC.id] = dataFromPC.pos;
  }
  
  if(ETfromFingerToMain.receiveData()){
    dataToPC.id   = dataFromFinger.id;
    dataToPC.pos  = foundSlaves[dataFromFinger.id];//dataFromFinger.pos;   
    dataToPC.A    = dataFromFinger.A;
    dataToPC.B    = dataFromFinger.B;
    ETfromMainToPC.sendData();
  }else{
    for(byte i=0; i<8; i++){
      if(foundSlaves[i]<255){
        dataToPC.id = i;
        dataToPC.pos = foundSlaves[i];   
        dataToPC.A = 255;
        ETfromMainToPC.sendData();
      }
    }
  }
//  delay(250);
}

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

bool is_next_on_right(int prev, int curr){
  bool is_right = false;
  
  int delta = curr - prev;
  if(delta > 0 && delta < 4) is_right = true;
  if(delta < -4 && delta > -8) is_right = true;
  return is_right;
}
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
