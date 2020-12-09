package Sensor

import (
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"math/rand"
	"raspberry/Models"
	"time"
)

func BatteryListener(c *chan Models.BatteryReport) {
	id := rand.Int()
	config := &serial.Config{
		Name: "/dev/ttyAMA0",
		Baud: 9600,
	}
	s, err := serial.OpenPort(config)
	if err != nil {
		logrus.Fatalf("Failed to open the serial: %s\n", err.Error())
	}
	defer s.Close()
	buf := make([]byte, 1024)
	result := Models.BatteryReport{}
	for {
		num, err := s.Read(buf)
		if err != nil {
			result.State = 1
			result.Info = err.Error()
			result.ID = id
			*c <- result
			s.Flush()
			break
		}
		if num > 0 {
			result.State = 0
			result.Info = string(buf)
			result.ID = id
			*c <- result
		}
		time.Sleep(1 * time.Second)
	}
}
