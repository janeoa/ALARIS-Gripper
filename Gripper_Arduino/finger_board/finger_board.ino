#include <Wire.h>
#include <EasyTransferI2C.h>
#include "Motor.h"
#include "Roller.h"
// Include the required Wire library for I2C<br>#include <Wire.h>
/**
 * IC2 - M1 -  2/4 - 3  -  mid    A2^
 * IC2 - M2 -  6/7 - 5  -  top    A1
 * IC3 - M3 -  8/9 -10  -  roll   A0
 */

/**
 * ADC0 - Light Res
 * ADC1 - R1 (down)           top
 * ADC2 - R3 (closer to i2c)  mid
 */
Motor   mid(2, 4,  3, A2, 500, 780);
Motor   top(6, 7,  5, A1, 250, 410);
Roller roll(8, 9, 10, A3);

#define THIS_FINGER_ID 0

struct RECEIVE_DATA_STRUCTURE{
  byte pos;
  byte A;
  byte B;
};

struct SEND_STATE{
  byte id;
  byte pos;
  byte npos;
  byte state;
  byte A;
  byte B;
};


EasyTransferI2C ETmain; 
EasyTransferI2C ETsend; 
RECEIVE_DATA_STRUCTURE maindata;
SEND_STATE  senddata;

void setup() {
  Serial.begin(115200);
  Wire.begin(THIS_FINGER_ID); 
  senddata.id = THIS_FINGER_ID;
  ETmain.begin(details(maindata), &Wire);
  ETsend.begin(details(senddata), &Wire);
  Wire.onReceive(receive);
//  Wire.onReceive(receiveEvent);
//  Wire.onRequest(requestEvent);
  
  mid.setGoal(50);
  top.setGoal(50);
//  roll.setGoal(4);
}

void receive(int numBytes) {}

void loop() {

  if(ETmain.receiveData()){
    if(0 < maindata.pos && maindata.pos < 8){
      roll.setGoal(maindata.pos);
      mid.setGoal(maindata.A);
      top.setGoal(maindata.B);
    }else{
      senddata.pos = roll.getPrev();
      senddata.npos = roll.getCurr();
      senddata.state = roll.getState();
      senddata.A = mid.getPos();
      senddata.B = top.getPos();
            
      ETsend.sendData(69);
    }
//    Serial.print(F("pos: "));
//    Serial.print(maindata.pos);
//    Serial.print(F(", A: "));
//    Serial.print(maindata.A);
//    Serial.print(F(", B: "));
//    Serial.println(maindata.B);
  }
  
  mid.tick();
  top.tick();
  roll.tick();
}
