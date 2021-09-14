#ifndef Roller_h
#define Roller_h

//#define THIS_FINGER_ID 0
//#define THIS_FINGER_ID 5
//#define THIS_FINGER_ID 2
/**
 * States
 */
#define READY 0
#define MOVE 1
#define CALIB 2

class Roller{
  public:
    Roller(byte pinA, byte pinB, byte pinE, byte anal, byte id, int th){
      in1 = pinA; in2 = pinB; enA = pinE; sensorPin = anal;
      THIS_FINGER_ID = id;
      threshould = th;
      // Set all the motor control pins to outputs
      pinMode(enA, OUTPUT);
      pinMode(in1, OUTPUT);
      pinMode(in2, OUTPUT);
      
      // Turn off motors - Initial state
      digitalWrite(in1, LOW);
      digitalWrite(in2, LOW);
    }

    void tick();

    void setMove(bool dir){
      _dir = dir;
      state = MOVE; 
    }

    int getAnal(){return sensorValue;}

    byte getState(){return state;}
  private:
    byte enA,in1,in2,sensorPin;

    int lightCalib;
    int sensorValue;
    byte state = CALIB;
    int lightCal;
    bool _dir = 0;
    byte THIS_FINGER_ID = 0;
    int threshould = 700;

    void directionControl(bool toTheRight) {
      analogWrite(enA, 255);
      digitalWrite(in1, toTheRight);
      digitalWrite(in2, !toTheRight);
    }
    
    bool is_next_on_right(int prev, int curr){
      bool is_right = false;
      
      int delta = curr - prev;
      if(delta > 0 && delta < 4) is_right = true;
      if(delta < -4 && delta > -8) is_right = true;
      return is_right;
    }
};

#endif
