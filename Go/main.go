package main

import (
	"github.com/stianeikeland/go-rpio/v4"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

const (
	touchPin = 20
	gapTime  = 20
	monitor  = ":0.0"
)

func main() {
	/* If you remotely debug/run the please add the xauth entry
	cmd := exec.Command("sudo", "-u", "pi", "xauth", "list", monitor)
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Get xauth failed ! %s", err.Error())
	}
	display := strings.Replace(output.String(), "\n", "", -1)
	exec.Command("sudo", "-S", "xauth", "add", display).Run()
	*/
	if err := rpio.Open(); err != nil {
		log.Printf("Failed to open rpio: %s", err.Error())
	}
	c := make(chan int)
	go touchListener(c)
	for {
		flag := <-c
		switch flag {
		case 1:
			exec.Command("xset", "-display", monitor, "dpms", "force", "off").Run()
		case 2:
			exec.Command("xset", "-display", monitor, "dpms", "force", "on").Run()
		}
	}
}

func touchListener(c chan int) {
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
					c <- 2
				} else {
					c <- 1
				}
				*isTouch = false
				*signalArray = ""
			}
		}
	}
}
