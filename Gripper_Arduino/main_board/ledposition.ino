int pstn = 0; // for incoming serial data
int prev = 0;
void setup()
{
  Serial.begin(9600);
  pinMode(9, OUTPUT);
  pinMode(8, OUTPUT);
  pinMode(7, OUTPUT);
  pinMode(6, OUTPUT);
  pinMode(5, OUTPUT);
  pinMode(4, OUTPUT);
  pinMode(3, OUTPUT);
  pinMode(2, OUTPUT);
}

void loop()
{
    while (Serial.available()==0){}             // wait for user input
    pstn = Serial.parseInt(); 
    Serial.print("I received: ");
    Serial.println(pstn);
    digitalWrite(pstn, HIGH);
    digitalWrite(prev, LOW);
    prev = pstn;
}
