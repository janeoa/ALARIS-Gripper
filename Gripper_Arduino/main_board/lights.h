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
  Lights(const byte* pins){
    _pins = new byte[8];
    for (int i(0); i < 8; i++) {
      this->_pins[i] = pins[i];
    }
  }
  virtual ~Lights(){
    delete[] _pins;
  }
  byte getPin(byte pin){
    return _pins[pin];
  }
  void setLight(byte pin);
  void setLight(byte pina, byte pinb);
  void tick();
private:
  byte* _pins;
  byte _state = CALIBRATING;
};

#endif
