#include "Arduino.h"
#include "Roller.h"

void Roller::tick(){
  sensorValue = analogRead(sensorPin);
  
  switch(state){
    case READY:{
      digitalWrite(in1, LOW);
      digitalWrite(in2, LOW);  
    } break;
    case MOVE:{
      if (sensorValue < lightCalib + 200){
         directionControl(is_next_on_right(prev, curr));
      }
      else if (sensorValue > lightCalib + 200) {
         state = READY;
//         Serial.println("found");
      }  
    } break;
    case CALIB:{
      lightCalib = analogRead(sensorPin);
//      Serial.println(sensorValue);
      state = READY;    
    } break;
  }

  if (x > -1){
     state = MOVE;
     prev = curr;
     curr = x;
     x = -1;
  }  
}
