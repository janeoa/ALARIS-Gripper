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
      if (sensorValue < threshould){
           directionControl(_dir);
      }else{
         state = READY;
      }  
    } break;
    case CALIB:{
      lightCalib = analogRead(sensorPin);
      state = READY;    
    } break;
    default:
    break;
  }
}
