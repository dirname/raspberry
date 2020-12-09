package sensor

import (
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
			if len(*signalArray) > 20 {
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
