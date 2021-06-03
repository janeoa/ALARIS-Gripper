#include "Arduino.h"
#include "Lights.h"

void Lights::tick(){
  switch(_state){
    case CALIBRATING:{
      for (int j=0;j<4;j++){
        for (int i = 0; i < 8; i++){
          digitalWrite(_pins[i], HIGH);
          delay(25);
          digitalWrite(_pins[i], LOW);
        }
      }
      _state = READY;
    }break;
    case READY:{
      for (int i = 0; i < 8; i++){
        digitalWrite(_pins[i], LOW);
      }
    }break;
  }
}

void Lights::setLight(int pin){
  for (int i = 0; i < 8; i++){
    digitalWrite(_pins[i], LOW);
  }
  digitalWrite(_pins[pin],HIGH);
  _state = SET;
}

void Lights::setLight(int pina, int pinb){
  for (int i = 0; i < 8; i++){
    digitalWrite(_pins[i], LOW);
  }
  digitalWrite(_pins[pina],HIGH);
  digitalWrite(_pins[pinb],HIGH);
  _state = SET;
}
