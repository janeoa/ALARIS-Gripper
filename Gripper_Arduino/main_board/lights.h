#ifndef Lights_h
#define Lights_h

/**
 * States: READY, MOVING
 */
#define READY       0
#define CALIBRATING 1
#define SET         2
#define SYNC        3

#include "Arduino.h"

class Lights{
public:
  Lights(const int* pins){
    _pins = new int[8];
    for (int i(0); i < 8; i++) {
      this->_pins[i] = pins[i];
    }
  }
  virtual ~Lights(){
    delete[] _pins;
  }
  int getPin(int pin){
    return _pins[pin];
  }
  void setLight(int pin);
  void setLight(int pina, int pinb);
  void tick();
private:
  int* _pins;
  byte _state = CALIBRATING;
};

#endif
