package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
	"net"
	"raspberry/command"
	"raspberry/database"
	"raspberry/logs"
	"raspberry/models"
	"raspberry/sensor"
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

func ChannelListener(touchPanel *chan int, battery *chan models.BatteryReport) {
	go sensor.TouchListener(touchPanel)
	go sensor.BatteryListener(battery)
	screen := command.Screen{DisplayName: monitor}
	ups := models.UPS{}
	if err := database.RedisClient.Set(database.Ctx, "battery", &ups, 0).Err(); err != nil {
		logrus.Fatalf("Failed to set UPS address to redis %s", err.Error())
	}
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
				go sensor.BatteryListener(battery)
			}
		}

	}
}

func initSensor() {
	if err := rpio.Open(); err != nil {
		logrus.Fatalf("Failed to open rpio: %s\n", err.Error())
	}
	touchChannel := make(chan int)
	batteryChannel := make(chan models.BatteryReport)
	go ChannelListener(&touchChannel, &batteryChannel)
}

func main() {
	defer database.RedisClient.Close()
	logs.Setup()
	logrus.Infof("Your Intranet IP: %s\n", getIntranetIP())
	initSensor()
	gin.SetMode(gin.DebugMode)
	router := initRouter()
	router.Run(":5000")

	/* If you remotely debug/run the please add the xauth entry
	cmd := exec.command("sudo", "-u", "pi", "xauth", "list", monitor)
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Get xauth failed ! %s", err.Error())
	}
	display := strings.Replace(output.String(), "\n", "", -1)
	exec.command("sudo", "-S", "xauth", "add", display).Run()
	*/
}
