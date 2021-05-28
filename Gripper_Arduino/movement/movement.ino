#define READY 0
#define MOVE 1
#define CALIB 2

// Motor A connections
int enA = 13;
int in1 = 12;
int in2 = 11;
int lightCalib;
int sensorValue;

int state = CALIB;

//resistance drops, we can check changes in analogread
int sensorPin = A0; // select the input pin for LDR

int lightCal;

int pstn = 0; // for incoming serial data
int prev = 0;

void setup() {
  Serial.begin(9600); 
  pinMode(9, OUTPUT);
  pinMode(8, OUTPUT);
  pinMode(7, OUTPUT);
  pinMode(6, OUTPUT);
  pinMode(5, OUTPUT);
  pinMode(4, OUTPUT);
  pinMode(3, OUTPUT);
  pinMode(2, OUTPUT);
  
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

void loop() {
  if(state == CALIB){
    analogWrite(enA, 255);
    digitalWrite(in1, HIGH);
    digitalWrite(in2, LOW);
    delay(2000);
    digitalWrite(in1, LOW);
    digitalWrite(in2, HIGH);
    delay(2000);
    digitalWrite(in1, LOW);
    digitalWrite(in2, LOW);
    
    lightCalib = analogRead(sensorPin);
    Serial.println(sensorValue);
    state = READY;    
  }
  if (Serial.available()>= 1){
    pstn = Serial.parseInt(); 
    Serial.print("I received: ");
    Serial.println(pstn);
    digitalWrite(pstn, HIGH);
    digitalWrite(prev, LOW);
    prev = pstn;
    state = MOVE;
    }             // wait for user input
  if(state == MOVE){
  sensorValue = analogRead(sensorPin); // read the value from the sensor
  Serial.println(sensorValue); //prints the values coming from the sensor on the screen
    if (sensorValue < lightCalib + 100){
       directionControl();
       state = READY;
       Serial.println("found");
  }
  else if (sensorValue > lightCalib + 100) {
    digitalWrite(in1, LOW);
    digitalWrite(in2, LOW);
  }
  }
  else {
    digitalWrite(in1, LOW);
    digitalWrite(in2, LOW);    
  }
}

// This function lets you control spinning direction of motors
void directionControl() {
  // Set motors to maximum speed
  // For PWM maximum possible values are 0 to 255
  analogWrite(enA, 255);

  // Turn on motor A & B
  digitalWrite(in1, HIGH);
  digitalWrite(in2, LOW);
  delay(2000);
}
//comment
