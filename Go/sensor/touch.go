package sensor

import (
	"github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
	"regexp"
	"strconv"
	"time"
)

const (
	touchPin = 20
	gapTime  = 20
)

func TouchListener(c *chan int) {
	if err := rpio.Open(); err != nil {
		logrus.Fatalf("Failed to open rpio: %s\n", err.Error())
	}
	pin := rpio.Pin(touchPin) // find touch panel
	pin.Input()               // set mode to input
	var isTouch = new(bool)
	var signalArray = new(string)
	var signal = new(rpio.State)
	for true {
		*signal = pin.Read()
		if *signal == 0 {
			*isTouch = true
		}
		time.Sleep(time.Millisecond * gapTime)
		if *isTouch {
			*signalArray += strconv.Itoa(int(*signal))
			if len(*signalArray) > 20 && *isTouch {
				re := regexp.MustCompile("^0.*10.*1$")
				if re.MatchString(*signalArray) {
					*c <- 2
				} else {
					*c <- 1
				}
				*isTouch = false
				*signalArray = ""
			}
		}
	}
}
