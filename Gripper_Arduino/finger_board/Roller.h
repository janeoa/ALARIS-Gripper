#ifndef Roller_h
#define Roller_h

/**
 * States
 */
#define READY 0
#define MOVE 1
#define CALIB 2

class Roller{
  public:
    Roller(){
      // Set all the motor control pins to outputs
      pinMode(enA, OUTPUT);
      pinMode(in1, OUTPUT);
      pinMode(in2, OUTPUT);
      
      // Turn off motors - Initial state
      digitalWrite(in1, LOW);
      digitalWrite(in2, LOW);
    }

    void tick();
  private:
    const byte enA = 13;
    const byte in1 = 12;
    const byte in2 = 11;
    const byte sensorPin = A0;

    int lightCalib;
    int sensorValue;
    int state = CALIB;
    int lightCal;
    int x;
};

#endif
