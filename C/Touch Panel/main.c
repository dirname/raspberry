#include <stdio.h>
#include <wiringPi.h>

const int touchPin = 20;

int main(void) {
    if (wiringPiSetupGpio() == -1) {
        printf("Setup wiringPi failed !\n");
        return 1;
    }
    pinMode(touchPin, INPUT);
    int *signalArray[20];
    int *index = 0;
    while (TRUE) {
        *signalArray[*index] = digitalRead(touchPin);
        *index++;
        if (*index >= 20) {
            *index = 0;
            printf("%d", *index);
        }
    }
}
