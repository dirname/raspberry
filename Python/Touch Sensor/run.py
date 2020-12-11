import time

import RPi.GPIO as GPIO

GPIO.setmode(GPIO.BCM)
GPIO.setup(18, GPIO.IN)

print("go")

try:
    while True:
        print(GPIO.input(18))
        print(1)
        # GPIO.output(12, GPIO.HIGH)
        # GPIO.output(12, GPIO.LOW)
        # time.sleep(3)
except KeyboardInterrupt:
    GPIO.cleanup()
    pass
