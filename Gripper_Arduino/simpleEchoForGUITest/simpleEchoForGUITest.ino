void setup() {
  // put your setup code here, to run once:
  Serial.begin(115200);
}


long last = millis();

char buff[100];

void loop() {
  // put your main code here, to run repeatedly:
  if(millis()-last>400){
    Serial.println(millis());
    last = millis();
  }

  if(Serial.available()){
    Serial.readBytesUntil("\n", buff, 100);
    Serial.println(buff);
    for(int i=0; i<100; i++){
      buff[i] = 0x00;
    }
  }
}
