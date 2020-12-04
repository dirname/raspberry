import time

import RPi.GPIO as GPIO
from gpiozero import Buzzer

GPIO.cleanup()
GPIO.setmode(GPIO.BCM)
GPIO.setup(27, GPIO.OUT)

try:
    while True:
        GPIO.output(27, GPIO.HIGH)
        time.sleep(0.2)
        GPIO.output(27, GPIO.LOW)
except KeyboardInterrupt:
    pass
