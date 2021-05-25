// Motor A connections
int enA = 13;
int in1 = 12;
int in2 = 11;

//int sensorPin = A0; // select the input pin for LDR
//int sensorValue = 0;
  
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
}

void loop() {
  while (Serial.available()==0){}             // wait for user input
  pstn = Serial.parseInt(); 
  Serial.print("I received: ");
  Serial.println(pstn);
  digitalWrite(pstn, HIGH);
  digitalWrite(prev, LOW);
  prev = pstn;
  
  directionControl();
  delay(1000);
//  sensorValue = analogRead(sensorPin); // read the value from the sensor
//  Serial.println(sensorValue); //prints the values coming from the sensor on the screen
//  delay(100);
  
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
  
  // Now change motor directions
  digitalWrite(in1, LOW);
  digitalWrite(in2, HIGH);
  delay(2000);
  
  // Turn off motors
  digitalWrite(in1, LOW);
  digitalWrite(in2, LOW);
}
