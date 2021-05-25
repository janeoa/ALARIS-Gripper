// C++ code
//
int pstn = 0; // for incoming serial data
int prev = 0;
void setup()
{
  Serial.begin(9600);
  pinMode(13, OUTPUT);
}

void loop()
{
    if (Serial.available() > 0) {
    // read
    pstn = Serial.parseInt();

    // say what you got:
    Serial.print("I received: ");
    Serial.println(pstn);
    }
  	digitalWrite(prev, LOW);
    digitalWrite(pstn, HIGH);
  	prev = pstn;
}