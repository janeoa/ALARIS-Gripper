#ifndef Motor_h
#define Motor_h

#include "Arduino.h"


class Motor
{
  public:
    Motor(int pinA, int pinB, int pinE, int anal, int minp, int maxp);
    void setGoal(int goal);
    void tick();
    int getGoal(){return _goal;}
    int getPos();
  private:
    int _pinA;
    int _pinB;
    int _pinE;
    int _anal;
    int _minp;
    int _maxp;
    
    int _rawadc;
    int _pos;
    int _goal;
};


#endif
