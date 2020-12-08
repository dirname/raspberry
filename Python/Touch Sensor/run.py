import time

import RPi.GPIO as GPIO

GPIO.setmode(GPIO.BCM)
GPIO.setup(12, GPIO.OUT)

try:
    while True:
        GPIO.output(12, GPIO.HIGH)
        # GPIO.output(12, GPIO.LOW)
        # time.sleep(3)
except KeyboardInterrupt:
    GPIO.cleanup()
    pass
