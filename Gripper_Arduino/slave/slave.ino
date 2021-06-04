#include <Wire.h>

#define READY 0
#define MOVE 1
#define CALIB 2

int enA = 13;
int in1 = 12;
int in2 = 11;
int lightCalib;
int sensorValue;

int state = CALIB;

int sensorPin = A0; 

int lightCal;

int x = -1;
int prev = 4;
int curr = 4;

void setup() {
  Serial.begin(9600);
  Wire.begin(9);
  Wire.onReceive(receiveEvent); 
  
  // Set all the motor control pins to outputs
  pinMode(enA, OUTPUT);
  pinMode(in1, OUTPUT);
  pinMode(in2, OUTPUT);
  
  // Turn off motors - Initial state
  digitalWrite(in1, LOW);
  digitalWrite(in2, LOW);

  pinMode(sensorPin, INPUT);
  lightCal = analogRead(sensorPin);
}
char msg[50];

void loop() {
  sensorValue = analogRead(sensorPin);
  
  if(state == CALIB){   
    lightCalib = analogRead(sensorPin);
    Serial.println(sensorValue);
    state = READY;    
  }

  if(state == READY){
     digitalWrite(in1, LOW);
     digitalWrite(in2, LOW);
  }
  
  if (x > -1){
     state = MOVE;
     prev = curr;
     curr = x;
     x = -1;
  }  
         
  if(state == MOVE){
     // read the value from the sensor
    if (sensorValue < lightCalib + 200){
       directionControl(is_next_on_right(prev, curr));
    }
    else if (sensorValue > lightCalib + 200) {
       state = READY;
       Serial.println("found");
    }
  }

  sprintf(msg, "%1d\t%4d\t%1d\t%1d\t%1d", state, sensorValue, x, prev, curr);
  Serial.println(msg);
}

// This function lets you control spinning direction of motors
void directionControl(bool toTheRight) {
  // Set motors to maximum speed
  // For PWM maximum possible values are 0 to 255
  analogWrite(enA, 255);

  // Turn on motor A & B
  digitalWrite(in1, toTheRight);
  digitalWrite(in2, !toTheRight);
}

void receiveEvent(int bytes) {
  x = Wire.read();    // read one character from the I2C
}

bool is_next_on_right(int prev, int curr){
  bool is_right = false;
  
  int delta = curr - prev;
  if(delta > 0 && delta < 4) is_right = true;
  if(delta < -4 && delta > -8) is_right = true;
  return is_right;
}
