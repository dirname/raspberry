#include <stdio.h>
#include <wiringPi.h>

void test();

int main(void) {
    test();
}

void test(void) {

    int LED = 8;
    wiringPiSetup();

    pinMode(LED, INPUT);

    int number = 10;
    int count = 0;
    while (count < 10) {

        printf("LED:%d is on\n", LED);
        digitalWrite(LED, HIGH);
        delay(500);

        printf("LED:%d is off\n", LED);
        digitalWrite(LED, LOW);
        delay(500);

        count++;
    }
}