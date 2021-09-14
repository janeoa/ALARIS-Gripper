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

/* finger 0 BEGIN */
//#define THIS_FINGER_ID 0
//Motor   mid(2, 4,  3, A2, 560, 300);
//Motor   top(6, 7,  5, A1, 410, 700);
//Roller roll(8, 9, 10, A3, THIS_FINGER_ID, 800);
/* finger 0 END */

/* finger 2 BEGIN */
#define THIS_FINGER_ID 2
Motor   mid(2, 4,  3, A2, 700, 280);
Motor   top(6, 7,  5, A1, 800, 390);
Roller roll(8, 9, 10, A0,THIS_FINGER_ID, 600);
/* finger 2 END */

/* finger 5 BEGIN */
//#define THIS_FINGER_ID 5
//Motor   mid(2, 4,  3, A2, 500, 700);//591 - 328
//Motor   top(6, 7,  5, A1, 700, 300);
//Roller roll(8, 9, 10, A0, THIS_FINGER_ID, 800);
/* finger 5 END */

struct RECEIVE_DATA_STRUCTURE{
  byte dir;
  byte A;
  byte B;
};

struct SEND_STATE{
  byte id;
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

//  roll.setMove(false);
  mid.setGoal(50);
  top.setGoal(30);
}

void receive(int numBytes) {}

unsigned long last = millis();

void loop() {
  if(ETmain.receiveData()){
    if(maindata.dir == 0){
      roll.setMove(true);
    }
    if(maindata.dir == 1){
      roll.setMove(false);
    }
    mid.setGoal(maindata.A);
    top.setGoal(maindata.B);
//    }else{
//      if(millis()-last>200){
//        last = millis();
//        senddata.id = THIS_FINGER_ID;
//        senddata.A = mid.getPos();
//        senddata.B = top.getPos();
//        ETsend.sendData(69);
//      }
//    }
  }
  
  mid.tick();
  top.tick();
  roll.tick();
}
