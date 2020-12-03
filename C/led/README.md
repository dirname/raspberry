# Light up a LED

```shell
gcc -o led -lwiringPi main.c
./led
```

# CGO
```Go
package main

/*
#include <stdio.h>
#include <wiringPi.h>

#cgo LDFLAGS: -lwiringPi

void test(void) {

    int LED = 8;
    wiringPiSetup();

    pinMode(LED, OUTPUT);

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
*/
import "C"

func main() {
	C.test()
}

```