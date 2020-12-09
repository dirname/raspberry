package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
	"net"
	"raspberry/Command"
	"raspberry/Logs"
	"raspberry/Models"
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

func ChannelListener(touchPanel *chan int, battery *chan Models.BatteryReport) {
	go Sensor.TouchListener(touchPanel)
	go Sensor.BatteryListener(battery)
	screen := Command.Screen{DisplayName: monitor}
	ups := Models.UPS{}
	for {
		select {
		case flag := <-*touchPanel:
			switch flag {
			case 1:
				screen.TurnOff()
			case 2:
				screen.WakeUp()
			}
		case flag := <-*battery:
			switch flag.State {
			case 0:
				ups.Info = flag.Info
				ups.Parse()
				//log.Printf("Have get the report len: %s from %d\n", flag.Info, flag.ID)
			case 1:
				logrus.Errorf("Get report failed: %s from %d\n", flag.Info, flag.ID)
				go Sensor.BatteryListener(battery)
			}
		}

	}
}

func initSensor() {
	if err := rpio.Open(); err != nil {
		logrus.Fatalf("Failed to open rpio: %s\n", err.Error())
	}
	touchChannel := make(chan int)
	batteryChannel := make(chan Models.BatteryReport)
	go ChannelListener(&touchChannel, &batteryChannel)
}

func main() {
	Logs.Setup()
	logrus.Infof("Your Intranet IP: %s\n", getIntranetIP())
	initSensor()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Run(":5000")
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
}
