#include "Arduino.h"
#include "Motor.h"

#define tolerance 100

Motor::Motor(int pinA, int pinB, int pinE, int anal, int minp, int maxp){
  _pinA = pinA;
  _pinB = pinB;
  _pinE = pinE;
  _anal = anal;
  _minp = minp;
  _maxp = maxp;

  pinMode(pinA, OUTPUT);
  pinMode(pinB, OUTPUT);
  pinMode(pinE, OUTPUT);
  digitalWrite(pinE, HIGH);
}

void Motor::setGoal(int goal){
  if(0<=goal && goal<=100){
    _goal = map(goal, 0, 100, _minp, _maxp);
  }
}

void Motor::tick(){
  _rawadc = analogRead(_anal);
  if(_goal > _rawadc - tolerance && _goal < _rawadc + tolerance){
    digitalWrite(_pinA, LOW);
    digitalWrite(_pinB, LOW);
  }else if(_goal > _rawadc){
    digitalWrite(_pinA, HIGH);
    digitalWrite(_pinB, LOW);
  }else{
    digitalWrite(_pinA, LOW);
    digitalWrite(_pinB, HIGH);
  }
}

int Motor::getPos(){
  _rawadc = analogRead(_anal);
//  Serial.print("EPTA");
//  Serial.println(_anal);
  return map(_rawadc, 0, 100, _minp, _maxp);;
}
