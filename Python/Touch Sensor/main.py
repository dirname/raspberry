import os
import re
import time
import RPi.GPIO as GPIO


class Touch:
    def __init__(self):
        self.infrared = 20  # GPIO Number
        self.gap_time = 0.02  # Set the interval to determine the event
        self.signal_threshold = 20
        self.last_sign = 1
        self.last_long = False
        self.touch_signal = False
        display = os.popen("sudo -u pi xauth list :0.0").read()
        os.system("sudo -S xauth add " + display)
        GPIO.setmode(GPIO.BCM)
        GPIO.setup(self.infrared, GPIO.IN)

    def check_event(self, signal):
        if signal == 0:
            self.last_long = True
            os.system("xset -display :0.0 dpms force off")
            print("double")
        else:
            if self.last_long is False:
                os.system("xset -display :0.0 dpms force on")
                print("single")
            else:
                print("double signal")
                self.last_long = False
            self.last_sign = False
        self.last_sign = 1
        self.detect()

    def detect(self):
        try:
            signal_array = ""
            while True:
                if GPIO.input(self.infrared) == 0:
                    self.touch_signal = True
                time.sleep(self.gap_time)
                if self.touch_signal:
                    signal_array += str(GPIO.input(self.infrared))
                if self.touch_signal and len(signal_array) > self.signal_threshold:
                    is_double = re.compile("^0.*10.*1$")
                    if is_double.match(signal_array):
                        # Double Click
                        os.system("xset -display :0.0 dpms force on")
                    else:
                        # Single Click
                        os.system("xset -display :0.0 dpms force off")
                    signal_array = ""
                    self.touch_signal = False  # End of a click time

                # Determine the long press event
                # if self.last_sign == 0:
                #     time.sleep(0.3)
                #     self.check_event(GPIO.input(self.infrared))
                #     break
                # self.last_sign = GPIO.input(self.infrared)
                # time.sleep(0.2)
        except KeyboardInterrupt:
            GPIO.cleanup()
            pass


if __name__ == "__main__":
    senor = Touch()
    senor.detect()
