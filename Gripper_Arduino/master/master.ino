#include <Wire.h>
#define BCAL 0
#define ECAL 1

int lcalib;
int x = 1;

int pstn = 0;
int prev = 0;

void setup() {
  Wire.begin();
  Serial.begin(9600);
  delay(500);

  pinMode(9, OUTPUT);
  pinMode(8, OUTPUT);
  pinMode(7, OUTPUT);
  pinMode(6, OUTPUT);
  pinMode(5, OUTPUT);
  pinMode(4, OUTPUT);
  pinMode(3, OUTPUT);
  pinMode(2, OUTPUT);
  lcalib = BCAL;
}

void loop() {
  if (lcalib == BCAL){
    for (int i = 2; i < 10; i++){
        digitalWrite(i, HIGH);
        delay(500);
        digitalWrite(i, LOW);
    }
    lcalib = ECAL;
  }
  if (Serial.available() >= 1) {
    pstn = Serial.parseInt();
    Serial.print("I received: ");
    Serial.println(pstn);
    digitalWrite(pstn, HIGH );
    digitalWrite(prev, LOW);
    prev = pstn;
    Wire.beginTransmission(9);
    Wire.write(x);
    Wire.endTransmission();
  }
}
