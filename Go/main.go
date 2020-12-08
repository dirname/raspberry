package main

import (
	"github.com/stianeikeland/go-rpio/v4"
	"log"
	"net"
	"raspberry/Command"
	"raspberry/Sensor"
)

const (
	monitor = ":0.0"
)

func getIntranetIP() string {
	adders, _ := net.InterfaceAddrs()
	for _, addr := range adders {
		if inet, ok := addr.(*net.IPNet); ok && !inet.IP.IsLoopback() {
			if inet.IP.To4() != nil {
				return inet.IP.String()
			}
		}
	}
	return ""
}

func main() {
	log.Printf("Your Intranet IP: %s\n", getIntranetIP())
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
	touchChanel := make(chan int)
	go Sensor.TouchListener(&touchChanel)
	screen := Command.Screen{DisplayName: monitor}
	for {
		flag := <-touchChanel
		switch flag {
		case 1:
			screen.TurnOff()
		case 2:
			screen.WakeUp()
		}
	}
}
